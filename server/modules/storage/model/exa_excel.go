package model

import system "gin-myboot/modules/system/model"

type ExcelInfo struct {
	FileName string                 `json:"fileName"` // 文件名
	InfoList []system.SysPermission `json:"infoList"`
}
