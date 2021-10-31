// 自动生成模板SysLog
package model

import (
	"gin-myboot/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type SysLog struct {
	global.Model
	Ip           string        `json:"ip" form:"ip" gorm:"column:ip;comment:请求ip"`                                   // 请求ip
	Method       string        `json:"method" form:"method" gorm:"column:method;comment:请求方法"`                       // 请求方法
	Path         string        `json:"path" form:"path" gorm:"column:path;comment:请求路径"`                             // 请求路径
	Status       int           `json:"status" form:"status" gorm:"column:status;comment:请求状态"`                       // 请求状态
	Latency      time.Duration `json:"latency" form:"latency" gorm:"column:latency;comment:延迟" swaggertype:"string"` // 延迟
	Agent        string        `json:"agent" form:"agent" gorm:"column:agent;comment:代理"`                            // 代理
	ErrorMessage string        `json:"errorMessage" form:"errorMessage" gorm:"column:error_message;comment:错误信息"`    // 错误信息
	Body         string        `json:"body" form:"body" gorm:"type:longtext;column:body;comment:请求Body"`             // 请求Body
	Resp         string        `json:"resp" form:"resp" gorm:"type:longtext;column:resp;comment:响应Body"`             // 响应Body
	UserID       uint64         `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id"`                      // 用户id
	Name         string        `json:"name" form:"name" gorm:"column:name;comment:方法操作名称"`
	Username     string        `json:"username" form:"username" gorm:"column:username;comment:用户名"`
	LogType      int           `json:"logType" form:"logType" gorm:"column:log_type;comment:日志类型"`
	RequestUrl   string        `json:"requestUrl" form:"requestUrl" gorm:"column:request_url;comment:请求url"`
	IpInfo       string        `json:"ipInfo" form:"ipInfo" gorm:"column:ip_info;comment:IP信息"`
	Device       string        `json:"device" form:"device" gorm:"column:device;comment:设备"`
	User         SysUser       `json:"user"`
}
