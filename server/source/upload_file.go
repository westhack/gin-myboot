package source

import (
	"gin-myboot/global"
	storage "gin-myboot/modules/storage/model"
	"github.com/gookit/color"
	"gorm.io/gorm"
	"time"
)

var File = new(file)

type file struct{}

var files = []storage.SysUploadFile{
	{global.Model{ID: 1, CreatedAt: global.LocalTime{Time:time.Now()}, UpdatedAt: time.Now()}, "5cc5a71a6e3b6.png", "https://ooo.0o0.ooo/2019/04/28/5cc5a71a6e3b6.png", "png", "5cc5a71a6e3b6", "local",1, "", 0},
	{global.Model{ID: 2, CreatedAt: global.LocalTime{Time:time.Now()}, UpdatedAt: time.Now()}, "5cc5a71a6e3b7.png", "https://ooo.0o0.ooo/2019/04/28/5cc5a71a6e3b6.png", "png", "5cc5a71a6e3b7", "local", 1, "", 0},
}

//@description: sys_upload_files 表初始化数据
func (f *file) Init() error {
	return global.GormDB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]storage.SysUploadFile{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_upload_files 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&files).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_upload_files 表初始数据成功!")
		return nil
	})
}
