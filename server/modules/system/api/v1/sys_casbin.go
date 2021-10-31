package v1

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/response"
	"gin-myboot/modules/system/model/request"
	systemRes "gin-myboot/modules/system/model/response"
	"gin-myboot/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CasbinApi struct {
}

// UpdateCasbin
// @Tags Casbin
// @Summary 更新角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /v1/system/casbin/UpdateCasbin [post]
func (cas *CasbinApi) UpdateCasbin(c *gin.Context) {
	var cmr request.CasbinInReceive
	_ = c.ShouldBindJSON(&cmr)
	if err := utils.Verify(cmr, utils.RoleIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := casbinService.UpdateCasbin(cmr.RoleId, cmr.CasbinInfos); err != nil {
		global.Logger.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetPolicyPathByRoleId
// @Tags Casbin
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/getPolicyPathByRoleId [post]
func (cas *CasbinApi) GetPolicyPathByRoleId(c *gin.Context) {
	var casbin request.CasbinInReceive
	_ = c.ShouldBindJSON(&casbin)
	if err := utils.Verify(casbin, utils.RoleIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	paths := casbinService.GetPolicyPathByRoleId(casbin.RoleId)
	response.OkWithDetailed(systemRes.PolicyPathResponse{Paths: paths}, "获取成功", c)
}
