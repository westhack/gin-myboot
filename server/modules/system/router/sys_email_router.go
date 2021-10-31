package router

import (
	"gin-myboot/middleware"
	v1 "gin-myboot/modules/system/api/v1"
	"github.com/gin-gonic/gin"
)

type EmailRouter struct {
}

func (s *EmailRouter) InitEmailRouter(Router *gin.RouterGroup) {
	emailRouter := Router.Group("v1/system/email").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.SystemLog())
	var systemApi = v1.ApiGroupApp.SystemApi
	{
		emailRouter.POST("emailTest", systemApi.EmailTest) // 发送测试邮件
	}
}
