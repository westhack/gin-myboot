package api

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/response"
	common "gin-myboot/modules/common/service"
	system "gin-myboot/modules/system/model"
	"gin-myboot/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DemoApi struct {
	BaseApi
}


func GetModel() (db *gorm.DB) {
	var user system.SysLog
	return global.GormDB.Model(&user)
}

var TableName = "sys_users"

func NewDemoApi() DemoApi {

	service := common.Demo{
		BaseService: common.BaseService{
			TableName: TableName,
			Model: common.Model{AllLimit: 10, IdName: "id", Preloads: []string{}, StatusName: "status", GetModel: GetModel},
		},
	}

	return DemoApi{
		BaseApi{
			Service: &service,
			Rules: Rules{
				Create: map[string][]string{
					"body": {utils.NotEmpty()},
				},
				Update: map[string][]string{
					"id": {utils.NotEmpty(), utils.Ge("100")},
				},
			},
		},
	}
}

func (api *DemoApi) GetList(c *gin.Context)  {
	response.FailWithMessage("获取失败2", c)
}