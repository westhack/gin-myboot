package api

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/response"
	"gin-myboot/modules/install/model"
	"gin-myboot/modules/install/service"
	"gin-myboot/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
)

type InstallApi struct {
}

var InstallApiApp = new(InstallApi)
var installService = service.SystemInstallServiceApp

// SetSystemConfig
// @Tags Install
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body model.InstallConfig true "设置配置文件内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /install/setSystemConfig [post]
func (s *InstallApi) SetSystemConfig(c *gin.Context) {

	_, err := ioutil.ReadFile("install.lock")
	if err == nil {
		response.FailWithDetailed(gin.H{"install": "ok"}, "系统已安装，如需重新安装请删除 install.lock", c)
		return
	}

	var sys model.InstallConfig
	_ = c.ShouldBindJSON(&sys)
	if err := installService.SetSystemConfig(sys); err != nil {
		global.Logger.Error("设置失败!", zap.Any("err", err))
		response.FailWithMessage("设置失败", c)
	} else {

		global.Logger.Error("init", zap.Any("InstallConfig", sys))
		if err := initDBService.InitDB(sys.Mysql); err != nil {
			global.Logger.Error("自动创建数据库失败!", zap.Any("err", err))
			response.FailWithMessage("自动创建数据库失败，请查看后台日志，检查后在进行初始化", c)
			return
		}

		err := ioutil.WriteFile("install.lock", []byte("ok"), 0644)
		if err != nil {
			global.Logger.Error("create install.lock error", zap.Any("err", err))
		} else {
			err = utils.Reload()
			if err != nil {
				global.Logger.Error("重启系统失败!", zap.Any("err", err))
				response.FailWithMessage("重启系统失败", c)
				return
			} else {
				response.OkWithDetailed(gin.H{"install": "ok"}, "重启系统成功", c)
				return
			}
		}

		response.OkWithData("设置成功", c)
	}
}

// GetServerInfo
// @Tags Install
// @Summary 获取服务器信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /install/getServerInfo [post]
func (s *InstallApi) GetServerInfo(c *gin.Context) {
	if server, err := installService.GetServerInfo(); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
	}
}

// GetSystemConfig
// @Tags Install
// @Summary 获取配置文件表单
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /install/getSystemConfig [post]
func (s *InstallApi) GetSystemConfig(c *gin.Context) {

	_, err := ioutil.ReadFile("install.lock")
	if err == nil {
		response.FailWithDetailed(gin.H{"install": "ok"}, "系统已安装，如需重新安装请删除 install.lock", c)
		return
	}

	if err, config := installService.GetSystemConfig(); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"config": config}, "获取成功", c)
	}
}

// Pong
// @Tags Install
// @Summary pong
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"安装成功"}"
// @Router /install/pong [get]
func (s *InstallApi) Pong(c *gin.Context) {

	_, err := ioutil.ReadFile("install.lock")
	if err == nil {
		response.OkWithDetailed(gin.H{"install": "ok"}, "系统成功", c)
		return
	}

	response.FailWithMessage("安装中...", c)
}
