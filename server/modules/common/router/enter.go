package router

import (
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	AuthRouter
}

var RouterGroupApp = new(RouterGroup)

func InitCommonRouter(Router *gin.RouterGroup)  {
	RouterGroupApp.InitAuthRouter(Router)    // 用户登录、注册相关
	RouterGroupApp.InitCaptchaRouter(Router) // 图形验证码、手机验证码
}