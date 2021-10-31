package router

import (
    "gin-myboot/middleware"
    "gin-myboot/modules/article/api"
    "github.com/gin-gonic/gin"
)

type SysArticleRouter struct {
}

// InitSysArticleRouter 初始化 SysArticle 路由信息
func InitSysArticleRouter(Router *gin.RouterGroup) {
    articleRouter := Router.Group("article/article").Use(middleware.SystemLog())
    var articleApi = api.SysArticleApiApp
    {
        articleRouter.POST("create", articleApi.Create)   // 新建SysArticle
        articleRouter.POST("delete", articleApi.Delete) // 删除SysArticle
        articleRouter.POST("deleteByIds", articleApi.DeleteByIds) // 批量删除SysArticle
        articleRouter.POST("update", articleApi.Update)    // 更新SysArticle
        articleRouter.GET("find", articleApi.Find)        // 根据ID获取SysArticle
        articleRouter.POST("getList", articleApi.GetList)  // 获取SysArticle列表
    }
}
