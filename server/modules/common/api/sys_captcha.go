package api

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/response"
	"gin-myboot/modules/system/model/request"
	systemRes "gin-myboot/modules/system/model/response"
	"gin-myboot/utils"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
//var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type CaptchaApi struct {
}

// Captcha
// @Tags Captcha
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /v1/common/captcha/get [post]
func (b *CaptchaApi) Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.Config.Captcha.ImgHeight, global.Config.Captcha.ImgWidth, global.Config.Captcha.KeyLong, 0.7, 80)
	//cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.Logger.Error("验证码获取失败!", zap.Any("err", err))
		response.FailWithMessage("验证码获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}

// SendSms
// @Tags Captcha
// @Summary 获取手机验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SmsCodeRequest true "手机号，留空获取用户自身手机号"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /v1/common/captcha/sendSms [post]
func (b *CaptchaApi) SendSms(c *gin.Context) {
	var req request.SmsCodeRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}

	if req.Mobile == "" {
		err, user := utils.GetUser(c)
		if err != nil {
			response.FailWithMessage("手机号不能为空", c)
			return
		}
		req.Mobile = user.Mobile
	}

	if req.Mobile == "" {
		response.FailWithMessage("手机号不能为空", c)
		return
	}

	err, smsCode := utils.SendSmsCode(req.Mobile)
	if err != nil {
		response.FailWithMessage("获取失败" + err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{"smsCode": smsCode, "mobile": req.Mobile}, "手机验证码获取成功", c)
}
