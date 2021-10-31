package model

type SysRolePermissionMenu struct {
	SysPermission
	PermissionId uint64                  `json:"permissionId" gorm:"comment:菜单ID"`
	RoleId       uint64                  `json:"-" gorm:"comment:角色ID"`
	Children     []SysRolePermissionMenu `json:"children" gorm:"-"`
}

func (s SysRolePermissionMenu) TableName() string {
	return "role_permission_menus"
}
