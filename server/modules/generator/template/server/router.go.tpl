package router

import (
    "gin-myboot/modules/{{.ModuleName}}/api"
    "gin-myboot/middleware"
    "github.com/gin-gonic/gin"
)

type {{.StructName}}Router struct {
}

// Init{{.StructName}}Router 初始化 {{.StructName}} 路由信息
func Init{{.StructName}}Router(Router *gin.RouterGroup) {
    {{.Abbreviation}}Router := Router.Group("{{.ModuleName}}/{{.Abbreviation}}").Use(middleware.SystemLog())
    var {{.Abbreviation}}Api = api.{{.StructName}}ApiApp
    {
        {{.Abbreviation}}Router.POST("create", {{.Abbreviation}}Api.Create)   // 新建{{.StructName}}
        {{.Abbreviation}}Router.POST("delete", {{.Abbreviation}}Api.Delete) // 删除{{.StructName}}
        {{.Abbreviation}}Router.POST("deleteByIds", {{.Abbreviation}}Api.DeleteByIds) // 批量删除{{.StructName}}
        {{.Abbreviation}}Router.POST("update", {{.Abbreviation}}Api.Update)    // 更新{{.StructName}}
        {{.Abbreviation}}Router.GET("find", {{.Abbreviation}}Api.Find)        // 根据ID获取{{.StructName}}
        {{.Abbreviation}}Router.POST("getList", {{.Abbreviation}}Api.GetList)  // 获取{{.StructName}}列表
    }
}
