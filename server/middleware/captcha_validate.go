package middleware

import (
	"bytes"
	"encoding/json"
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
	Code string `json:"code"`
}

type SmsValidateForm struct {
	Mobile string `json:"mobile"`
	SmsCode string `json:"smsCode"`
}


// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
//var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

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
