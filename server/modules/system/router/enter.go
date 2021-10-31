package router

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	EmailRouter
	JwtRouter
	SystemRouter
	RedisRouter
}

var SystemRouterGroupApp = new(RouterGroup)

func InitSystemRouter(Router *gin.RouterGroup) {
	SystemRouterGroupApp.InitEmailRouter(Router)  // 邮件相关路由
	SystemRouterGroupApp.InitSystemRouter(Router) // 系统业务
	SystemRouterGroupApp.InitJwtRouter(Router)    // jwt相关路由
	SystemRouterGroupApp.InitRedisRouter(Router)  // redis相关路由
}
