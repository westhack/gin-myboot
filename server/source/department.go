package source

import (
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	"github.com/gookit/color"
	"time"

	"gorm.io/gorm"
)

var Department = new(department)

type department struct{}

//@description: sys_departments 表数据初始化
func (d *department) Init() error {
	var departments = []system.SysDepartment{
		{Model: global.Model{ID: 1, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, ParentId: 0, Title: "总部", Status: true, IsParent: true},
		{Model: global.Model{ID: 2, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, ParentId: 1, Title: "技术部", Status: true, IsParent: false},
		{Model: global.Model{ID: 3, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, ParentId: 2, Title: "研发中心", Status: true, IsParent: false},
		{Model: global.Model{ID: 4, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, ParentId: 3, Title: "大数据", Status: true, IsParent: false},
		{Model: global.Model{ID: 5, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()},ParentId: 3, Title: "人工智障", Status: true, IsParent: false},
		{Model: global.Model{ID: 6, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, ParentId: 0,Title: "成都分部", Status: true, IsParent: true},
		{Model: global.Model{ID: 7, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, ParentId:3 ,Title: "Vue", Status: true, IsParent: false},
		{Model: global.Model{ID: 8, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()},ParentId: 3, Title: "JAVA", Status: true, IsParent: false},
		{Model: global.Model{ID: 9, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, ParentId: 0,Title: "人事部", Status: true, IsParent: true},
		{Model: global.Model{ID: 10, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()},ParentId: 2, Title: "游客", Status: true, IsParent: false},
		{Model: global.Model{ID: 11, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, ParentId: 2,Title: "VIP", Status: true, IsParent: false},
	}
	return global.GormDB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 6}).Find(&[]system.SysDepartment{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_departments 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&departments).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_departments 表初始数据成功!")
		return nil
	})
}
