package router

import (
	"gin-myboot/middleware"
	"gin-myboot/modules/captcha/api"
	"github.com/gin-gonic/gin"
)

func InitCaptchaRouter(Router *gin.RouterGroup) {
	router := Router.Group("captcha").Use(middleware.SystemLog())

	var api = api.CaptchaApiApp
	{
		router.POST("get", api.Get)
		router.POST("check", api.Check)
		router.POST("verification", api.Verification)
	}
}
