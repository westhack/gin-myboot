package source

import (
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	"github.com/gookit/color"
	"time"

	"gorm.io/gorm"
)

var DictDetail = new(dictDetail)

type dictDetail struct{}

//@description: dict_details 表数据初始化
func (d *dictDetail) Init() error {
	var details = []system.SysDictDetail{
		{global.Model{ID: 1, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "smallint", "1", "", status, 1, 2},
		{global.Model{ID: 2, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "mediumint", "2", "", status, 2, 2},
		{global.Model{ID: 3, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "int", "3", "", status, 3, 2},
		{global.Model{ID: 4, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "bigint", "4", "", status, 4, 2},
		{global.Model{ID: 5, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "date", "0", "", status, 0, 3},
		{global.Model{ID: 6, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "time", "1", "", status, 1, 3},
		{global.Model{ID: 7, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "year", "2", "", status, 2, 3},
		{global.Model{ID: 8, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "datetime", "3", "", status, 3, 3},
		{global.Model{ID: 9, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "timestamp", "5", "", status, 5, 3},
		{global.Model{ID: 10, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "float", "0", "", status, 0, 4},
		{global.Model{ID: 11, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "double", "1", "", status, 1, 4},
		{global.Model{ID: 12, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "decimal", "2", "", status, 2, 4},
		{global.Model{ID: 13, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "char", "0", "", status, 0, 5},
		{global.Model{ID: 14, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "varchar", "1", "", status, 1, 5},
		{global.Model{ID: 15, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "tinyblob", "2", "", status, 2, 5},
		{global.Model{ID: 16, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "tinytext", "3", "", status, 3, 5},
		{global.Model{ID: 17, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "text", "4", "", status, 4, 5},
		{global.Model{ID: 18, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "blob", "5", "", status, 5, 5},
		{global.Model{ID: 19, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "mediumblob", "6", "", status, 6, 5},
		{global.Model{ID: 20, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "mediumtext", "7", "", status, 7, 5},
		{global.Model{ID: 21, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "longblob", "8", "", status, 8, 5},
		{global.Model{ID: 22, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "longtext", "9", "", status, 9, 5},
		{global.Model{ID: 23, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "tinyint", "0", "", status, 0, 6},

		{global.Model{ID: 24, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "本地", "local", "", status, 0, 7},
		{global.Model{ID: 25, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "七牛", "qiniu", "", status, 0, 7},
		{global.Model{ID: 26, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "腾讯云", "tencent-cos", "", status, 0, 7},
		{global.Model{ID: 27, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "阿里云", "aliyun-oss", "", status, 0, 7},

		{global.Model{ID: 28, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "系统消息", "0", "", status, 0, 8},
		{global.Model{ID: 29, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "普通消息", "1", "", status, 0, 8},

		{global.Model{ID: 30, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "int", "int", "", status, 0, 9},
		{global.Model{ID: 31, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "bigint", "bigint", "", status, 0, 9},
		{global.Model{ID: 32, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "varchar", "varchar", "", status, 0, 9},
		{global.Model{ID: 33, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "datetime", "datetime", "", status, 0, 9},
		{global.Model{ID: 34, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "decimal", "decimal", "", status, 0, 9},
		{global.Model{ID: 35, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "char", "char", "", status, 0, 9},
		{global.Model{ID: 36, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "tinyint", "tinyint", "", status, 0, 9},
		{global.Model{ID: 37, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "text", "text", "", status, 0, 9},
		{global.Model{ID: 38, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "longtext", "longtext", "", status, 0, 9},

		{global.Model{ID: 39, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "string", "string", "", status, 0, 10},
		{global.Model{ID: 40, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "int", "int", "", status, 0, 10},
		{global.Model{ID: 41, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "uint", "uint", "", status, 0, 10},
		{global.Model{ID: 42, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "int32", "int32", "", status, 0, 10},
		{global.Model{ID: 43, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "uint32", "uint32", "", status, 0, 10},
		{global.Model{ID: 44, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "int64", "int64", "", status, 0, 10},
		{global.Model{ID: 45, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "uint64", "uint64", "", status, 0, 10},
		{global.Model{ID: 46, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "float32", "float32", "", status, 0, 10},
		{global.Model{ID: 47, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "float64", "float64", "", status, 0, 10},
		{global.Model{ID: 48, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "bool", "bool", "", status, 0, 10},
		{global.Model{ID: 49, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "time.Time", "time.Time", "", status, 0, 10},
		{global.Model{ID: 50, CreatedAt: global.LocalTime{Time: time.Now()}, UpdatedAt: time.Now()}, "interface{}", "interface{}", "", status, 0, 10},
	}
	return global.GormDB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 23}).Find(&[]system.SysDictDetail{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_dict_details 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&details).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_dict_details 表初始数据成功!")
		return nil
	})
}
