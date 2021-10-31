package model

import (
	"gin-myboot/global"
)

type SysDepartmentHeader struct {
	global.Model
	DepartmentId uint64    `json:"departmentId" gorm:"comment:部门ID"`
	UserId       uint64    `json:"userId" gorm:"comment:用户ID"`
	Type         uint    `json:"type" gorm:"comment:负责人类型 默认0主要 1副职"`
	User         SysUser `json:"user" gorm:"foreignKey:user_id;references:id;comment:用户"`
}
