package service

import (
	"encoding/json"
	"gin-myboot/global"
	 "gin-myboot/modules/install/model"
	"gin-myboot/utils"
	"github.com/ghodss/yaml"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
)

type SystemInstallService struct {
}

var SystemInstallServiceApp = new(SystemInstallService)

// SetSystemConfig
// @description   set system config,
// @function: SetSystemConfig
// @description: 设置配置文件
// @param: system model.System
// @return: err error
func (systemInstallService *SystemInstallService) SetSystemConfig(config model.InstallConfig) (err error) {
	cs := utils.StructToMap(config)
	for k, v := range cs {
		if k != "isInitDb" {
			global.Viper.Set(k, v)
		}
	}
	err = global.Viper.WriteConfig()
	return err
}

// GetServerInfo
// @function: GetServerInfo
// @description: 获取服务器信息
// @return: server *utils.Server, err error
func (systemInstallService *SystemInstallService) GetServerInfo() (server *utils.Server, err error) {
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


// GetSystemConfig
// @function: GetSystemConfig
// @description: 读取配置文件表单
// @return: err error, conf system.SysConfig
func (systemInstallService *SystemInstallService) GetSystemConfig() (err error, conf []interface{}) {
	// var config []system.SysConfig
	// err = global.GormDB.Where(" parent_name = ?", "system").Order("sort_order asc").Find(&config).Error

	var basePath, _ = os.Getwd()
	y, err := ioutil.ReadFile(basePath + "/modules/install/config.yaml")
	jsonStr, err := yaml.YAMLToJSON(y)
	if err != nil {
		return err, nil
	}

	var config []interface{}
	err = json.Unmarshal(jsonStr, &config)
	if err != nil {
		//fmt.Printf("err: %v\n", err)
		return err, nil
	}

	return nil, config
}