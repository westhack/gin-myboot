package v1

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/response"
	system "gin-myboot/modules/system/model"
	"gin-myboot/modules/system/model/request"
	"github.com/gin-gonic/gin"
)

type ConfigApi struct {
}

// Create
// @Tags Config
// @Summary 创建Config
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysConfig true "Config模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/system/config/create [post]
func (s *ConfigApi) Create(c *gin.Context) {
	var config system.SysConfig
	_ = c.ShouldBindJSON(&config)
	if err := configService.Create(config); err != nil {
		global.Error("创建失败!", err)
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete
// @Tags Config
// @Summary 删除Config
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysConfig true "Config模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/system/config/delete [post]
func (s *ConfigApi) Delete(c *gin.Context) {
	var config system.SysConfig
	_ = c.ShouldBindJSON(&config)
	if err := configService.Delete(config); err != nil {
		global.Error("删除失败!", err)
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// Update
// @Tags Config
// @Summary 更新Config
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysConfig true "Config模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /v1/system/config/update [post]
func (s *ConfigApi) Update(c *gin.Context) {
	var config system.SysConfig
	_ = c.ShouldBindJSON(&config)
	if err := configService.Update(&config); err != nil {
		global.Error("更新失败!", err)
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Find
// @Tags Config
// @Summary 用id查询Config
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysConfig true "ID或字典英名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /v1/system/config/find [get]
func (s *ConfigApi) Find(c *gin.Context) {
	var config system.SysConfig
	_ = c.ShouldBindQuery(&config)
	if err, data := configService.GetConfig(config.Type, config.ID); err != nil {
		global.Error("查询失败!", err)
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"config": data}, "查询成功", c)
	}
}

// GetList
// @Tags Config
// @Summary 分页获取Config列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.ConfigSearch true "搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/config/getList [post]
func (s *ConfigApi) GetList(c *gin.Context) {
	var pageInfo request.ConfigSearch
	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := configService.GetList(pageInfo); err != nil {
		global.Error("获取失败!", err)
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     1,
			PageSize: 100,
		}, "获取成功", c)
	}
}

// GetAll
// @Tags Config
// @Summary 分页获取Config列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.ConfigSearch true "搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/system/config/getAll [post]
func (s *ConfigApi) GetAll(c *gin.Context) {
	var pageInfo request.ConfigSearch
	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := configService.GetAll(pageInfo); err != nil {
		global.Error("获取失败!", err)
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     1,
			PageSize: 100,
		}, "获取成功", c)
	}
}

// SetValue
// @Tags Config
// @Summary 更新Config
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysConfig true "Config模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /v1/system/config/setValue [post]
func (s *ConfigApi) SetValue(c *gin.Context) {
	var req map[string]string
	_ = c.ShouldBindJSON(&req)
	if err := configService.SetValue(req); err != nil {
		global.Error("更新失败!", err)
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// Write
// @Tags Config
// @Summary 更新Config
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysConfig true "Config模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /v1/system/config/write [post]
func (s *ConfigApi) Write(c *gin.Context) {
	var req map[string]interface{}
	_ = c.ShouldBindJSON(&req)
	if err := configService.Write(req); err != nil {
		global.Error("更新失败!", err)
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}