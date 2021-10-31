package model

import (
	"gin-myboot/global"
)

type JwtBlacklist struct {
	global.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}
