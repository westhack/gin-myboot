package v1

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/response"
	system "gin-myboot/modules/system/model"
	"gin-myboot/modules/system/model/request"
	"gin-myboot/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DictDetailApi struct {
}

// Create
// @Tags SysDictDetail
// @Summary 创建SysDictDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictDetail true "SysDictDetail模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/system/dictDetail/create [post]
func (s *DictDetailApi) Create(c *gin.Context) {
	var detail system.SysDictDetail
	_ = c.ShouldBindJSON(&detail)
	if err := dictDetailService.Create(detail); err != nil {
		global.Logger.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete
// @Tags SysDictDetail
// @Summary 删除SysDictDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictDetail true "SysDictDetail模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/system/dictDetail/delete [delete]
func (s *DictDetailApi) Delete(c *gin.Context) {
	var detail system.SysDictDetail
	_ = c.ShouldBindJSON(&detail)
	if err := dictDetailService.Delete(detail); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update
// @Tags SysDictDetail
// @Summary 更新SysDictDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictDetail true "更新SysDictDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /v1/system/dictDetail/update [post]
func (s *DictDetailApi) Update(c *gin.Context) {
	var detail system.SysDictDetail
	_ = c.ShouldBindJSON(&detail)
	if err := dictDetailService.Update(&detail); err != nil {
		global.Logger.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Find
// @Tags SysDictDetail
// @Summary 用id查询SysDictDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysDictDetail true "用id查询SysDictDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /v1/system/dictDetail/find [post]
func (s *DictDetailApi) Find(c *gin.Context) {
	var detail system.SysDictDetail
	_ = c.ShouldBindQuery(&detail)
	if err := utils.Verify(detail, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, resysDictDetail := dictDetailService.GetById(detail.ID); err != nil {
		global.Logger.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"resysDictDetail": resysDictDetail}, "查询成功", c)
	}
}

// GetList
// @Tags SysDictDetail
// @Summary 分页获取SysDictDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.SysDictDetailSearch true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/dictDetail/getList [get]
func (s *DictDetailApi) GetList(c *gin.Context) {
	var pageInfo request.SysDictDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := dictDetailService.GetList(pageInfo); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
