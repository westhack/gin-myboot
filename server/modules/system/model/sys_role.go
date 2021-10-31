package model

import (
	"gin-myboot/global"
	"time"
)

type SysRole struct {
	ID              uint64              `json:"id" gorm:"primarykey"` // 主键ID
	CreatedAt       global.LocalTime    `json:"createdAt"`            // 创建时间
	UpdatedAt       time.Time           `json:"updatedAt"`            // 更新时间
	DeletedAt       *time.Time          `json:"-" sql:"index"`
	Title           string              `json:"title" gorm:"comment:名称"`       // 角色名
	Name            string              `json:"name" gorm:"comment:角色名"`       // 角色名
	ParentId        uint64              `json:"parentId" gorm:"comment:父角色ID"` // 父角色ID
	Children        []SysRole           `json:"children" gorm:"-"`
	Permissions     []SysPermission     `json:"permissions" gorm:"many2many:sys_role_permissions;foreignKey:id;joinForeignKey:role_id;References:id;joinReferences:permission_id"`
	Departments     []SysDepartment     `json:"departments" gorm:"many2many:sys_role_departments;foreignKey:id;joinForeignKey:role_id;References:id;joinReferences:department_id"`
	RoleDepartments []SysRoleDepartment `json:"roleDepartments" gorm:"foreignKey:role_id;references:id"`
	RolePermissions []SysRolePermission `json:"rolePermissions" gorm:"foreignKey:role_id;references:id"`
	DefaultRouter   string              `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
	DefaultRole     *bool               `json:"defaultRole" gorm:"comment:是否为注册默认角色"`
	Description     string              `json:"description" gorm:"comment:描述"`
}

//type User struct {
//	gorm.Model
//	Profiles []Profile `gorm:"many2many:user_profiles;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;joinReferences:ProfileRefer"`
//	Refer    uint      `gorm:"index:,unique"`
//}
//
//type Profile struct {
//	gorm.Model
//	Name      string
//	UserRefer uint `gorm:"index:,unique"`
//}
//
//// Which creates join table: user_profiles
////   foreign key: user_refer_id, reference: users.refer
////   foreign key: profile_refer, reference: profiles.user_refer
