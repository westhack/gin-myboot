package model

type SysRoleDepartment struct {
	DepartmentId uint64 `gorm:"column:department_id"`
	RoleId       uint64 `gorm:"column:role_id"`
}
