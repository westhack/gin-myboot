package api

import (
	"gin-myboot/global"
	"gin-myboot/modules/captcha/model"
	"gin-myboot/modules/captcha/service"
	"gin-myboot/modules/common/model/response"
	"github.com/gin-gonic/gin"
)

type CaptchaApi struct {
}

var CaptchaApiApp = new(CaptchaApi)

// Get 获取滑块验证码
// @Tags Captcha
// @Summary 获取滑块验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CaptchaRequest true "验证码类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /captcha/get [post]
func (api *CaptchaApi) Get(c *gin.Context) {
	client := global.Redis
	var captchaReq model.CaptchaRequest
	var checkRes model.CaptchaInfo
	checkRes.RepCode = "6110"
	err := c.ShouldBindJSON(&captchaReq)
	if err != nil {

		checkRes.RepCode = "6110"
		checkRes.RepMsg = "获取失败"
		c.JSON(200, checkRes)

		return
	}

	if captchaReq.CaptchaType == "clickWord" {
		clickWord, err := service.GetClickWord(client)
		if err != nil {
			return
		}

		checkRes.RepCode = "0000"
		checkRes.RepData = clickWord
		c.JSON(200, checkRes)

		return
	}

	puzzle, err := service.GetBlockPuzzle(client)
	if err != nil {
		return
	}

	checkRes.RepCode = "0000"
	checkRes.RepData = puzzle

	c.JSON(200, checkRes)
}

// Check 滑块验证码验证
// @Tags Captcha
// @Summary 滑块验证码验证
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CaptchaCheckRequest true "验证码数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证失败"}"
// @Router /captcha/check [post]
func (api *CaptchaApi) Check(c *gin.Context) {
	client := global.Redis
	var captchaCheckReq model.CaptchaCheckRequest
	var captchaInfo model.CaptchaInfo
	err := c.ShouldBindJSON(&captchaCheckReq)
	if err != nil {

		captchaInfo.RepCode = "6110"
		captchaInfo.RepMsg = "验证码已失效，请重新获取"
		c.JSON(200, captchaInfo)
		return
	}

	if captchaCheckReq.CaptchaType == "clickWord" {
		_, err := service.ClickWordCheck(client, captchaCheckReq, true)
		if err != nil {
			captchaInfo.RepCode = "6110"
			captchaInfo.RepMsg = "验证码已失效，请重新获取"
			c.JSON(200, captchaInfo)
			return
		}

		captchaInfo.RepCode = "0000"
		captchaInfo.RepMsg = "验证成功"
		c.JSON(200, captchaInfo)

		return
	} else {
		_, err := service.BlockPuzzleCheck(client, captchaCheckReq, true)
		if err != nil {
			captchaInfo.RepCode = "6110"
			captchaInfo.RepMsg = "验证码已失效，请重新获取"
			c.JSON(200, captchaInfo)
			return
		}

		captchaInfo.RepCode = "0000"
		captchaInfo.RepMsg = "验证成功"
		c.JSON(200, captchaInfo)
	}
}

// Verification 二次验证演示
// @Tags Captcha
// @Summary 二次验证演示
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CaptchaVerificationRequest true "验证码数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证失败"}"
// @Router /captcha/verification [post]
func (api *CaptchaApi) Verification(c *gin.Context) {
	client := global.Redis
	var captchaVerificationRequest model.CaptchaVerificationRequest

	err := c.ShouldBindJSON(&captchaVerificationRequest)
	if err != nil {
		response.FailWithMessage("验证失败", c)
		return
	}

	_, err = service.Verification(client, captchaVerificationRequest)
	if err != nil {
		response.FailWithMessage("验证失败"+err.Error(), c)
		return
	}
	// 其它业务
	response.OkWithMessage("验证成功", c)
	return
}
