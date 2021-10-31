package router

import (
	"gin-myboot/middleware"
	v1 "gin-myboot/modules/system/api/v1"
	"github.com/gin-gonic/gin"
)

type JwtRouter struct {
}

func (s *JwtRouter) InitJwtRouter(Router *gin.RouterGroup) {
	jwtRouter := Router.Group("v1/system/jwt").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.SystemLog())
	var jwtApi = v1.ApiGroupApp.JwtApi
	{
		jwtRouter.POST("jsonInBlacklist", jwtApi.JsonInBlacklist) // jwt加入黑名单
	}
}
