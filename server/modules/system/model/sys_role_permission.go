package model

type SysUserRolePermission struct {
	SysPermission
	PermissionId string                  `json:"permissionId" gorm:"comment:菜单ID"`
	RoleId       string                  `json:"-" gorm:"comment:角色ID"`
	Children     []SysUserRolePermission `json:"children" gorm:"-"`
}

func (s SysUserRolePermission) TableName() string {
	return "role_permissions"
}

type SysRolePermission struct {
	RoleId       uint64 `gorm:"column:role_id;comment:角色ID"`
	PermissionId uint64 `gorm:"column:permission_id;comment:菜单ID"`
}
