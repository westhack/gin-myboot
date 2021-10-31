package service

import (
	"gin-myboot/modules/common/model"
	"gin-myboot/modules/common/model/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseServiceInterface interface {
	Create(req map[string]interface{}) (err error, ret interface{})
	Update(req map[string]interface{}) (err error, ret interface{})
	Find(id uint64) (err error, ret interface{})
	GetList(queryParams request.QueryParams) (err error, ret interface{})
	GetAll() (err error, ret interface{})
	ChangeStatus(id []uint64, status interface{}) (err error, ret interface{})
	DeleteByIds(id []uint64) (err error, ret interface{})
	Delete(id uint64) (err error, ret interface{})
}


type BaseService struct {
	TableName string
	Model    Model
}

type Model struct {
	AllLimit   int
	IdName     string
	Preloads   []string
	StatusName string
	GetModel   func() (db *gorm.DB)
}

func (this *BaseService) Create(req map[string]interface{}) (err error, ret interface{}) {
	db := this.Model.GetModel()

	err = db.Create(req).Error

	return err, req
}

func (this *BaseService) Update(req map[string]interface{}) (err error, ret interface{}) {
	db := this.Model.GetModel()

	idName := this.Model.IdName
	err = db.Where(this.Model.IdName+" = ?", req[idName]).Updates(req).Error

	return err, req
}

func (this *BaseService) Find(id uint64) (err error, ret interface{}) {
	db := this.Model.GetModel()

	result := map[string]interface{}{}
	err = db.First(&result).Error

	return err, result
}

func (this *BaseService) GetList(queryParams request.QueryParams) (err error, ret interface{}) {
	db := this.Model.GetModel()

	var results []map[string]interface{}

	var total int64

	limit := queryParams.PageSize
	offset := queryParams.PageSize * (queryParams.Page - 1)

	err = db.Scopes(model.Search(queryParams.Search)).
		Scopes(model.SortOrder(queryParams.SortOrder)).Count(&total).Error

	err = db.Scopes(model.Search(queryParams.Search)).
		Scopes(model.SortOrder(queryParams.SortOrder)).
		Limit(limit).Offset(offset).Find(&results).Error

	ret = gin.H{
		"total":    total,
		"items":    results,
		"pageSize": limit,
		"page":     queryParams.Page,
	}

	return err, ret
}

func (this *BaseService) GetAll() (err error, ret interface{}) {
	db := this.Model.GetModel()

	AllLimit := this.Model.AllLimit

	if AllLimit > 0 {
		db.Limit(AllLimit)
	}

	var results []map[string]interface{}
	err = db.Find(&results).Error

	return err, results
}

func (this *BaseService) ChangeStatus(id []uint64, status interface{}) (err error, ret interface{}) {
	db := this.Model.GetModel()

	idName := this.Model.IdName
	err = db.Where(idName + " = ? ", id).Update(this.Model.StatusName, status).Error

	return err, status
}

func (this *BaseService) Delete(id uint64) (err error, ret interface{}) {
	db := this.Model.GetModel()

	idName := this.Model.IdName
	var result interface{}
	err = db.Unscoped().Where(idName + " = ? ", id).Delete(&result).Error

	return err, result
}

func (this *BaseService) DeleteByIds(id []uint64) (err error, ret interface{}) {
	db := this.Model.GetModel()

	idName := this.Model.IdName
	var result interface{}
	err = db.Unscoped().Where(idName+" in ? ", id).Delete(&result).Error

	return err, result
}
