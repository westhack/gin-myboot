package router

import (
	"gin-myboot/middleware"
	"gin-myboot/modules/common/api"
	"github.com/gin-gonic/gin"
)

type CaptchaRouter struct {
}

func (s *AuthRouter) InitCaptchaRouter(Router *gin.RouterGroup) {
	router := Router.Group("v1/common/captcha").Use(middleware.SystemLog())

	var api = api.ApiGroupApp.CaptchaApi
	{
		router.GET("get", api.Captcha)
		router.POST("sendSms", api.SendSms)
	}
}


