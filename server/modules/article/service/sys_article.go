package service

import (
    "gin-myboot/global"
    article "gin-myboot/modules/article/model"
    "gin-myboot/modules/common/model"
    "gin-myboot/modules/common/model/request"
)

type SysArticleService struct {
}

var SysArticleServiceApp = new(SysArticleService)

// Create 创建SysArticle记录
func (articleService *SysArticleService) Create(article article.SysArticle) (err error) {
    err = global.GormDB.Create(&article).Error
    return err
}

// Delete 删除SysArticle记录
func (articleService *SysArticleService)Delete(id uint64) (err error) {
    err = global.GormDB.Where("id = ?", id).Delete(&article.SysArticle{}).Error
    return err
}

// DeleteByIds 批量删除SysArticle记录
func (articleService *SysArticleService)DeleteByIds(id []uint64) (err error) {
    err = global.GormDB.Delete(&[]article.SysArticle{},"id in ?",id).Error
    return err
}

// UpdateSysArticle 更新SysArticle记录
func (articleService *SysArticleService)Update(article article.SysArticle) (err error) {
    err = global.GormDB.Updates(&article).Error
    return err
}

// FindById 根据id获取SysArticle记录
func (articleService *SysArticleService)FindById(id uint64) (err error, article article.SysArticle) {
    err = global.GormDB.Where("id = ?", id).First(&article).Error
    return
}

// GetList 分页获取SysArticle记录
func (articleService *SysArticleService)GetList(queryParams request.QueryParams) (err error, list interface{}, total int64) {
    limit := queryParams.PageSize
    offset := queryParams.PageSize * (queryParams.Page - 1)

    if queryParams.SortOrder.Column == "" {
        queryParams.SortOrder.Column = "id"
        queryParams.SortOrder.Order = "desc"
    }

    // 创建db
    db := global.GormDB.Model(&article.SysArticle{}).Scopes(model.Search(queryParams.Search))
    var articles []article.SysArticle
    err = db.Count(&total).Error
    err = db.Scopes(model.SortOrder(queryParams.SortOrder)).Limit(limit).Offset(offset).Find(&articles).Error
    return err, articles, total
}
