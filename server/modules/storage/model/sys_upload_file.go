package model

import (
	"gin-myboot/global"
)

type SysUploadFile struct {
	global.Model
	Name     string `json:"name" gorm:"comment:文件名"`                                                 // 文件名
	Url      string `json:"url" gorm:"comment:文件地址"`                                                 // 文件地址
	Tag      string `json:"tag" gorm:"comment:文件标签"`                                                 // 文件标签
	Uuid     string `json:"uuid" gorm:"comment:编号"`                                                  // 编号
	Type     string `json:"type" gorm:"comment:类型:local=本地,qiniu=七牛,tencent-cos=腾讯云,aliyun-oss=阿里云"` // 存储类型
	UserId   uint64 `json:"userId" gorm:"comment:上传用户"`                                              // 上传用户
	Md5      string `json:"md5" gorm:"comment:文件MD5"`
	FileSize uint32 `json:"fileSize" gorm:"comment:文件大小"`
}
