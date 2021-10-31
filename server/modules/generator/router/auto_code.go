package router

import (
	"gin-myboot/modules/generator/api"
	"github.com/gin-gonic/gin"
)

func InitAutoCodeRouter(Router *gin.RouterGroup) {
	autoCodeRouter := Router.Group("generator/autoCode")
	var roleApi = api.AutoCodeApiApp
	{
		autoCodeRouter.POST("delete", roleApi.Delete) // 删除回滚记录
		autoCodeRouter.POST("deleteByIds", roleApi.DeleteByIds) // 删除回滚记录
		autoCodeRouter.POST("getMeta", roleApi.GetMeta)             // 根据id获取meta信息
		autoCodeRouter.POST("getList", roleApi.GetList) // 获取回滚记录分页
		autoCodeRouter.POST("rollback", roleApi.RollBack)           // 回滚
		autoCodeRouter.POST("preview", roleApi.PreviewTemp)         // 获取自动创建代码预览
		autoCodeRouter.POST("createTemp", roleApi.CreateTemp)       // 创建自动化代码
		autoCodeRouter.GET("getTables", roleApi.GetTables)          // 获取对应数据库的表
		autoCodeRouter.GET("getDatabases", roleApi.GetDatabases)    // 获取数据库
		autoCodeRouter.GET("getColumns", roleApi.GetColumns)        // 获取指定表所有字段信息
		autoCodeRouter.GET("getModules", roleApi.GetModules)        // 获取指定表所有字段信息
	}
}
