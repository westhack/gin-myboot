package v1

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// EmailTest
// @Tags System
// @Summary 发送测试邮件
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /v1/system/email/emailTest [post]
func (s *SystemApi) EmailTest(c *gin.Context) {
	if err := emailService.EmailTest(); err != nil {
		global.Logger.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}
