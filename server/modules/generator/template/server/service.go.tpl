package service

import (
    "gin-myboot/global"
    {{.ModuleName}} "gin-myboot/modules/{{.ModuleName}}/model"
    "gin-myboot/modules/common/model"
    "gin-myboot/modules/common/model/request"
)

type {{.StructName}}Service struct {
}

var {{.StructName}}ServiceApp = new({{.StructName}}Service)

// Create 创建{{.StructName}}记录
func ({{.Abbreviation}}Service *{{.StructName}}Service) Create({{.Abbreviation}} {{.ModuleName}}.{{.StructName}}) (err error) {
    err = global.GormDB.Create(&{{.Abbreviation}}).Error
    return err
}

// Delete 删除{{.StructName}}记录
func ({{.Abbreviation}}Service *{{.StructName}}Service)Delete(id uint64) (err error) {
    err = global.GormDB.Where("id = ?", id).Delete(&{{.ModuleName}}.{{.StructName}}{}).Error
    return err
}

// DeleteByIds 批量删除{{.StructName}}记录
func ({{.Abbreviation}}Service *{{.StructName}}Service)DeleteByIds(id []uint64) (err error) {
    err = global.GormDB.Delete(&[]{{.PackageName}}.{{.StructName}}{},"id in ?",id).Error
    return err
}

// Update{{.StructName}} 更新{{.StructName}}记录
func ({{.Abbreviation}}Service *{{.StructName}}Service)Update({{.Abbreviation}} {{.ModuleName}}.{{.StructName}}) (err error) {
    err = global.GormDB.Updates(&{{.Abbreviation}}).Error
    return err
}

// FindById 根据id获取{{.StructName}}记录
func ({{.Abbreviation}}Service *{{.StructName}}Service)FindById(id uint64) (err error, {{.Abbreviation}} {{.ModuleName}}.{{.StructName}}) {
    err = global.GormDB.Where("id = ?", id).First(&{{.Abbreviation}}).Error
    return
}

// GetList 分页获取{{.StructName}}记录
func ({{.Abbreviation}}Service *{{.StructName}}Service)GetList(queryParams request.QueryParams) (err error, list interface{}, total int64) {
    limit := queryParams.PageSize
    offset := queryParams.PageSize * (queryParams.Page - 1)

    if queryParams.SortOrder.Column == "" {
        queryParams.SortOrder.Column = "id"
        queryParams.SortOrder.Order = "desc"
    }

    // 创建db
    db := global.GormDB.Model(&{{.ModuleName}}.{{.StructName}}{}).Scopes(model.Search(queryParams.Search))
    var {{.Abbreviation}}s []{{.ModuleName}}.{{.StructName}}
    err = db.Count(&total).Error
    err = db.Scopes(model.SortOrder(queryParams.SortOrder)).Limit(limit).Offset(offset).Find(&{{.Abbreviation}}s).Error
    return err, {{.Abbreviation}}s, total
}
