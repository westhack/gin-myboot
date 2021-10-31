package model

import (
	"gin-myboot/global"
)

type SysDepartment struct {
	global.Model
	ParentId        uint64              `json:"parentId" gorm:"comment:父菜单ID"` // 父菜单ID
	SortOrder       float64             `json:"sortOrder" gorm:"comment:排序标记"`
	Status          bool                `json:"status" gorm:"comment:是否启用"`
	IsParent        bool                `json:"isParent" gorm:"comment:是否为父节点(含子节点) 默认false"`
	Children        []SysDepartment     `json:"children" gorm:"-"`
	Title           string              `json:"title" gorm:"comment:名称"`
	RoleDepartments []SysRoleDepartment `json:"roleDepartments" gorm:"foreignKey:department_id;references:id"`
	//MainHeaders     []SysUser           `json:"mainHeaders" gorm:"many2many:sys_department_headers;foreignKey:id;joinForeignKey:department_id;References:id;joinReferences:user_id"`
	//ViceHeaders     []SysUser           `json:"viceHeaders" gorm:"many2many:sys_department_headers;foreignKey:id;joinForeignKey:department_id;References:id;joinReferences:user_id"`
	MainHeaders []SysUser `json:"mainHeaders" gorm:"-"`
	ViceHeaders []SysUser `json:"viceHeaders" gorm:"-"`
	MainHeader  []uint64  `json:"mainHeader" gorm:"-"`
	ViceHeader  []uint64  `json:"viceHeader" gorm:"-"`
}
