package model

import (
	"gin-myboot/global"
	"gorm.io/gorm"
)

type SysConfig struct {
	global.Model
	ParentName  string      `json:"parentName" gorm:"column:parent_name;comment:父类"`
	Module      string      `json:"module" gorm:"column:module;comment:分组"`
	Label       string      `json:"label" gorm:"column:label;comment:名称（中）"`
	Name        string      `json:"name" gorm:"column:name;comment:名称（英）"`
	Type        string      `json:"type" gorm:"column:type;comment:类型"`
	Value       string      `json:"value" gorm:"column:value;comment:值;type:text;"`
	DataSource  string      `json:"dataSource" gorm:"column:data_source;comment:数据;type:longtext;"`
	RuleSource  string      `json:"ruleSource" gorm:"column:rule_source;comment:规则;type:text;"`
	Multiple    *bool       `json:"multiple" gorm:"column:multiple;comment:是否多值;"`
	Status      *bool       `json:"status" gorm:"column:status;comment:状态"`
	SortOrder   float64     `json:"sortOrder" gorm:"column:sort_order;comment:排序"`
	Description string      `json:"description" gorm:"column:description;comment:描述"`
	ValueType   string      `json:"valueType" gorm:"column:value_type;comment:值格式化"`
	Fields      string      `json:"fields" gorm:"column:fields;comment:动态表单字段;type:text;"`
	Limit       uint        `json:"limit" gorm:"column:limit;comment:动态input限制数量"`
	Children    []SysConfig `json:"children" gorm:"-"`
	Data        interface{} `json:"data" gorm:"-"`
	Rules       interface{} `json:"rules" gorm:"-"`
}

func (u *SysConfig) AfterFind(tx *gorm.DB) (err error) {

	return
}
