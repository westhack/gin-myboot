package service

type ServiceGroup struct {
	JwtService
	ApiService
	RoleService
	CasbinService
	DictService
	DictDetailService
	EmailService
	LogService
	SystemConfigService
	UserService
	PermissionMenuService
	PermissionService
	AdminService
	ConfigService
	DepartmentService
	MessageService
	RedisService
}

var SystemServiceGroup = new(ServiceGroup)