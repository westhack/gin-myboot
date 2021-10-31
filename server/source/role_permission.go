package source

import (
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var RolePermission = new(rolePermission)

type rolePermission struct{}

var rolePermissionModel = []system.SysRolePermission{
	{1, 1},
	{1, 2},
	{1, 3},
	{1, 4},
	{1, 5},
	{1, 6},
	{1, 7},
	{1, 8},
	{1, 9},
	{1, 10},
	{1, 11},
	{1, 12},
	{1, 13},
	{1, 14},
	{1, 15},
	{1, 16},
	{1, 17},
	{1, 18},
	{1, 19},
	{1, 20},
	{1, 21},
	{1, 22},
	{1, 23},
	{1, 24},
	{1, 25},
	{1, 27},
	{3, 1},
	{3, 2},
	{3, 3},
	{3, 4},
	{3, 5},

}

//@description: sys_role_permissions 数据初始化
func (a *rolePermission) Init() error {
	return global.GormDB.Model(&system.SysRolePermission{}).Transaction(func(tx *gorm.DB) error {
		if tx.Where("role_id = 1").Find(&[]system.SysRolePermission{}).RowsAffected == 25 {
			color.Danger.Println("\n[Mysql] --> sys_role_permissions 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&rolePermissionModel).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_role_permissions 表初始数据成功!")
		return nil
	})
}
