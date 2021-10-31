package api

import (
	"errors"
	"fmt"
	"gin-myboot/global"
	"gin-myboot/modules/common/model/response"
	"gin-myboot/modules/generator/model"
	"gin-myboot/modules/generator/model/request"
	"gin-myboot/modules/generator/service"
	"os"

	"gin-myboot/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/url"
)

type AutoCodeApi struct {
}

var autoCodeVerify = utils.Rules{
	"ModuleName":   {utils.NotEmpty()},
	"Abbreviation": {utils.NotEmpty()},
	"StructName":   {utils.NotEmpty()},
	//"PackageName":  {utils.NotEmpty()},
	"Fields":       {utils.NotEmpty()},
}

var AutoCodeApiApp = new(AutoCodeApi)
var autoCodeHistoryService = service.AutoCodeHistoryServiceApp
var autoCodeService = service.AutoCodeServiceApp

// Delete
// @Tags AutoCode
// @Summary 删除回滚记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AutoHistoryByID true "删除回滚记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /generator/autoCode/delete [post]
func (autoApi *AutoCodeApi) Delete(c *gin.Context) {
	var id request.AutoHistoryByID
	_ = c.ShouldBindJSON(&id)
	err := autoCodeHistoryService.Delete(id.ID)
	if err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	}
	response.OkWithMessage("删除成功", c)

}

// DeleteByIds
// @Tags AutoCode
// @Summary 删除回滚记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AutoHistoryByIDs true "删除回滚记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /generator/autoCode/deleteByIds [post]
func (autoApi *AutoCodeApi) DeleteByIds(c *gin.Context) {
	var id request.AutoHistoryByIDs
	_ = c.ShouldBindJSON(&id)
	err := autoCodeHistoryService.DeleteByIds(id.ID)
	if err != nil {
		global.Logger.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	}
	response.OkWithMessage("批量删除成功", c)

}

// GetList
// @Tags AutoCode
// @Summary 查询回滚记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SysAutoHistory true "查询回滚记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /generator/autoCode/getList [post]
func (autoApi *AutoCodeApi) GetList(c *gin.Context) {
	var search request.SysAutoHistory
	_ = c.ShouldBindJSON(&search)
	err, list, total := autoCodeHistoryService.GetList(search.PageInfo)
	if err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     search.Page,
			PageSize: search.PageSize,
		}, "获取成功", c)
	}
}

// RollBack 
// @Tags AutoCode
// @Summary 回滚
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AutoHistoryByID true "回滚自动生成代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"回滚成功"}"
// @Router /generator/autoCode/rollback [post]
func (autoApi *AutoCodeApi) RollBack(c *gin.Context) {
	var id request.AutoHistoryByID
	_ = c.ShouldBindJSON(&id)
	if err := autoCodeHistoryService.RollBack(id.ID); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("回滚成功", c)
}

// GetMeta 
// @Tags AutoCode
// @Summary 回滚
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AutoHistoryByID true "获取meta信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /generator/autoCode/getMeta [post]
func (autoApi *AutoCodeApi) GetMeta(c *gin.Context) {
	var id request.AutoHistoryByID
	_ = c.ShouldBindJSON(&id)
	if v, err := autoCodeHistoryService.GetMeta(id.ID); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithDetailed(gin.H{"meta": v}, "获取成功", c)
	}

}

