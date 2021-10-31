package v1

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/request"
	"gin-myboot/modules/common/model/response"
	systemReq "gin-myboot/modules/system/model/request"
	"gin-myboot/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct {
}

// GetList
// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchParams true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/user/getList [post]
func (b *UserApi) GetList(c *gin.Context) {
	var searchParams request.QueryParams
	_ = c.ShouldBindJSON(&searchParams)
	if err := utils.Verify(searchParams, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, list, total := userService.GetList(searchParams); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     searchParams.Page,
			PageSize: searchParams.PageSize,
		}, "获取成功", c)
	}
}

// Create
// @Tags SysUser
// @Summary 创建用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UserFormRequest true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /v1/system/user/create [post]
func (b *UserApi) Create(c *gin.Context) {
	var user systemReq.UserFormRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailWithMessage("创建失败，"+utils.GetError(err, user), c)
		return
	}
	if err, ReqUser := userService.Create(&user); err != nil {
		global.Error("创建失败", err)
		response.FailWithMessage("创建失败，"+err.Error(), c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "创建成功", c)
	}
}

// Update
// @Tags SysUser
// @Summary 修改用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UserFormRequest true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /v1/system/user/update [post]
func (b *UserApi) Update(c *gin.Context) {
	var user systemReq.UserFormRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailWithMessage("更新失败，"+utils.GetError(err, user), c)
		return
	}
	if err, ReqUser := userService.Update(&user); err != nil {
		global.Error("更新失败", err)
		response.FailWithMessage("更新失败，"+err.Error(), c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "更新成功", c)
	}
}

// Delete
// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "用户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/system/user/delete [post]
func (b *UserApi) Delete(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := utils.GetUserID(c)
	if jwtId == reqId.ID {
		response.FailWithMessage("删除失败, 无法删除自己", c)
		return
	}
	if err := userService.Delete(reqId.ID); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds
// @Tags SysUser
// @Summary 批量删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "用户IDs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/system/user/deleteByIds [post]
func (b *UserApi) DeleteByIds(c *gin.Context) {
	var reqId request.GetByIds
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := utils.GetUserID(c)

	for _, id := range reqId.ID {
		if jwtId == id {
			response.FailWithMessage("删除失败, 无法上传自己", c)
			return
		}
	}

	if err := userService.DeleteByIds(reqId.ID); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// ChangeStatus
// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "用户IDs，状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /v1/system/user/changeStatus [post]
func (b *UserApi) ChangeStatus(c *gin.Context) {
	var req request.ChangeStatus
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := userService.ChangeStatus(req.ID, req.Status); err != nil {
		global.Logger.Error("修改用户状态!", zap.Any("err", err))
		response.FailWithMessage("修改用户状态失败", c)
	} else {
		response.OkWithMessage("修改用户状态成功", c)
	}
}

// ResetPassword
// @Tags SysUser
// @Summary 重置密码为 123456
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "用户IDs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /v1/system/user/resetPassword [post]
func (b *UserApi) ResetPassword(c *gin.Context) {
	var reqId request.GetByIds
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := userService.ResetPassword(reqId.ID); err != nil {
		global.Logger.Error("重置密码失败!", zap.Any("err", err))
		response.FailWithMessage("重置密码失败", c)
	} else {
		response.OkWithMessage("重置密码成功", c)
	}
}
