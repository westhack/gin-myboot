package main

import (
	"gin-myboot/core"
	"gin-myboot/global"
	"gin-myboot/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /api/
func main() {
	global.Viper = core.Viper()      // 初始化Viper
	global.Logger = core.Zap()       // 初始化zap日志库
	global.GormDB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	if global.GormDB != nil {
		initialize.MysqlTables(global.GormDB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GormDB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