// PreviewTemp
// @Tags AutoCode
// @Summary 预览创建后的代码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AutoCodeStruct true "预览创建代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /generator/autoCode/preview [post]
func (autoApi *AutoCodeApi) PreviewTemp(c *gin.Context) {
	var a model.AutoCodeStruct
	_ = c.ShouldBindJSON(&a)
	if err := utils.Verify(a, autoCodeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	a.PackageName = a.ModuleName
	autoCode, err := autoCodeService.PreviewTemp(a)
	if err != nil {
		global.Logger.Error("预览失败!", zap.Any("err", err))
		response.FailWithMessage("预览失败", c)
	} else {
		response.OkWithDetailed(gin.H{"autoCode": autoCode}, "预览成功", c)
	}
}

// CreateTemp
// @Tags AutoCode
// @Summary 自动代码模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AutoCodeStruct true "创建自动代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /generator/autoCode/createTemp [post]
func (autoApi *AutoCodeApi) CreateTemp(c *gin.Context) {
	var a model.AutoCodeStruct
	_ = c.ShouldBindJSON(&a)
	if err := utils.Verify(a, autoCodeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	a.PackageName = a.ModuleName
	var apiIds []uint64
	if a.AutoCreateApiToSql {
		if ids, err := autoCodeService.AutoCreateApi(&a); err != nil {
			global.Logger.Error("自动化创建失败!请自行清空垃圾数据!", zap.Any("err", err))
			c.Writer.Header().Add("success", "false")
			c.Writer.Header().Add("message", url.QueryEscape("自动化创建失败!请自行清空垃圾数据!"))
			return
		} else {
			apiIds = ids
		}
	}
	err := autoCodeService.CreateTemp(a, apiIds...)
	if err != nil {
		if errors.Is(err, model.AutoMoveErr) {
			c.Writer.Header().Add("success", "false")
			c.Writer.Header().Add("msgtype", "success")
			c.Writer.Header().Add("message", url.QueryEscape(err.Error()))
		} else {
			c.Writer.Header().Add("success", "false")
			c.Writer.Header().Add("message", url.QueryEscape(err.Error()))
			_ = os.Remove("./ginmyboot.zip")
		}
	} else {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "ginmyboot.zip")) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.Header().Add("success", "true")
		c.File("./ginmyboot.zip")
		_ = os.Remove("./ginmyboot.zip")
	}
}

// GetTables
// @Tags AutoCode
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /generator/autoCode/getTables [get]
func (autoApi *AutoCodeApi) GetTables(c *gin.Context) {
	dbName := c.DefaultQuery("dbName", global.Config.Mysql.Dbname)
	err, tables := autoCodeService.GetTables(dbName)
	if err != nil {
		global.Logger.Error("查询table失败!", zap.Any("err", err))
		response.FailWithMessage("查询table失败", c)
	} else {
		var items []map[string]string
		for _, t := range tables {
			m := make(map[string]string)
			if t.TableComment == "" {
				m["description"] = t.TableComment
				m["label"] = t.TableName
			} else {
				m["description"] = t.TableName
				m["label"] = t.TableName + " - " + t.TableComment
			}
			m["value"] = t.TableName
			items = append(items, m)
		}

		response.OkWithDetailed(gin.H{"items": items}, "获取成功", c)
	}
}

// GetDatabases
// @Tags AutoCode
// @Summary 获取当前所有数据库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /generator/autoCode/getDatabases [get]
func (autoApi *AutoCodeApi) GetDatabases(c *gin.Context) {
	if err, dbs := autoCodeService.GetDatabases(); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		var items []map[string]string
		for _, db := range dbs {
			m := make(map[string]string)
			m["label"] = db.Database
			m["value"] = db.Database
			items = append(items, m)
		}

		response.OkWithDetailed(gin.H{"items": items, "dbName": global.Config.Mysql.Dbname}, "获取成功", c)
	}
}

// GetColumns
// @Tags AutoCode
// @Summary 获取当前表所有字段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /generator/autoCode/getColumns [get]
func (autoApi *AutoCodeApi) GetColumns(c *gin.Context) {
	dbName := c.DefaultQuery("dbName", global.Config.Mysql.Dbname)
	tableName := c.Query("tableName")
	if err, columns := autoCodeService.GetColumns(tableName, dbName); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {

		for i, column := range columns {
			if column.DataType == "text" || column.DataType == "longtext" {
				columns[i].DataTypeLong = ""
			}
		}

		response.OkWithDetailed(gin.H{"items": columns}, "获取成功", c)
	}
}

// GetModules
// @Tags AutoCode
// @Summary 获取模块列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /generator/autoCode/getModules [get]
func (autoApi *AutoCodeApi) GetModules(c *gin.Context) {
	if err, list := autoCodeService.GetModules(); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"items": list}, "获取成功", c)
	}
}
