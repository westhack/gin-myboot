package service

import (
	"errors"
	"gin-myboot/global"
	system "gin-myboot/modules/system/model"
	"gin-myboot/modules/system/model/request"
	"gorm.io/gorm"
)

type DictService struct {
}

// Create
// @function: Create
// @description: 创建字典数据
// @param: sysDict model.SysDict
// @return: err error
func (dictService *DictService) Create(sysDict system.SysDict) (err error) {
	if (!errors.Is(global.GormDB.First(&system.SysDict{}, "type = ?", sysDict.Type).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的type，不允许创建")
	}
	err = global.GormDB.Create(&sysDict).Error
	return err
}

// Delete
// @function: Delete
// @description: 删除字典数据
// @param: sysDict model.SysDict
// @return: err error
func (dictService *DictService) Delete(id uint64) (err error) {

	var sysDict system.SysDict

	err = global.GormDB.Where("id =?", id).Preload("DictDetails").First(&sysDict).Error
	if err != nil {
		return err
	}

	err = global.GormDB.Delete(&sysDict).Delete(&sysDict.DictDetails).Error
	return err
}

// DeleteByIds
// @function: DeleteByIds
// @description: 批量删除字典数据
// @param: sysDict model.SysDict
// @return: err error
func (dictService *DictService) DeleteByIds(ids []uint64) (err error) {

	for _, id := range ids {
		err := dictService.Delete(id)
		if err != nil {
			return err
		}
	}

	return err
}

// Update
// @function: Update
// @description: 更新字典数据
// @param: sysDict *model.SysDict
// @return: err error
func (dictService *DictService) Update(sysDict *system.SysDict) (err error) {
	var dict system.SysDict
	sysDictMap := map[string]interface{}{
		"Name":        sysDict.Name,
		"Type":        sysDict.Type,
		"Status":      sysDict.Status,
		"Description": sysDict.Description,
	}
	db := global.GormDB.Where("id = ?", sysDict.ID).First(&dict)
	if dict.Type == sysDict.Type {
		err = db.Updates(sysDictMap).Error
	} else {
		if (!errors.Is(global.GormDB.First(&system.SysDict{}, "type = ?", sysDict.Type).Error, gorm.ErrRecordNotFound)) {
			return errors.New("存在相同的type，不允许创建")
		}
		err = db.Updates(sysDictMap).Error

	}
	return err
}

// GetByIdOrType
// @function: GetByIdOrType
// @description: 根据id或者type获取字典单条数据
// @param: Type string, Id uint
// @return: err error, sysDict model.SysDict
func (dictService *DictService) GetByIdOrType(Type string, Id uint64) (err error, sysDict system.SysDict) {
	err = global.GormDB.Where("type = ? OR id = ?", Type, Id).Preload("DictDetails").First(&sysDict).Error
	return
}

// GetList
// @function: GetList
// @description: 分页获取字典列表
// @param: info request.SysDictSearch
// @return: err error, list interface{}, total int64
func (dictService *DictService) GetList(info request.SysDictSearch) (err error, list interface{}, total int64) {

	// 创建db
	db := global.GormDB.Model(&system.SysDict{})
	var sysDicts []system.SysDict
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Type != "" {
		db = db.Where("`type` LIKE ?", "%"+info.Type+"%")
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.Description != "" {
		db = db.Where("`description` LIKE ?", "%"+info.Description+"%")
	}
	err = db.Count(&total).Error
	err = db.Preload("DictDetails").Find(&sysDicts).Error

	return err, sysDicts, total
}

// GetAll
// @function: GetAll
// @description: 获取全部字典列表
// @return: err error, list interface{}, total int64
func (dictService *DictService) GetAll() (err error, list interface{}) {
	db := global.GormDB.Model(&system.SysDict{})
	var sysDicts []system.SysDict

	err = db.Preload("DictDetails").Find(&sysDicts).Error
	return err, sysDicts
}

// SaveDetail
// @function: SaveDetail
// @description: 更新字典数据
// @param: req *request.SaveDetailRequest
// @return: err error
func (dictService *DictService) SaveDetail(req *request.SaveDetailRequest) (err error) {
	var dict system.SysDict

	err = global.GormDB.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Where("id = ? ", req.DictId).Preload("DictDetails").First(&dict).Error
		if txErr != nil {
			return txErr
		}

		txErr = tx.Unscoped().Delete(&system.SysDictDetail{}, "dict_id = ?", req.DictId).Error
		if txErr != nil {
			global.Debug("SaveDetail", txErr.Error())
			return txErr
		}

		var details []system.SysDictDetail

		status := true
		for _, detail := range req.DictDetails {
			d := system.SysDictDetail{
				DictId: req.DictId,
				Label:  detail.Label,
				Value:  detail.Value,
				Color:  detail.Color,
				Status: &status,
			}
			details = append(details, d)
		}

		err := tx.Create(&details).Error

		return err
	})

	return err
}
