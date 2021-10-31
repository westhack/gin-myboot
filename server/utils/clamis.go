package utils

import (
	"errors"
	"gin-myboot/global"
	"gin-myboot/modules/common/model/response"
	system "gin-myboot/modules/system/model"
	systemReq "gin-myboot/modules/system/model/request"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) uint64 {
	if claims, exists := c.Get("claims"); !exists {
		global.Logger.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件!")
		return 0
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.ID
	}
}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		global.Logger.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件!")
		return uuid.UUID{}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.UUID
	}
}

// GetUserRoleId 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserRoleId(c *gin.Context) uint64 {
	if claims, exists := c.Get("claims"); !exists {
		global.Logger.Error("请检查路由是否使用jwt中间件!")
		return 0
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.RoleId
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) *systemReq.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		global.Logger.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件!")
		return nil
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse
	}
}

// GetUserRoleName 从Gin的Context中获取从jwt解析出来的用户角色name
func GetUserRoleName(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		global.Logger.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件!")
		return ""
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.RoleName
	}
}


// GetCurrUser 从Gin的Context中获取从jwt解析出来的用户ID，获取用户信息
func GetCurrUser(c *gin.Context) (err error, user system.SysUser) {
	if claims, exists := c.Get("claims"); !exists {
		global.Logger.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件!", zap.Any("cc", claims))

		response.UnauthorizedWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
		return errors.New("未登录或非法访问"), user
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		userId := waitUse.ID
		err = global.GormDB.Preload("Roles").Preload("Role").First(&user, "id = ?", userId).Error

		return err, user
	}
}


// GetUser 从Gin的Context中获取从jwt解析出来的用户ID，获取用户信息
func GetUser(c *gin.Context) (err error, user system.SysUser) {
	if claims, exists := c.Get("claims"); !exists {
		global.Logger.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件!")
		return errors.New("用户不存在"), user
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		userId := waitUse.ID
		err = global.GormDB.Preload("Roles").Preload("Role").First(&user, "id = ?", userId).Error

		return err, user
	}
}