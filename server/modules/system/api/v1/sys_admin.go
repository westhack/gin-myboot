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

type AdminApi struct {
}

// GetList
// @Tags SysAdmin
// @Summary 分页获取管理员列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchParams true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/admin/getList [post]
func (b *AdminApi) GetList(c *gin.Context) {
	var searchParams request.QueryParams
	_ = c.ShouldBindJSON(&searchParams)
	if err := utils.Verify(searchParams, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, list, total := adminService.GetList(searchParams); err != nil {
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

// GetAll
// @Tags SysAdmin
// @Summary 获取所有可用管理员
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/admin/getAll [get]
func (b *AdminApi) GetAll(c *gin.Context) {
	if err, list, _ := adminService.GetAll(10000); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"items": list}, "获取成功", c)
	}
}

// Create
// @Tags SysAdmin
// @Summary 创建管理员
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AdminFormRequest true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /v1/system/admin/create [post]
func (b *AdminApi) Create(c *gin.Context) {
	var user systemReq.AdminFormRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailWithMessage("创建失败，"+utils.GetError(err, user), c)
		return
	}

	user.Id = 0

	if err, resUser := adminService.Create(&user); err != nil {
		global.Error("创建失败", err)
		response.FailWithMessage("创建失败，"+err.Error(), c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": resUser}, "创建成功", c)
	}
}

// Update
// @Tags SysAdmin
// @Summary 修改管理员
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AdminFormRequest true "ID, 用户名, 昵称, 头像链接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /v1/system/admin/update [post]
func (b *AdminApi) Update(c *gin.Context) {
	var user systemReq.AdminFormRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailWithMessage("更新失败，"+utils.GetError(err, user), c)
		return
	}

	if err, ReqUser := adminService.Update(&user); err != nil {
		global.Error("更新失败", err)
		response.FailWithMessage("更新失败，"+err.Error(), c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "更新成功", c)
	}
}

// Delete
// @Tags SysAdmin
// @Summary 删除管理员
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "用户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/system/admin/delete [post]
func (b *AdminApi) Delete(c *gin.Context) {
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

	if err := adminService.Delete(reqId.ID); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds
// @Tags SysAdmin
// @Summary 批量删除管理员
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "用户IDs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/system/admin/deleteByIds [post]
func (b *AdminApi) DeleteByIds(c *gin.Context) {
	var reqId request.GetByIds
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	jwtId := utils.GetUserID(c)

	for _, id := range reqId.ID {
		if jwtId == id {
			response.FailWithMessage("删除失败, 无法删除自己", c)
			return
		}
	}

	if err := adminService.DeleteByIds(reqId.ID); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// ChangeStatus
// @Tags SysAdmin
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "用户IDs，状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /v1/system/admin/changeStatus [post]
func (b *AdminApi) ChangeStatus(c *gin.Context) {
	var req request.ChangeStatus
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := adminService.ChangeStatus(req.ID, req.Status); err != nil {
		global.Logger.Error("修改用户状态!", zap.Any("err", err))
		response.FailWithMessage("修改用户状态失败", c)
	} else {
		response.OkWithMessage("修改用户状态成功", c)
	}
}

// ResetPassword
// @Tags SysAdmin
// @Summary 重置密码为 123456
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "用户IDs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /v1/system/admin/resetPassword [post]
func (b *AdminApi) ResetPassword(c *gin.Context) {
	var reqId request.GetByIds
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := adminService.ResetPassword(reqId.ID); err != nil {
		global.Logger.Error("重置密码失败!", zap.Any("err", err))
		response.FailWithMessage("重置密码失败", c)
	} else {
		response.OkWithMessage("重置密码成功", c)
	}
}
