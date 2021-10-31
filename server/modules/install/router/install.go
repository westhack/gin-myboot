package router

import (
	"gin-myboot/modules/install/api"
	"github.com/gin-gonic/gin"
)

func InitInstallRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("install")
	{
		sysRouter.POST("getSystemConfig", api.InstallApiApp.GetSystemConfig) // 获取配置文件内容
		sysRouter.POST("setSystemConfig", api.InstallApiApp.SetSystemConfig) // 设置配置文件内容
		sysRouter.POST("getServerInfo", api.InstallApiApp.GetServerInfo)     // 获取服务器信息

		sysRouter.POST("initdb", api.DBApiApp.InitDB)  // 初始化数据库
		sysRouter.POST("checkdb", api.DBApiApp.CheckDB) // 验证数据库

		sysRouter.GET("pong", api.InstallApiApp.Pong)
	}
}
