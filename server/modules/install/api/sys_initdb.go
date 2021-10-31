package api

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/response"
	"gin-myboot/modules/install/model"
	"gin-myboot/modules/install/service"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type DBApi struct {
}

var initDBService = service.InitDBServiceApp
var DBApiApp = new(DBApi)

// InitDB
// @Tags InitDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Param data body model.Mysql true "初始化数据库参数"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"自动创建数据库成功"}"
// @Router /install/initdb [post]
func (i *DBApi) InitDB(c *gin.Context) {
	if global.GormDB != nil {
		global.Logger.Error("已存在数据库配置!")
		response.FailWithMessage("已存在数据库配置", c)
		return
	}
	var dbInfo model.Mysql
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.Logger.Error("参数校验不通过!", zap.Any("err", err))
		response.FailWithMessage("参数校验不通过", c)
		return
	}
	if err := initDBService.InitDB(dbInfo); err != nil {
		global.Logger.Error("自动创建数据库失败!", zap.Any("err", err))
		response.FailWithMessage("自动创建数据库失败，请查看后台日志，检查后在进行初始化", c)
		return
	}
	response.OkWithData("自动创建数据库成功", c)
}

// CheckDB
// @Tags CheckDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"探测完成"}"
// @Router /install/checkdb [post]
func (i *DBApi) CheckDB(c *gin.Context) {
	if global.GormDB != nil {
		global.Logger.Info("数据库无需初始化")
		response.OkWithDetailed(gin.H{"needInit": false}, "数据库无需初始化", c)
		return
	} else {
		global.Logger.Info("前往初始化数据库")
		response.OkWithDetailed(gin.H{"needInit": true}, "前往初始化数据库", c)
		return
	}
}
