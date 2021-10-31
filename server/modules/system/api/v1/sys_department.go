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

type DepartmentApi struct {

}


// GetList
// @Tags Department
// @Summary 获取部门列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchDepartmentParams true "部门名称"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/department/getList [post]
func (api *DepartmentApi) GetList(c *gin.Context)  {
	var req systemReq.SearchDepartmentParams
	_ = c.ShouldBindJSON(&req)

	if err, list, _ := departmentService.GetList(req); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"items": list}, "获取成功", c)
	}
}


// Create
// @Tags Department
// @Summary 创建部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDepartment true "部门名称,..."
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/system/department/create [post]
func (api *DepartmentApi) Create(c *gin.Context)  {
	var req system.SysDepartment
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("创建失败，" + utils.GetError(err, req), c)
		return
	}

	if err := departmentService.Create(&req); err != nil {
		global.Logger.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithDetailed(req, "创建成功", c)
	}
}

// Update
// @Tags Department
// @Summary 更新部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDepartment true "部门名称,..."
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/department/update [post]
func (api *DepartmentApi) Update(c *gin.Context)  {
	var req system.SysDepartment
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("修改失败，" + utils.GetError(err, req), c)
		return
	}
	if err := departmentService.Update(&req); err != nil {
		global.Logger.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败 " + err.Error(), c)
	} else {
		response.OkWithDetailed(req, "修改成功", c)
	}
}

// Delete
// @Tags Department
// @Summary 删除部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "部门ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/system/department/delete [post]
func (api *DepartmentApi) Delete(c *gin.Context) {
	var id request.GetById
	_ = c.ShouldBindJSON(&id)

	if err := departmentService.Delete(id.ID); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}


// DeleteByIds
// @Tags Department
// @Summary 批量删除部门
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "部门IDs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/system/department/delete [post]
func (api *DepartmentApi) DeleteByIds(c *gin.Context) {
	var ids request.GetByIds
	_ = c.ShouldBindJSON(&ids)

	if err := departmentService.DeleteByIds(ids.ID); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}


// GetAll
// @Tags Department
// @Summary 获取可用部门列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/department/getAll [post]
func (api *DepartmentApi) GetAll(c *gin.Context)  {
	if err, list, _ := departmentService.GetAll(); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"items": list}, "获取成功", c)
	}
}