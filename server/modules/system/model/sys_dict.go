// 自动生成模板SysDict
package model

import (
	"gin-myboot/global"
)

// 如果含有time.Time 请自行import time包
type SysDict struct {
	global.Model
	Name        string          `json:"name" form:"name" gorm:"column:name;comment:字典名（中）"`   // 字典名（中）
	Type        string          `json:"type" form:"type" gorm:"column:type;comment:字典名（英）"`   // 字典名（英）
	Status      *bool           `json:"status" form:"status" gorm:"column:status;comment:状态"` // 状态
	SortOrder   float64             `json:"sortOrder" form:"sort_order" gorm:"column:sort_order;comment:排序标记"`
	Description string          `json:"description" form:"description" gorm:"column:description;comment:描述"` // 描述
	DictDetails []SysDictDetail `json:"dictDetails" gorm:"foreignKey:dict_id;references:id"`
}
