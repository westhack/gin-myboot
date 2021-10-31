package v1

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/request"
	"gin-myboot/modules/common/model/response"
	systemReq "gin-myboot/modules/system/model/request"
	"gin-myboot/utils"
	"github.com/gin-gonic/gin"
)

type PermissionApi struct {}

// GetAll
// @Tags Permission
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchPermissionParams true "查询条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/permission/getAll [post]
func (a *PermissionApi) GetAll(c *gin.Context) {
	var searchPermissionParams systemReq.SearchPermissionParams
	_ = c.ShouldBindJSON(&searchPermissionParams)

	if err, list, total := permissionService.GetList(searchPermissionParams); err != nil {
		global.Error("获取失败!", err)
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     1,
			PageSize: int(total),
		}, "获取成功", c)
	}
}

// GetTree
// @Tags Permission
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchPermissionParams true "查询条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/permission/getTree [post]
func (a *PermissionApi) GetTree(c *gin.Context) {
	var searchPermissionParams systemReq.SearchPermissionParams
	_ = c.ShouldBindJSON(&searchPermissionParams)

	if err, list := permissionService.GetTree(searchPermissionParams); err != nil {
		global.Error("获取失败!", err)
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"items":list}, "获取成功", c)
	}
}

// Update
// @Tags Permission
// @Summary 修改权限菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PermissionFormRequest true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /v1/system/permission/update [post]
func (a *PermissionApi) Update(c *gin.Context) {
	var permission systemReq.PermissionFormRequest

	if err := c.ShouldBindJSON(&permission); err != nil {
		response.FailWithMessage("修改失败，" + utils.GetError(err, permission), c)
		return
	}

	if err := permissionService.Update(&permission); err != nil {
		global.Error("修改失败!", err)

		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// Create
// @Tags Permission
// @Summary 新增权限菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PermissionFormRequest true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /v1/system/permission/create [post]
func (a *PermissionApi) Create(c *gin.Context) {
	var permission systemReq.PermissionFormRequest

	if err := c.ShouldBindJSON(&permission); err != nil {
		response.FailWithMessage("创建失败，" + utils.GetError(err, permission), c)
		return
	}

	if err := permissionService.Create(&permission); err != nil {
		global.Error("创建失败!", err)

		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete
// @Tags Permission
// @Summary 通过id删除
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "id 数据id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /v1/system/permission/delete [post]
func (a *PermissionApi) Delete(c *gin.Context) {
	var req request.GetById

	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("删除失败，" + utils.GetError(err, req), c)
		return
	}

	if req.ID == 0 {
		response.FailWithMessage("请选择一条数据", c)
		return
	}

	if err := permissionService.Delete(req.ID); err != nil {
		global.Error("删除失败!", err)

		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds
// @Tags Permission
// @Summary 批量通过id删除
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "id  数组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /v1/system/permission/deleteByIds [post]
func (a *PermissionApi) DeleteByIds(c *gin.Context) {
	var ids request.GetByIds

	if err := c.ShouldBindJSON(&ids); err != nil {
		response.FailWithMessage("删除失败，" + utils.GetError(err, ids), c)
		return
	}

	if len(ids.ID) == 0 {
		response.FailWithMessage("请选择一条数据", c)
		return
	}

	if err := permissionService.DeleteByIds(ids.ID); err != nil {
		global.Error("删除失败!", err)

		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}