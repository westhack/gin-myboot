package model

import "errors"

// 初始版本自动化代码工具
type AutoCodeStruct struct {
	ModuleName          string   `json:"moduleName"`         // 模块名称
	StructName          string   `json:"structName"`         // Struct名称
	TableName           string   `json:"tableName"`          // 表名
	PackageName         string   `json:"packageName"`        // 文件名称
	FileName            string   `json:"fileName"`           // go文件名称
	Abbreviation        string   `json:"abbreviation"`       // Struct简称
	Description         string   `json:"description"`        // Struct中文名称
	AutoCreateApiToSql  bool     `json:"autoCreateApiToSql"` // 是否自动创建api
	AutoMoveFile        bool     `json:"autoMoveFile"`       // 是否自动移动文件
	IsSearch            bool     `json:"isSearch"`
	IsTableCreateUpdate bool     `json:"isTableCreateUpdate"`
	IsFormCreateUpdate  bool     `json:"isFormCreateUpdate"`
	IsDblclickUpdate    bool     `json:"isDblclickUpdate"`
	IsBatchDelete       bool     `json:"isBatchDelete"`
	IsTableDelete       bool     `json:"isTableDelete"`
	IsDownload          bool     `json:"isDownload"`
	ShowPopover         bool     `json:"showPopover"`
	RowKey              string   `json:"rowKey"`
	Fields              []*Field `json:"fields"`
}

type Field struct {
	FieldName       string   `json:"fieldName"`       // Field名
	FieldDesc       string   `json:"fieldDesc"`       // 中文名
	FieldType       string   `json:"fieldType"`       // Field数据类型
	FieldJson       string   `json:"fieldJson"`       // FieldJson
	DataType        string   `json:"dataType"`        // 数据库字段类型
	DataTypeLong    string   `json:"dataTypeLong"`    // 数据库字段长度
	ColumnComment   string   `json:"columnComment"`   // 数据库字段描述
	ColumnName      string   `json:"columnName"`      // 数据库字段
	FieldSearchType string   `json:"fieldSearchType"` // 搜索条件
	DictType        string   `json:"dictType"`        // 字典
	InputType       string   `json:"inputType"`       // 表单类型
	InputRules      []string `json:"inputRules"`      // 表单验证规则
	HiddenPopover   bool     `json:"hiddenPopover"`
	TableWidth      string   `json:"tableWidth"`
	TableAlign      string   `json:"tableAlign"`
}

var AutoMoveErr error = errors.New("创建代码成功并移动文件成功")
