package api

import (
    "gin-myboot/global"
    "gin-myboot/modules/{{.ModuleName}}/model"
    "gin-myboot/modules/common/model/request"
    "gin-myboot/modules/common/model/response"
    "gin-myboot/modules/{{.ModuleName}}/service"
    "gin-myboot/utils"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type {{.StructName}}Api struct {
}

var {{.StructName}}ApiApp = new({{.StructName}}Api)

var {{.Abbreviation}}Service = service.{{.StructName}}ServiceApp


// Create 创建{{.StructName}}
// @Tags {{.StructName}}
// @Summary 创建{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "创建{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /{{.ModuleName}}/{{.Abbreviation}}/create [post]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Create(c *gin.Context) {
    var {{.Abbreviation}} model.{{.StructName}}
    _ = c.ShouldBindJSON(&{{.Abbreviation}})
    if err := {{.Abbreviation}}Service.Create({{.Abbreviation}}); err != nil {
        global.Logger.Error("创建失败!", zap.Any("err", err))
        response.FailWithMessage("创建失败", c)
    } else {
        response.OkWithMessage("创建成功", c)
    }
}

// Delete 删除{{.StructName}}
// @Tags {{.StructName}}
// @Summary 删除{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /{{.ModuleName}}/{{.Abbreviation}}/delete [post]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Delete(c *gin.Context) {
    var id request.GetById
    _ = c.ShouldBindJSON(&id)
    if err := {{.Abbreviation}}Service.Delete(id.ID); err != nil {
        global.Logger.Error("删除失败!", zap.Any("err", err))
        response.FailWithMessage("删除失败", c)
    } else {
        response.OkWithMessage("删除成功", c)
    }
}

// DeleteByIds 批量删除{{.StructName}}
// @Tags {{.StructName}}
// @Summary 批量删除{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "批量删除{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /{{.ModuleName}}/{{.Abbreviation}}/deleteByIds [post]
func ({{.Abbreviation}}Api *{{.StructName}}Api) DeleteByIds(c *gin.Context) {
    var ids request.GetByIds
    _ = c.ShouldBindJSON(&ids)
    if err := {{.Abbreviation}}Service.DeleteByIds(ids.ID); err != nil {
        global.Logger.Error("批量删除失败!", zap.Any("err", err))
        response.FailWithMessage("批量删除失败", c)
    } else {
        response.OkWithMessage("批量删除成功", c)
    }
}

// Update 更新{{.StructName}}
// @Tags {{.StructName}}
// @Summary 更新{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.{{.StructName}} true "更新{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /{{.ModuleName}}/{{.Abbreviation}}/update [post]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Update(c *gin.Context) {
    var {{.Abbreviation}} model.{{.StructName}}
    _ = c.ShouldBindJSON(&{{.Abbreviation}})
    if err := {{.Abbreviation}}Service.Update({{.Abbreviation}}); err != nil {
        global.Logger.Error("更新失败!", zap.Any("err", err))
        response.FailWithMessage("更新失败", c)
    } else {
        response.OkWithMessage("更新成功", c)
    }
}

// Find 用id查询{{.StructName}}
// @Tags {{.StructName}}
// @Summary 用id查询{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetById true "用id查询{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /{{.ModuleName}}/{{.Abbreviation}}/find [get]
func ({{.Abbreviation}}Api *{{.StructName}}Api) Find(c *gin.Context) {
    var id request.GetById
    _ = c.ShouldBindQuery(&id)
    if err, re{{.Abbreviation}} := {{.Abbreviation}}Service.FindById(id.ID); err != nil {
        global.Logger.Error("查询失败!", zap.Any("err", err))
        response.FailWithMessage("查询失败", c)
    } else {
        response.OkWithData(gin.H{"re{{.Abbreviation}}": re{{.Abbreviation}}}, c)
    }
}

// GetList 分页获取{{.StructName}}列表
// @Tags {{.StructName}}
// @Summary 分页获取{{.StructName}}列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchParams true "分页获取{{.StructName}}列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /{{.ModuleName}}/{{.Abbreviation}}/getList [post]
func ({{.Abbreviation}}Api *{{.StructName}}Api) GetList(c *gin.Context) {
    var searchParams request.QueryParams
    _ = c.ShouldBindJSON(&searchParams)
    if err := utils.Verify(searchParams, utils.PageInfoVerify); err != nil {
    	response.FailWithMessage(err.Error(), c)
    	return
    }
    if err, list, total := {{.Abbreviation}}Service.GetList(searchParams); err != nil {
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
