package source

import (
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	"github.com/gookit/color"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var Admin = new(admin)

type admin struct{}

var admins = []system.SysUser{
	{Model: global.Model{ID: 1, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", Nickname: "超级管理员", Avatar: "https://ooo.0o0.ooo/2019/04/28/5cc5a71a6e3b6.png", RoleId: 1, IsAdmin: true, Type: 1, Status: 1},
	{Model: global.Model{ID: 2, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "a303176530", Password: "3ec063004a6f31642261936a379fde3d", Nickname: "westhack", Avatar: "https://ooo.0o0.ooo/2019/04/28/5cc5a71a6e3b6.png", RoleId: 2, IsAdmin: true, Type: 1, Status: 1},
}

//@description: sys_users 表数据初始化
func (a *admin) Init() error {
	return global.GormDB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]system.SysUser{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_users 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_users 表初始数据成功!")
		return nil
	})
}
