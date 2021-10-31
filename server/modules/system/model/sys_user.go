package model

import (
	"gin-myboot/global"
	"gorm.io/gorm"

	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	global.Model
	UUID         uuid.UUID       `json:"uuid" gorm:"comment:用户UUID"`    // 用户UUID
	Username     string          `json:"username" gorm:"comment:用户登录名"` // 用户登录名
	Password     string          `json:"-" gorm:"comment:用户登录密码"`       // 用户登录密码
	Mobile       string          `json:"mobile" gorm:"comment:手机号"`     // 用户登录密码
	Nickname     string          `json:"nickname" gorm:"default:系统用户;comment:用户昵称"`
	Status       int             `json:"status" form:"status" gorm:"column:status;comment:状态:1-正常,0-禁用，2-待审核"` // 状态
	IsAdmin      bool            `json:"isAdmin" form:"isAdmin" gorm:"comment:是否是管理员"`                         // 是否是管理员
	Type         int             `json:"type" form:"type" gorm:"comment:用户类型 0普通用户 1管理员"`
	Avatar       string          `json:"avatar" gorm:"default:https://ooo.0o0.ooo/2019/04/28/5cc5a71a6e3b6.png;comment:用户头像"` // 用户头像
	Role         SysRole         `json:"role" gorm:"foreignKey:role_id;references:id;comment:用户角色"`
	RoleId       uint64          `json:"roleId" gorm:"comment:用户角色ID"` // 用户角色ID
	Department   SysDepartment   `json:"department" gorm:"foreignKey:department_id;references:id;comment:用户部门"`
	DepartmentId uint64          `json:"departmentId" gorm:"comment:部门ID"`                  // 部门ID
	SideMode     string          `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`       // 用户侧边主题
	ActiveColor  string          `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"` // 活跃颜色
	BaseColor    string          `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`      // 基础颜色
	Roles        []SysRole       `json:"roles" gorm:"many2many:sys_user_roles;foreignKey:id;joinForeignKey:user_id;References:id;joinReferences:role_id"`
	Permissions  []SysPermission `json:"permissions" gorm:"-"`
	StatusText   string          `json:"statusText" gorm:"-"`
	Email        string          `json:"email" gorm:"comment:邮箱"`
	Address      string          `json:"address" gorm:"comment:省市县地址"`
	Street       string          `json:"street" gorm:"comment:街道地址"`
	Sex          string          `json:"sex" gorm:"comment:性别"`
	Description  string          `json:"description" gorm:"comment:描述/详情/备注"`
}

func (u *SysUser) AfterFind(tx *gorm.DB) (err error) {
	if u.Status == 1 {
		u.StatusText = "正常"
	} else if u.Status == 0 {
		u.StatusText = "禁用"
	} else if u.Status == 2 {
		u.StatusText = "待审核"
	} else if u.Status == -1 {
		u.StatusText = "锁住"
	}
	return
}
