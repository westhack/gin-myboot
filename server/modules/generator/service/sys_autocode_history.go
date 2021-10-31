package service

import (
	"fmt"
	"gin-myboot/global"
	"gin-myboot/modules/common/model/request"
	"gin-myboot/modules/generator/model"
	"gin-myboot/modules/system/service"
	"gin-myboot/utils"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap"
)

type AutoCodeHistoryService struct {
}

var AutoCodeHistoryServiceApp = new(AutoCodeHistoryService)

// CreateAutoCodeHistory RouterPath : RouterPath@RouterString;RouterPath2@RouterString2
func (autoCodeHistoryService *AutoCodeHistoryService) CreateAutoCodeHistory(meta, structName, structCNName, autoCodePath string, injectionMeta string, tableName string, apiIds string) error {
	return global.GormDB.Create(&model.SysAutoCodeHistory{
		RequestMeta:   meta,
		AutoCodePath:  autoCodePath,
		InjectionMeta: injectionMeta,
		StructName:    structName,
		StructCNName:  structCNName,
		TableName:     tableName,
		ApiIDs:        apiIds,
	}).Error
}

// RollBack 回滚
func (autoCodeHistoryService *AutoCodeHistoryService) RollBack(id uint) error {
	md := model.SysAutoCodeHistory{}
	if err := global.GormDB.First(&md, id).Error; err != nil {
		return err
	}
	// 清除API表
	err := service.ApiServiceApp.DeleteApisByIds(strings.Split(md.ApiIDs, ";"))
	if err != nil {
		global.Logger.Error("ClearTag DeleteApiByIds:", zap.Error(err))
	}
	// 获取全部表名
	err, dbNames := AutoCodeServiceApp.GetTables(global.Config.Mysql.Dbname)
	if err != nil {
		global.Logger.Error("ClearTag GetTables:", zap.Error(err))
	}
	// 删除表
	for _, name := range dbNames {
		if strings.Contains(strings.ToUpper(strings.Replace(name.TableName, "_", "", -1)), strings.ToUpper(md.TableName)) {
			// 删除表
			if err = AutoCodeServiceApp.DropTable(name.TableName); err != nil {
				global.Logger.Error("ClearTag DropTable:", zap.Error(err))

			}
		}
	}
	// 删除文件

	for _, path := range strings.Split(md.AutoCodePath, ";") {
		// 迁移
		nPath := filepath.Join(global.Config.AutoCode.Root,
			"rm_file", time.Now().Format("20060102"), filepath.Base(filepath.Dir(filepath.Dir(path))), filepath.Base(filepath.Dir(path)), filepath.Base(path))
		err = utils.FileMove(path, nPath)
		if err != nil {
			fmt.Println(">>>>>>>>>>>>>>>>>>>", err)
		}
		//_ = utils.DeLFile(path)
	}
	// 清除注入
	for _, v := range strings.Split(md.InjectionMeta, ";") {
		// RouterPath@functionName@RouterString
		meta := strings.Split(v, "@")
		if len(meta) == 3 {
			_ = utils.AutoClearCode(meta[0], meta[2])
		}
	}
	md.Flag = 1
	return global.GormDB.Save(&md).Error
}

func (autoCodeHistoryService *AutoCodeHistoryService) GetMeta(id uint) (string, error) {
	var meta string
	return meta, global.GormDB.Model(model.SysAutoCodeHistory{}).Select("request_meta").First(&meta, id).Error
}

// GetList  获取系统历史数据
func (autoCodeHistoryService *AutoCodeHistoryService) GetList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GormDB
	var fileLists []model.SysAutoCodeHistory
	err = db.Find(&fileLists).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error // .Select("id,created_at,updated_at,struct_name,struct_cn_name,flag,table_name")
	return err, fileLists, total
}

// Delete 删除历史数据
func (autoCodeHistoryService *AutoCodeHistoryService) Delete(id uint) error {
	return global.GormDB.Delete(model.SysAutoCodeHistory{}, id).Error
}
// DeleteByIds 批量删除历史数据
func (autoCodeHistoryService *AutoCodeHistoryService) DeleteByIds(id []uint) error {
	return global.GormDB.Where("id in ?", id).Delete(model.SysAutoCodeHistory{}).Error
}
