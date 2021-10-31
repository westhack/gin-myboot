package initialize

import (
	_ "gin-myboot/docs"
	"gin-myboot/global"
	"gin-myboot/middleware"
	"gin-myboot/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	global.Engine = Router
	Router.StaticFS(global.Config.Local.Path, http.Dir(global.Config.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.Logger.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.Logger.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.Logger.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	//获取路由组实例
	router.RouterGroupApp.Init(Router.Group(""))

	global.Logger.Info("router register success")
	return Router
}
