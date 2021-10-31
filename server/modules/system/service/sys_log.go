package service

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model"
	"gin-myboot/modules/common/model/request"
	system "gin-myboot/modules/system/model"
	"unicode/utf8"
)

type LogService struct {
}

// Create
// @function: Create
// @description: 创建记录
// @param: log model.SysLog
// @return: err error
func (logService *LogService) Create(log system.SysLog) (err error) {
	b := utf8.ValidString(log.Body)
	if !b {
		log.Body = ""
	}
	err = global.GormDB.Create(&log).Error
	return err
}

// DeleteByIds
// @function: DeleteByIds
// @description: 批量删除记录
// @param: ids []uint64
// @return: err error
func (logService *LogService) DeleteByIds(ids []uint64) (err error) {
	err = global.GormDB.Delete(&[]system.SysLog{}, "id in (?)", ids).Error
	return err
}

// DeleteAll
// @function: DeleteAll
// @description: 删除全部
// @return: err error
func (logService *LogService) DeleteAll() (err error) {
	go func() {
		global.GormDB.Unscoped().Where("id > 0").Delete(&[]system.SysLog{})
	}()
	return err
}

// Delete
// @function: Delete
// @description: 删除操作记录
// @param: id uint64
// @return: err error
func (logService *LogService) Delete(id uint64) (err error) {
	err = global.GormDB.Delete(&[]system.SysLog{}, "id = ?", id).Error
	return err
}

// GetById
// @function: GetById
// @description: 根据id获取单条操作记录
// @param: id uint64
// @return: err error, log system.SysLog
func (logService *LogService) GetById(id uint64) (err error, log system.SysLog) {
	err = global.GormDB.Where("id = ?", id).First(&log).Error
	return
}

// GetList
// @function: GetList
// @description: 分页获取操作记录列表
// @param: info queryParams request.QueryParams
// @return: err error, list interface{}, total int64
func (logService *LogService) GetList(queryParams request.QueryParams) (err error, res interface{}, total int64) {
	limit := queryParams.PageSize
	offset := queryParams.PageSize * (queryParams.Page - 1)

	if queryParams.SortOrder.Column == "" {
		queryParams.SortOrder.Column = "id"
		queryParams.SortOrder.Order = "desc"
	}

	// 创建db
	db := global.GormDB.Model(&system.SysLog{})
	var list []system.SysLog
	db.Scopes(model.Search(queryParams.Search))

	err = db.Count(&total).Error

	if err != nil {
		return err, list, total
	} else {
		err = db.Scopes(model.SortOrder(queryParams.SortOrder)).
			Limit(limit).Offset(offset).Find(&list).Error

	}

	return err, list, total
}
