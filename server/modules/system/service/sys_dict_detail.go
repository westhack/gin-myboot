package service

import (
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	"gin-myboot/modules/system/model/request"
)

type DictDetailService struct {
}

// Create
// @function: Create
// @description: 创建字典详情数据
// @param: sysDictDetail model.SysDictDetail
// @return: err error
func (dictDetailService *DictDetailService) Create(sysDictDetail system.SysDictDetail) (err error) {
	err = global.GormDB.Create(&sysDictDetail).Error
	return err
}

// Delete
// @function: Delete
// @description: 删除字典详情数据
// @param: sysDictDetail model.SysDictDetail
// @return: err error
func (dictDetailService *DictDetailService) Delete(sysDictDetail system.SysDictDetail) (err error) {
	err = global.GormDB.Delete(&sysDictDetail).Error
	return err
}

// Update
// @function: Update
// @description: 更新字典详情数据
// @param: sysDictDetail *model.SysDictDetail
// @return: err error
func (dictDetailService *DictDetailService) Update(sysDictDetail *system.SysDictDetail) (err error) {
	err = global.GormDB.Save(sysDictDetail).Error
	return err
}

// GetById
// @function: GetById
// @description: 根据id获取字典详情单条数据
// @param: id uint
// @return: err error, sysDictDetail model.SysDictDetail
func (dictDetailService *DictDetailService) GetById(id uint64) (err error, sysDictDetail system.SysDictDetail) {
	err = global.GormDB.Where("id = ?", id).First(&sysDictDetail).Error
	return
}

// GetList
// @function: GetList
// @description: 分页获取字典详情列表
// @param: info request.SysDictDetailSearch
// @return: err error, list interface{}, total int64
func (dictDetailService *DictDetailService) GetList(info request.SysDictDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GormDB.Model(&system.SysDictDetail{})
	var sysDictDetails []system.SysDictDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != "" {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.DictId != 0 {
		db = db.Where("dict_id = ?", info.DictId)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sysDictDetails).Error

	return err, sysDictDetails, total
}
