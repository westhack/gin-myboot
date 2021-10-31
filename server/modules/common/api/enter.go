package api

import "gin-myboot/modules/system/service"

type ApiGroup struct {
	AuthApi
	CaptchaApi
}

var ApiGroupApp = new(ApiGroup)

var roleService = service.SystemServiceGroup.RoleService
var jwtService = service.SystemServiceGroup.JwtService
var userService = service.SystemServiceGroup.UserService
var permissionMenuService = service.SystemServiceGroup.PermissionMenuService
var permissionService = service.SystemServiceGroup.PermissionService
