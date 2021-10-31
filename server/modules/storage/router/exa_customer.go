package router

import (
	"gin-myboot/middleware"
	"gin-myboot/modules/storage/api"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct {
}

func (e *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("customer").Use(middleware.SystemLog())
	var exaCustomerApi = api.StorageApiGroupApp.CustomerApi
	{
		customerRouter.POST("customer", exaCustomerApi.CreateExaCustomer)     // 创建客户
		customerRouter.PUT("customer", exaCustomerApi.UpdateExaCustomer)      // 更新客户
		customerRouter.DELETE("customer", exaCustomerApi.DeleteExaCustomer)   // 删除客户
		customerRouter.GET("customer", exaCustomerApi.GetExaCustomer)         // 获取单一客户信息
		customerRouter.GET("customerList", exaCustomerApi.GetExaCustomerList) // 获取客户列表
	}
}
