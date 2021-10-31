package router

import (
	"gin-myboot/middleware"
	"gin-myboot/modules/common/api"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func (s *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {

	var commonApi = api.ApiGroupApp.AuthApi

	demo := Router.Group("v1/common/demo").Use(middleware.SystemLog())//.Use(middleware.ImageValidate())
	a := api.NewDemoApi()
	{
		demo.POST("getList", a.GetList)
		demo.POST("getAll", a.GetAll)
		demo.POST("delete", a.Delete)
		demo.POST("create", a.Create)
		demo.POST("update", a.Update)
	}

	noLoginRouter := Router.Group("v1/common/auth").Use(middleware.SystemLog())
	{
		noLoginRouter.POST("login", commonApi.Login)
		noLoginRouter.Use(middleware.SmsCodeValidate()).POST("register", commonApi.Register)
	}

	userRouter := Router.Group("v1/common/auth").Use(middleware.JWTAuth()).Use(middleware.SystemLog()).Use(middleware.CasbinHandler())
	{
		userRouter.POST("logout", commonApi.Logout)
		userRouter.POST("getUserInfo", commonApi.GetUserInfo)
		userRouter.POST("getUserMenus", commonApi.GetUserMenus)
		userRouter.POST("changePassword", commonApi.ChangePassword)
		userRouter.POST("update", commonApi.Update)
	}
}


