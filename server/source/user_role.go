package source

import (
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var UserRole = new(userRole)

type userRole struct{}

var userRoleModel = []system.SysUserRole{
	{1, 1},
	{1, 2},
	{1, 3},
	{2, 1},
}

//@description: sys_user_roles 数据初始化
func (a *userRole) Init() error {
	return global.GormDB.Model(&system.SysUserRole{}).Transaction(func(tx *gorm.DB) error {
		if tx.Where("user_id IN (1, 2)").Find(&[]system.SysUserRole{}).RowsAffected == 4 {
			color.Danger.Println("\n[Mysql] --> sys_user_roles 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&userRoleModel).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_user_roles 表初始数据成功!")
		return nil
	})
}
