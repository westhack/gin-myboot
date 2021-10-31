package v1

import "gin-myboot/modules/system/service"

type ApiGroup struct {
	SystemApiApi
	RoleApi
	CasbinApi
	DictApi
	DictDetailApi
	SystemApi
	JwtApi
	UserLogApi
	PermissionApi
	DepartmentApi
	UserApi
	AdminApi
	ConfigApi
	MessageApi
	RedisApi
}

var ApiGroupApp = new(ApiGroup)

var roleService = service.SystemServiceGroup.RoleService
var apiService = service.SystemServiceGroup.ApiService
var casbinService = service.SystemServiceGroup.CasbinService
var dictService = service.SystemServiceGroup.DictService
var dictDetailService = service.SystemServiceGroup.DictDetailService
var emailService = service.SystemServiceGroup.EmailService
var jwtService = service.SystemServiceGroup.JwtService
var userLogService = service.SystemServiceGroup.LogService
var userService = service.SystemServiceGroup.UserService
var systemConfigService = service.SystemServiceGroup.SystemConfigService
var permissionMenuService = service.SystemServiceGroup.PermissionMenuService
var permissionService = service.SystemServiceGroup.PermissionService
var configService = service.SystemServiceGroup.ConfigService
var departmentService = service.SystemServiceGroup.DepartmentService
var messageService = service.SystemServiceGroup.MessageService
var adminService = service.SystemServiceGroup.AdminService
var redisService = service.SystemServiceGroup.RedisService