package v1

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/request"
	"gin-myboot/modules/common/model/response"
	system "gin-myboot/modules/system/model"
	"gin-myboot/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserLogApi struct {
}

// Create
// @Tags SysLog
// @Summary 创建SysLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysLog true "创建SysLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/log/create [post]
func (s *UserLogApi) Create(c *gin.Context) {
	var sysUserLog system.SysLog
	_ = c.ShouldBindJSON(&sysUserLog)
	if err := userLogService.Create(sysUserLog); err != nil {
		global.Logger.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete
// @Tags SysLog
// @Summary 删除SysLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysLog true "SysLog模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/system/log/delete [post]
func (s *UserLogApi) Delete(c *gin.Context) {
	var id request.GetById
	_ = c.ShouldBindJSON(&id)
	if err := userLogService.Delete(id.ID); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds
// @Tags SysLog
// @Summary 批量删除SysLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "批量删除SysLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /v1/system/log/deleteByIds [post]
func (s *UserLogApi) DeleteByIds(c *gin.Context) {
	var ids request.GetByIds
	_ = c.ShouldBindJSON(&ids)
	if err := userLogService.DeleteByIds(ids.ID); err != nil {
		global.Logger.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Find
// @Tags SysLog
// @Summary 用id查询SysLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetById true "Id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /v1/system/log/find [post]
func (s *UserLogApi) Find(c *gin.Context) {
	var id request.GetById
	_ = c.ShouldBindQuery(&id)
	if err := utils.Verify(id, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, resysUserLog := userLogService.GetById(id.ID); err != nil {
		global.Logger.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"resysUserLog": resysUserLog}, "查询成功", c)
	}
}

// GetList
// @Tags SysLog
// @Summary 分页获取SysLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.QueryParams true "页码, 每页大小, 搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/log/getList [get]
func (s *UserLogApi) GetList(c *gin.Context) {
	var searchParams request.QueryParams
	_ = c.ShouldBindJSON(&searchParams)
	if err := utils.Verify(searchParams, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := userLogService.GetList(searchParams); err != nil {
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



// DeleteAll
// @Tags SysLog
// @Summary 删除SysLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Empty true "{}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/system/log/deleteAll [post]
func (s *UserLogApi) DeleteAll(c *gin.Context) {
	if err := userLogService.DeleteAll(); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}