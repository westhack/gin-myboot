package v1

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/response"
	system "gin-myboot/modules/system/model"
	"gin-myboot/modules/system/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DictApi struct {
}

// Create 
// @Tags SysDict
// @Summary 创建SysDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDict true "SysDict模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/system/dict/create [post]
func (s *DictApi) Create(c *gin.Context) {
	var dict system.SysDict
	_ = c.ShouldBindJSON(&dict)
	if err := dictService.Create(dict); err != nil {
		global.Logger.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败" + err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 
// @Tags SysDict
// @Summary 删除SysDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDict true "SysDict模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/system/dict/delete [post]
func (s *DictApi) Delete(c *gin.Context) {
	var dict system.SysDict
	_ = c.ShouldBindJSON(&dict)
	if err := dictService.Delete(dict); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败" + err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update 
// @Tags SysDict
// @Summary 更新SysDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDict true "SysDict模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /v1/system/dict/update [put]
func (s *DictApi) Update(c *gin.Context) {
	var dict system.SysDict
	_ = c.ShouldBindJSON(&dict)
	if err := dictService.Update(&dict); err != nil {
		global.Logger.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败" + err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Find
// @Tags SysDict
// @Summary 用id查询SysDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysDict true "ID或字典英名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /v1/system/dict/find [get]
func (s *DictApi) Find(c *gin.Context) {
	var dict system.SysDict
	_ = c.ShouldBindQuery(&dict)
	if err, sysDict := dictService.GetByIdOrType(dict.Type, dict.ID); err != nil {
		global.Logger.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"dict": sysDict}, "查询成功", c)
	}
}

// GetList 
// @Tags SysDict
// @Summary 分页获取SysDict列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SysDictSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/dict/getList [post]
func (s *DictApi) GetList(c *gin.Context) {
	var pageInfo request.SysDictSearch
	_ = c.ShouldBindJSON(&pageInfo)

	if err, list, total := dictService.GetList(pageInfo); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     1,
			PageSize: 1,
		}, "获取成功", c)
	}
}


// GetAll
// @Tags SysDict
// @Summary 分页获取SysDict列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/dict/getAll [post]
func (s *DictApi) GetAll(c *gin.Context) {

	if err, list := dictService.GetAll(); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"items": list}, "获取成功", c)
	}
}

// SaveDetail
// @Tags SysDict
// @Summary 分页获取SysDict列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SaveDetailRequest true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"保存成功"}"
// @Router /v1/system/dict/saveDetail [post]
func (s *DictApi) SaveDetail(c *gin.Context) {
	var saveDetails request.SaveDetailRequest
	_ = c.ShouldBindJSON(&saveDetails)

	if err := dictService.SaveDetail(&saveDetails); err != nil {
		global.Error("保存失败!", err)
		response.FailWithMessage("保存失败" + err.Error(), c)
	} else {
		response.OkWithMessage( "保存成功", c)
	}
}