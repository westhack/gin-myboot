package v1

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/request"
	"gin-myboot/modules/common/model/response"
	system "gin-myboot/modules/system/model"
	systemRes "gin-myboot/modules/system/model/response"
	"gin-myboot/utils"

	"github.com/gin-gonic/gin"
)

type SystemApiApi struct {
}


// GetAll
// @Tags SysApi
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/api/getAll [post]
func (s *SystemApiApi) GetAll(c *gin.Context) {
	if err, apis := apiService.GetAll(); err != nil {
		global.Error("获取失败!", err)
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"items": apis}, "获取成功", c)
	}
}


// GetList
// @Tags SysApi
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.QueryParams true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/api/getList [post]
func (s *SystemApiApi) GetList(c *gin.Context) {
	var searchParams request.QueryParams
	_ = c.ShouldBindJSON(&searchParams)
	if err := utils.Verify(searchParams, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, list, total := apiService.GetList(searchParams); err != nil {
		global.Error("获取失败!", err)
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
// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/system/api/create [post]
func (s *SystemApiApi) Create(c *gin.Context) {
	var api system.SysApi

	if err := c.ShouldBindJSON(&api); err != nil {
		response.FailWithMessage("创建失败，" + utils.GetError(err, api), c)
		return
	}

	if err := apiService.Create(api); err != nil {
		global.Error("创建失败!", err)
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}


// Update
// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /v1/system/api/update [post]
func (s *SystemApiApi) Update(c *gin.Context) {
	var api system.SysApi

	if err := c.ShouldBindJSON(&api); err != nil {
		response.FailWithMessage("修改失败，" + utils.GetError(err, api), c)
		return
	}

	if err := apiService.Update(api); err != nil {
		global.Error("修改失败!", err)
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}


// Delete
// @Tags SysApi
// @Summary 删除api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysApi true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/system/api/delete [post]
func (s *SystemApiApi) Delete(c *gin.Context) {
	var api system.SysApi
	_ = c.ShouldBindJSON(&api)
	if err := utils.Verify(api.Model, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := apiService.Delete(api); err != nil {
		global.Error("删除失败!", err)
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds
// @Tags SysApi
// @Summary 删除api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysApi true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/system/api/deleteByIds [post]
func (s *SystemApiApi) DeleteByIds(c *gin.Context) {
	var ids request.GetByIds
	_ = c.ShouldBindJSON(&ids)
	if err := utils.Verify(ids, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := apiService.DeleteByIds(ids.ID); err != nil {
		global.Error("删除失败!", err)
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Detail
// @Tags SysApi
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/api/detail [post]
func (s *SystemApiApi) Detail(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err, api := apiService.GetById(idInfo.ID)
	if err != nil {
		global.Error("获取失败!", err)
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(systemRes.SysAPIResponse{Api: api}, c)
	}
}

// GetRoutes
// @Tags SysApi
// @Summary 获取已注册路由列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/api/getRoutes [post]
func (s *SystemApiApi) GetRoutes(c *gin.Context) {
	routes := global.Engine.Routes()

	gin.Default().Routes()

	var list []system.SysApi
	for _, info := range routes {
		a := system.SysApi{
			Method: info.Method,
			Path: info.Path,
			Description: info.Path + ":" + info.Method,
			Value: info.Path + ":" + info.Method,
			Label: info.Path + ":" + info.Method,
		}
		list = append(list, a)
	}

	response.OkWithDetailed(gin.H{"items": list}, "获取成功", c)
}