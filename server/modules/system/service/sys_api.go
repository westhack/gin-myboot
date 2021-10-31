package service

import (
	"errors"
	"gin-myboot/global"
	"gin-myboot/modules/common/model"
	"gin-myboot/modules/common/model/request"
	system "gin-myboot/modules/system/model"

	"gorm.io/gorm"
)

type ApiService struct {
}

var ApiServiceApp = new(ApiService)

// Create
// @function: Create
// @description: 新增基础api
// @param: api model.SysApi
// @return: err error
func (apiService *ApiService) Create(api system.SysApi) (err error) {
	if !errors.Is(global.GormDB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.GormDB.Create(&api).Error
}

// Delete
// @function: Delete
// @description: 删除基础api
// @param: api model.SysApi
// @return: err error
func (apiService *ApiService) Delete(api system.SysApi) (err error) {
	err = global.GormDB.Delete(&api).Error
	CasbinServiceApp.ClearCasbin(1, api.Path, api.Method)
	return err
}

// GetList
// @function: GetList
// @description: 分页获取数据,
// @param: queryParams request.QueryParams
// @return: err error
func (apiService *ApiService) GetList(queryParams request.QueryParams) (err error, list interface{}, total int64) {
	limit := queryParams.PageSize
	offset := queryParams.PageSize * (queryParams.Page - 1)
	db := global.GormDB.Model(&system.SysApi{})

	var apiList []system.SysApi

	db.Scopes(model.Search(queryParams.Search))

	err = db.Count(&total).Error

	if err != nil {
		return err, apiList, total
	} else {
		err = db.Scopes(model.SortOrder(queryParams.SortOrder)).
			Limit(limit).Offset(offset).Find(&apiList).Error

	}
	return err, apiList, total
}

// GetAll
// @function: GetAll
// @description: 获取所有的api
// @return: err error, apis []model.SysApi
func (apiService *ApiService) GetAll() (err error, apis []system.SysApi) {
	err = global.GormDB.Find(&apis).Error
	return
}

// GetById
// @function: GetById
// @description: 根据id获取api
// @param: id float64
// @return: err error, api model.SysApi
func (apiService *ApiService) GetById(id uint64) (err error, api system.SysApi) {
	err = global.GormDB.Where("id = ?", id).First(&api).Error
	return
}

// Update
// @function: Update
// @description: 根据id更新api
// @param: api model.SysApi
// @return: err error
func (apiService *ApiService) Update(api system.SysApi) (err error) {
	var oldA system.SysApi
	err = global.GormDB.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(global.GormDB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}

	oldA.Path = api.Path
	oldA.Method = api.Method
	oldA.ApiGroup = api.ApiGroup
	oldA.Description = api.Description

	if err != nil {
		return err
	} else {
		err = CasbinServiceApp.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = global.GormDB.Save(&oldA).Error
		}
	}
	return err
}

// DeleteByIds
// @function: DeleteByIds
// @description: 删除选中API
// @param: apis []model.SysApi
// @return: err error
func (apiService *ApiService) DeleteByIds(ids []uint64) (err error) {
	err = global.GormDB.Delete(&[]system.SysApi{}, "id in ?", ids).Error
	return err
}

func (apiService *ApiService) DeleteApisByIds(ids []string) (err error) {
	return global.GormDB.Delete(system.SysApi{}, ids).Error
}
