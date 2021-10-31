package source

import (
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	"github.com/gookit/color"
	"time"

	"gorm.io/gorm"
)

var Dict = new(dict)

type dict struct{}

var status = new(bool)

//@description: sys_dicts 表数据初始化
func (d *dict) Init() error {
	*status = true
	var dictionaries = []system.SysDict{
		{Model: global.Model{ID: 1, CreatedAt: global.LocalTime{Time:time.Now()}, UpdatedAt: time.Now()}, Name: "性别", Type: "sex", Status: status, Description: "性别字典"},
		{Model: global.Model{ID: 2, CreatedAt: global.LocalTime{Time:time.Now()}, UpdatedAt: time.Now()}, Name: "数据库int类型", Type: "int", Status: status, Description: "int类型对应的数据库类型"},
		{Model: global.Model{ID: 3, CreatedAt: global.LocalTime{Time:time.Now()}, UpdatedAt: time.Now()}, Name: "数据库时间日期类型", Type: "time.Time", Status: status, Description: "数据库时间日期类型"},
		{Model: global.Model{ID: 4, CreatedAt: global.LocalTime{Time:time.Now()}, UpdatedAt: time.Now()}, Name: "数据库浮点型", Type: "float64", Status: status, Description: "数据库浮点型"},
		{Model: global.Model{ID: 5, CreatedAt: global.LocalTime{Time:time.Now()}, UpdatedAt: time.Now()}, Name: "数据库字符串", Type: "string", Status: status, Description: "数据库字符串"},
		{Model: global.Model{ID: 6, CreatedAt: global.LocalTime{Time:time.Now()}, UpdatedAt: time.Now()}, Name: "数据库bool类型", Type: "bool", Status: status, Description: "数据库bool类型"},
		{Model: global.Model{ID: 7, CreatedAt: global.LocalTime{Time:time.Now()}, UpdatedAt: time.Now()}, Name: "文件上传类型", Type: "uploadType", Status: status, Description: "文件上传类型"},
		{Model: global.Model{ID: 8, CreatedAt: global.LocalTime{Time:time.Now()}, UpdatedAt: time.Now()}, Name: "消息类型", Type: "messageType", Status: status, Description: "消息类型"},
		{Model: global.Model{ID: 9, CreatedAt: global.LocalTime{Time:time.Now()}, UpdatedAt: time.Now()}, Name: "数据库字段类型", Type: "dataType", Status: status, Description: "数据库字段类型"},
		{Model: global.Model{ID: 10, CreatedAt: global.LocalTime{Time:time.Now()}, UpdatedAt: time.Now()}, Name: "字段类型", Type: "fieldType", Status: status, Description: "字段类型"},
	}
	return global.GormDB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 6}).Find(&[]system.SysDict{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_dicts 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&dictionaries).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_dicts 表初始数据成功!")
		return nil
	})
}
