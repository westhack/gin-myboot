package source

import (
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	"github.com/gookit/color"
	"time"

	"gorm.io/gorm"
)

var Role = new(role)

type role struct{}

var authorities = []system.SysRole{
	{CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now(), Name: "admin", Title: "超级管理员", ParentId: 0, DefaultRouter: "dashboard"},
	{CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now(), Name: "admin1", Title: "子角色", ParentId: 1, DefaultRouter: "dashboard"},
	{CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now(), Name: "test", Title: "测试角色", ParentId: 0, DefaultRouter: "dashboard"},
}

//@description: sys_roles 表数据初始化
func (a *role) Init() error {
	return global.GormDB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("name IN ? ", []string{"admin", "admin1"}).Find(&[]system.SysRole{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_roles 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorities).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_roles 表初始数据成功!")
		return nil
	})
}
