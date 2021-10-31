package service

import (
	"gin-myboot/config"
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	"gin-myboot/utils"
	"go.uber.org/zap"
)

type SystemConfigService struct {
}

// GetSystemConfig
// @function: GetSystemConfig
// @description: 读取配置文件
// @return: err error, conf config.Server
func (systemConfigService *SystemConfigService) GetSystemConfig() (err error, conf config.Server) {
	return nil, global.Config
}

// SetSystemConfig
// @description   set system config,
// @function: SetSystemConfig
// @description: 设置配置文件
// @param: system model.System
// @return: err error
func (systemConfigService *SystemConfigService) SetSystemConfig(system system.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.Viper.Set(k, v)
	}
	err = global.Viper.WriteConfig()
	return err
}

// GetServerInfo
// @function: GetServerInfo
// @description: 获取服务器信息
// @return: server *utils.Server, err error
func (systemConfigService *SystemConfigService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.Logger.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Rrm, err = utils.InitRAM(); err != nil {
		global.Logger.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.Logger.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}