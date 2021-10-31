package api

import (
    "gin-myboot/global"
    "gin-myboot/modules/article/model"
    "gin-myboot/modules/article/service"
    "gin-myboot/modules/common/model/request"
    "gin-myboot/modules/common/model/response"
    "gin-myboot/utils"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type SysArticleApi struct {
}

var SysArticleApiApp = new(SysArticleApi)

var articleService = service.SysArticleServiceApp


// Create 创建SysArticle
// @Tags SysArticle
// @Summary 创建SysArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysArticle true "创建SysArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /article/article/create [post]
func (articleApi *SysArticleApi) Create(c *gin.Context) {
    var article model.SysArticle
    _ = c.ShouldBindJSON(&article)
    if err := articleService.Create(article); err != nil {
        global.Logger.Error("创建失败!", zap.Any("err", err))
        response.FailWithMessage("创建失败", c)
    } else {
        response.OkWithMessage("创建成功", c)
    }
}

// Delete 删除SysArticle
// @Tags SysArticle
// @Summary 删除SysArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除SysArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /article/article/delete [post]
func (articleApi *SysArticleApi) Delete(c *gin.Context) {
    var id request.GetById
    _ = c.ShouldBindJSON(&id)
    if err := articleService.Delete(id.ID); err != nil {
        global.Logger.Error("删除失败!", zap.Any("err", err))
        response.FailWithMessage("删除失败", c)
    } else {
        response.OkWithMessage("删除成功", c)
    }
}

// DeleteByIds 批量删除SysArticle
// @Tags SysArticle
// @Summary 批量删除SysArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "批量删除SysArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /article/article/deleteByIds [post]
func (articleApi *SysArticleApi) DeleteByIds(c *gin.Context) {
    var ids request.GetByIds
    _ = c.ShouldBindJSON(&ids)
    if err := articleService.DeleteByIds(ids.ID); err != nil {
        global.Logger.Error("批量删除失败!", zap.Any("err", err))
        response.FailWithMessage("批量删除失败", c)
    } else {
        response.OkWithMessage("批量删除成功", c)
    }
}

// Update 更新SysArticle
// @Tags SysArticle
// @Summary 更新SysArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysArticle true "更新SysArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /article/article/update [post]
func (articleApi *SysArticleApi) Update(c *gin.Context) {
    var article model.SysArticle
    _ = c.ShouldBindJSON(&article)
    if err := articleService.Update(article); err != nil {
        global.Logger.Error("更新失败!", zap.Any("err", err))
        response.FailWithMessage("更新失败", c)
    } else {
        response.OkWithMessage("更新成功", c)
    }
}

// Find 用id查询SysArticle
// @Tags SysArticle
// @Summary 用id查询SysArticle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.GetById true "用id查询SysArticle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /article/article/find [get]
func (articleApi *SysArticleApi) Find(c *gin.Context) {
    var id request.GetById
    _ = c.ShouldBindQuery(&id)
    if err, rearticle := articleService.FindById(id.ID); err != nil {
        global.Logger.Error("查询失败!", zap.Any("err", err))
        response.FailWithMessage("查询失败", c)
    } else {
        response.OkWithData(gin.H{"rearticle": rearticle}, c)
    }
}

// GetList 分页获取SysArticle列表
// @Tags SysArticle
// @Summary 分页获取SysArticle列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchParams true "分页获取SysArticle列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /article/article/getList [post]
func (articleApi *SysArticleApi) GetList(c *gin.Context) {
    var searchParams request.QueryParams
    _ = c.ShouldBindJSON(&searchParams)
    if err := utils.Verify(searchParams, utils.PageInfoVerify); err != nil {
    	response.FailWithMessage(err.Error(), c)
    	return
    }
    if err, list, total := articleService.GetList(searchParams); err != nil {
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
