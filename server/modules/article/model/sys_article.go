// 自动生成模板SysArticle
package model

import (
    "gin-myboot/global"
)

// SysArticle 结构体
// 如果含有time.Time 请自行import time包
type SysArticle struct {
      global.Model
      Content  string `json:"content" form:"content" gorm:"column:content;comment:内容;type:text;"`
      Title  string `json:"title" form:"title" gorm:"column:title;comment:标题;type:varchar(255);"`
}


// TableName SysArticle 表名
func (SysArticle) TableName() string {
  return "sys_articles"
}

