package core

import (
	"fmt"
	"gin-myboot/global"
	"gin-myboot/initialize"
	"gin-myboot/utils"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	// 初始化redis服务
	initialize.Redis()
	utils.InitValidator()
	Router := initialize.Routers()
	Router.Static("/doc", "./resource/doc")

	address := fmt.Sprintf(":%d", global.Config.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.Logger.Info("server run success on ", zap.String("address", address))
	fmt.Printf(`
	欢迎使用 gin-myboot
	当前版本:V1.0.1
    加群方式:微信号：westhack
	默认自动化文档地址:http://127.0.0.1%v/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8000
	如果项目让您获得了收益，请 star :https://github.com/westhack/gin-myboot
`, address)
	global.Logger.Error(s.ListenAndServe().Error())
}
