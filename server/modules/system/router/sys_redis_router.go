package router

import (
	"gin-myboot/middleware"
	v1 "gin-myboot/modules/system/api/v1"
	"github.com/gin-gonic/gin"
)

type RedisRouter struct {
}

func (s *RedisRouter) InitRedisRouter(Router *gin.RouterGroup) {
	redisRouter := Router.Group("v1/system/redis").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.SystemLog())
	var redisApi = v1.ApiGroupApp.RedisApi
	{
		redisRouter.POST("getList", redisApi.GetList)
		redisRouter.POST("create", redisApi.Create)
		redisRouter.POST("update", redisApi.Update)
		redisRouter.POST("delete", redisApi.Delete)
		redisRouter.POST("deleteByIds", redisApi.DeleteByIds)
		redisRouter.GET("find", redisApi.DeleteByIds)
	}
}
