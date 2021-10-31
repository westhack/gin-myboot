package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Response struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Msg       string      `json:"message"`
	Success   bool        `json:"success"`
	Timestamp int64       `json:"timestamp"`
}

const (
	SUCCESS            = 200 // 成功
	ERROR              = 500 // 失败
	AUTHENTICATED      = 400 // 请求错误
	UNAUTHORIZED       = 401 // 未登录
	UNAUTHENTICATED    = 403 // 未授权
	NOT_FOUND          = 404 // 内容不存在
	METHOD_NOT_ALLOWED = 405 // 方法不允许
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	var success bool
	if code == 200 {
		success = true
	} else {
		success = false
	}

	timestamp := time.Now().UnixMicro()
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
		success,
		timestamp,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func UnauthorizedWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(UNAUTHORIZED, data, message, c)
}

func UnauthenticatedWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(UNAUTHENTICATED, data, message, c)
}
