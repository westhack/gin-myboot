package model

import (
	"gin-myboot/global"
)

type SysPermission struct {
	global.Model
	Level       uint            `json:"-" gorm:"comment:层级"`
	ParentId    uint64          `json:"parentId" gorm:"comment:父菜单ID"`     // 父菜单ID
	Path        string          `json:"path" gorm:"comment:路由path;type:varchar(255);"`        // 路由path
	Name        string          `json:"name" gorm:"comment:路由name"`        // 路由name
	Hidden      bool            `json:"hidden" gorm:"comment:是否在列表隐藏"`     // 是否在列表隐藏
	Component   string          `json:"component" gorm:"comment:对应前端文件路径;type:varchar(255);"` // 对应前端文件路径
	SortOrder   float64         `json:"sortOrder" gorm:"comment:排序标记"`
	DefaultMenu bool            `json:"defaultMenu" gorm:"comment:是否是默认菜单"`
	Status      bool            `json:"status" gorm:"comment:状态"`
	IsButton    bool            `json:"isButton" gorm:"comment:是否权限按钮"`
	Redirect    string          `json:"redirect" gorm:"comment:跳转地址;type:varchar(255);"`
	Api         string          `json:"api" gorm:"comment:API权限;type:text;"`
	SysRoles    []SysRole       `json:"roles" gorm:"many2many:sys_role_permissions;foreignKey:id;joinForeignKey:permission_id;References:id;joinReferences:role_id"`
	Children    []SysPermission `json:"children" gorm:"-"`
	KeepAlive   bool            `json:"keepAlive" gorm:"comment:是否缓存"` // 是否缓存
	Title       string          `json:"title" gorm:"comment:菜单标题"`    // 菜单名
	Icon        string          `json:"icon" gorm:"comment:菜单图标"`      // 菜单图标
}
