package middleware

import (
	"bytes"
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	"gin-myboot/modules/system/model/request"
	"gin-myboot/modules/system/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var userLogService = service.SystemServiceGroup.LogService

func SystemLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		var userId uint64
		username := ""
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				global.Logger.Error("read body from request error:", zap.Any("err", err))
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		}
		if claims, ok := c.Get("claims"); ok {
			waitUse := claims.(*request.CustomClaims)
			userId = waitUse.ID
			username = waitUse.Username
		} else {
			id, err := strconv.ParseUint(c.Request.Header.Get("x-user-id"), 10, 64)
			if err != nil {
				userId = 0
			}
			userId = id

			username = string(c.Request.Header.Get("x-username"))
		}
		record := system.SysLog{
			Ip:         c.ClientIP(),
			Method:     c.Request.Method,
			RequestUrl: c.Request.URL.RequestURI(),
			Agent:      c.Request.UserAgent(),
			Body:       string(body),
			UserID:     userId,
			Username:   username,
		}
		// 存在某些未知错误 TODO
		//values := c.Request.Header.Values("content-type")
		//if len(values) >0 && strings.Contains(values[0], "boundary") {
		//	record.Body = "file"
		//}
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		latency := time.Now().Sub(now)
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()

		if err := userLogService.Create(record); err != nil {
			global.Logger.Error("create user logs record error:", zap.Any("err", err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
