package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gin-myboot/global"
	"gin-myboot/modules/common/model/response"
	"gin-myboot/utils"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type ImageValidateForm struct {
	CaptchaId string `json:"captchaId"`
	Code      string `json:"code"`
}

type SmsValidateForm struct {
	Mobile  string `json:"mobile"`
	SmsCode string `json:"smsCode"`
}

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
//var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

// ImageValidate 图片验证码中间件
func ImageValidate() gin.HandlerFunc {
	return func(c *gin.Context) {

		var imageValidateForm ImageValidateForm

		var body []byte
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				global.Logger.Error("read body from request error:", zap.Any("err", err))
				response.FailWithDetailed(gin.H{}, "验证码错误", c)
				c.Abort()
				return
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
				if err == json.Unmarshal(body, &imageValidateForm) && err == nil {
					global.Logger.Error("ImageValidate", zap.Any("imageValidateForm", imageValidateForm))
				}
			}
		} else {
			imageValidateForm.Code = c.Query("code")
			imageValidateForm.CaptchaId = c.Query("captchaId")
		}

		if imageValidateForm.Code == "" {
			response.FailWithDetailed(gin.H{}, "验证码不能为空", c)
			c.Abort()
			return
		}

		if store.Verify(imageValidateForm.CaptchaId, imageValidateForm.Code, true) {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "验证码错误", c)
			c.Abort()
			return
		}
	}
}

// SmsCodeValidate 短信验证中间件
func SmsCodeValidate() gin.HandlerFunc {
	return func(c *gin.Context) {

		var smsValidateForm SmsValidateForm

		var body []byte
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				global.Logger.Error("read body from request error:", zap.Any("err", err))
				response.FailWithDetailed(gin.H{}, "手机验证码错误", c)
				c.Abort()
				return
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
				if err == json.Unmarshal(body, &smsValidateForm) && err == nil {
					global.Logger.Error("SmsCodeValidate", zap.Any("smsValidateForm", smsValidateForm))
				}
			}
		} else {
			smsValidateForm.SmsCode = c.Query("smsCode")
			smsValidateForm.Mobile = c.Query("mobile")
		}

		if smsValidateForm.SmsCode == "" {
			response.FailWithDetailed(gin.H{}, "手机验证码不能为空", c)
			c.Abort()
			return
		}

		if err, _ := utils.CheckSmsCode(smsValidateForm.Mobile, smsValidateForm.SmsCode, true); err == nil {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "手机验证码错误1", c)
			c.Abort()
			return
		}
	}
}

type CaptchaVerificationRequest struct {
	CaptchaVerification string `json:"captchaVerification"`
}

// CaptchaValidate 滑块验证中间件
func CaptchaValidate() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req CaptchaVerificationRequest

		var body []byte
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				global.Logger.Error("read body from request error:", zap.Any("err", err))
				response.FailWithDetailed(gin.H{}, "验证码错误", c)
				c.Abort()
				return
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
				if err == json.Unmarshal(body, &req) && err == nil {
					global.Logger.Error("CaptchaValidate", zap.Any("req", req))
				}
			}
		} else {
			req.CaptchaVerification = c.Query("captchaVerification")
		}

		if req.CaptchaVerification == "" {
			response.FailWithDetailed(gin.H{}, "验证码错误", c)
			c.Abort()
			return
		}

		ctx := context.Background()

		key := "RUNNING:CAPTCHA:second-" + req.CaptchaVerification
		result, err := global.Redis.Get(ctx, key).Result()
		if err != nil {
			response.FailWithDetailed(gin.H{}, "验证码错误", c)
			c.Abort()
			return
		}

		fmt.Printf("=======> Verification %s \n", result)

		global.Redis.Del(ctx, key).Result()

		c.Next()
	}
}
