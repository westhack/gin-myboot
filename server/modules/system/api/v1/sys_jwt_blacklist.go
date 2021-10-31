package v1

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/response"
	system "gin-myboot/modules/system/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type JwtApi struct {
}

// JsonInBlacklist
// @Tags Jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /v1/system/jwt/jsonInBlacklist [post]
func (j *JwtApi) JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := system.JwtBlacklist{Jwt: token}
	if err := jwtService.JsonInBlacklist(jwt); err != nil {
		global.Logger.Error("jwt作废失败!", zap.Any("err", err))
		response.FailWithMessage("jwt作废失败", c)
	} else {
		response.OkWithMessage("jwt作废成功", c)
	}
}
