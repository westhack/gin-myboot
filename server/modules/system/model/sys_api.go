package model

import (
	"gin-myboot/global"
	"gorm.io/gorm"
)

type SysApi struct {
	global.Model
	Path        string `json:"path" gorm:"comment:api路径" binding:"required" required_err:"api路径不能为空"`   // api路径
	Description string `json:"description" gorm:"comment:api中文描述"`                                      // api中文描述
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组" binding:"required" required_err:"api组不能为空"` // api组
	Method      string `json:"method" gorm:"default:POST;comment:方法"`                                   // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
	Label       string `json:"label" gorm:"-"`
	Value       string `json:"value" gorm:"-"`
}

func (u *SysApi) AfterFind(tx *gorm.DB) (err error) {
	u.Label = u.Description + " " + u.Path + " " + u.Method
	u.Value = u.Path + ":" + u.Method
	return
}
