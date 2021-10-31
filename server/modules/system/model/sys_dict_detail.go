// 自动生成模板SysDictDetail
package model

import (
	"gin-myboot/global"
)

// 如果含有time.Time 请自行import time包
type SysDictDetail struct {
	global.Model
	Label     string  `json:"label" form:"label" gorm:"column:label;comment:字典健;type:varchar(255);"`               // 字典健
	Value     string  `json:"value" form:"value" gorm:"column:value;comment:字典值;type:varchar(1000);"`               // 字典值
	Color     string  `json:"color" form:"color" gorm:"column:color;comment:颜色"`                // 字典值
	Status    *bool   `json:"status" form:"status" gorm:"column:status;comment:启用状态"`           // 启用状态
	SortOrder float64 `json:"sortOrder" form:"sortOrder" gorm:"column:sort_order;comment:排序标记"` // 排序标记
	DictId    uint64   `json:"dictId" form:"dictId" gorm:"column:dict_id;comment:关联标记"`          // 关联标记
}
