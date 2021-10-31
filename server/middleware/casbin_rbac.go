package middleware

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/response"
	"gin-myboot/modules/system/model/request"
	"gin-myboot/modules/system/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

var casbinService = service.SystemServiceGroup.CasbinService

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*request.CustomClaims)
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := strconv.FormatUint(waitUse.RoleId, 10)
		e := casbinService.Casbin()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		global.Error("===========> sub ", sub)
		global.Error("===========> obj ", obj)
		global.Error("===========> act ", act)
		global.Error("===========> success ", success)
		if global.Config.System.Env == "develop1" || success {
			c.Next()
		} else {
			response.UnauthenticatedWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
	}
}
