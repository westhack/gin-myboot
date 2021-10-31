package v1

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/request"
	"gin-myboot/modules/common/model/response"
	system "gin-myboot/modules/system/model"
	systemReq "gin-myboot/modules/system/model/request"
	"gin-myboot/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RoleApi struct {
}

// Create
// @Tags Role
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysRole true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/system/role/create [post]
func (a *RoleApi) Create(c *gin.Context) {
	var role systemReq.RoleFormRequest
	if err := c.ShouldBindJSON(&role); err != nil {
		response.FailWithMessage("创建失败，" + utils.GetError(err, role), c)
		return
	}

	if err := roleService.Create(role); err != nil {
		global.Error("更新失败!", err)
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Delete
// @Tags Role
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysRole true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/system/role/delete [post]
func (a *RoleApi) Delete(c *gin.Context) {
	var id request.GetById
	_ = c.ShouldBindJSON(&id)
	if err := utils.Verify(id, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := roleService.Delete(id.ID); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds
// @Tags Role
// @Summary 批量删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysRole true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/system/role/deleteByIds [post]
func (a *RoleApi) DeleteByIds(c *gin.Context) {
	var ids request.GetByIds
	_ = c.ShouldBindJSON(&ids)
	if err := utils.Verify(ids, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := roleService.DeleteByIds(ids.ID); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update
// @Tags Role
// @Summary 更新角色信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysRole true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /v1/system/role/update [post]
func (a *RoleApi) Update(c *gin.Context) {
	var role systemReq.RoleFormRequest
	if err := c.ShouldBindJSON(&role); err != nil {
		response.FailWithMessage("创建失败，" + utils.GetError(err, role), c)
		return
	}

	if err := roleService.Update(role); err != nil {
		global.Error("更新失败!", err)
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetList
// @Tags Role
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/role/getList [post]
func (a *RoleApi) GetList(c *gin.Context) {

	if err, list, _ := roleService.GetList(); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(gin.H{"items": list}, "获取成功", c)
	}
}

// SetDataRole
// @Tags Role
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysRole true "设置角色资源权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /v1/system/role/setDataRole [post]
func (a *RoleApi) SetDataRole(c *gin.Context) {
	var role system.SysRole
	_ = c.ShouldBindJSON(&role)
	if err := utils.Verify(role, utils.RoleIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := roleService.SetDataRole(role); err != nil {
		global.Logger.Error("设置失败!", zap.Any("err", err))
		response.FailWithMessage("设置失败"+err.Error(), c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

// GetAll
// @Tags Role
// @Summary 获取全部角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/role/getAll [post]
func (a *RoleApi) GetAll(c *gin.Context) {

	if err, list, _ := roleService.GetAll(); err != nil {
		global.Error("获取失败!", err)
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(gin.H{"items": list}, "获取成功", c)
	}
}