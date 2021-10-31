package model

type SysUserRole struct {
	UserId uint64 `gorm:"column:user_id"`
	RoleId uint64 `gorm:"column:role_id"`
}

func (s *SysUserRole) TableName() string {
	return "sys_user_roles"
}
