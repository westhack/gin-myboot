package model

import (
	"gin-myboot/global"
	"gorm.io/gorm"
)

type SysMessage struct {
	global.Model
	SessionId  string         `json:"sessionId" gorm:"comment:会话ID"`
	Content    string         `json:"content" gorm:"comment:消息内容;type:text;"`
	Data       string         `json:"data" gorm:"comment:关联数据;type:text;"`
	DataId     uint64         `json:"dataId" gorm:"comment:关联ID"`
	DataType   string         `json:"dataType" gorm:"comment:关联类型"`
	FormUser   SysUser        `json:"formUser" gorm:"foreignKey:from_user_id;references:id"`
	FromUserId uint64         `json:"fromUserId" gorm:"comment:发送用户ID"`
	Status     bool           `json:"status" gorm:"comment:状态"`
	Title      string         `json:"title" gorm:"comment:标题"`
	Type       int            `json:"type" gorm:"comment:类型"`
	User       SysUser        `json:"user" gorm:"foreignKey:user_id;references:id"`
	UserId     uint64         `json:"userId" gorm:"comment:接受用户ID"`
	ViewTime   string         `json:"viewTime" gorm:"comment:查看时间"`
	Icon       string         `json:"icon" gorm:"comment:消息图标"`
	Image      string         `json:"image" gorm:"comment:消息缩略图"`
}

func (u *SysMessage) AfterFind(tx *gorm.DB) (err error) {
	if u.FormUser.ID > 0 {
		u.Icon = u.FormUser.Avatar
	} else {
		u.Icon = "/assets/images/message.png"
	}
	return
}
