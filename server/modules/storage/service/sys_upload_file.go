package service

import (
	"errors"
	"gin-myboot/global"
	"gin-myboot/modules/common/model"
	"gin-myboot/modules/common/model/request"
	storage "gin-myboot/modules/storage/model"
	"gin-myboot/utils/upload"
	"mime/multipart"
	"strings"
)

// Upload 
// @function: Upload
// @description: 创建文件上传记录
// @param: file model.SysUploadFile
// @return: error
func (e *UploadFileService) Upload(file storage.SysUploadFile, userId uint64) error {
	file.Type = global.Config.System.OssType
	file.UserId = userId
	return global.GormDB.Create(&file).Error
}

// FindFile
// @function: FindFile
// @description: 删除文件切片记录
// @param: id uint
// @return: error, model.SysUploadFile
func (e *UploadFileService) FindFile(id uint64) (error, storage.SysUploadFile) {
	var file storage.SysUploadFile
	err := global.GormDB.Where("id = ?", id).First(&file).Error
	return err, file
}

// DeleteFile
// @function: DeleteFile
// @description: 删除文件记录
// @param: file model.SysUploadFile
// @return: err error
func (e *UploadFileService) DeleteFile(file storage.SysUploadFile) (err error) {
	var fileFromDb storage.SysUploadFile
	err, fileFromDb = e.FindFile(file.ID)
	oss := upload.NewOssByType(fileFromDb.Type)
	if err = oss.DeleteFile(fileFromDb.Uuid); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.GormDB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return err
}

// DeleteFileByIds
// @function: DeleteFileByIds
// @description: 删除文件记录
// @param: ids request.GetByIds
// @return: err error
func (e *UploadFileService) DeleteFileByIds(req request.GetByIds) (err error) {

	var files []storage.SysUploadFile
	err = global.GormDB.Where("id in ?", req.ID).Unscoped().Find(&files).Error

	global.Error("=====>", files)
	for _, file := range files {
		oss := upload.NewOssByType(file.Type)
		if err = oss.DeleteFile(file.Uuid); err != nil {
			global.Error("====>NewOss DeleteFile", err)
			return errors.New("文件删除失败：" + file.Name)
		}
		err = global.GormDB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
		if err != nil {
			return errors.New("文件删除失败：" + file.Name)
		}
	}

	return err
}


// CreateFile
// @function: CreateFile
// @description: 添加文件
// @param: file storage.SysUploadFile
// @return: err error
func (e *UploadFileService) CreateFile(file storage.SysUploadFile) (err error) {

	err = global.GormDB.Where("id =?", file.ID).Updates(&file).Error

	return err
}


// UpdateFile
// @function: CreateFile
// @description: 修改文件
// @param: file storage.SysUploadFile
// @return: err error
func (e *UploadFileService) UpdateFile(file storage.SysUploadFile) (err error) {

	err = global.GormDB.Create(&file).Error

	return err
}

// GetFileList
// @function: GetFileList
// @description: 分页获取数据
// @param: info request.QueryParams
// @return: err error, list interface{}, total int64
func (e *UploadFileService) GetFileList(queryParams request.QueryParams) (err error, list interface{}, total int64) {
	limit := queryParams.PageSize
	offset := queryParams.PageSize * (queryParams.Page - 1)
	db := global.GormDB
	var fileLists []storage.SysUploadFile

	db.Scopes(model.Search(queryParams.Search)).Scopes(model.SortOrder(queryParams.SortOrder))

	err = db.Find(&fileLists).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return err, fileLists, total
}

// UploadFile
// @function: UploadFile
// @description: 根据配置文件判断是文件上传到本地或者七牛云
// @param: header *multipart.FileHeader, noSave string, userId uint
// @return: err error, file model.SysUploadFile
func (e *UploadFileService) UploadFile(header *multipart.FileHeader, noSave string, userId uint64) (err error, file storage.SysUploadFile) {
	oss := upload.NewOss()
	filePath, uuid, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := storage.SysUploadFile{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Uuid:  uuid,
		}
		return e.Upload(f, userId), f
	}
	return
}


// GetUserFileList
// @function: GetUserFileList
// @description: 分页获取数据
// @param: info request.QueryParams
// @return: err error, list interface{}, total int64
func (e *UploadFileService) GetUserFileList(queryParams request.QueryParams, userId uint64) (err error, list interface{}, total int64) {
	limit := queryParams.PageSize
	offset := queryParams.PageSize * (queryParams.Page - 1)
	db := global.GormDB
	var fileLists []storage.SysUploadFile

	db.Where("user_id = ?", userId)
	db.Scopes(model.Search(queryParams.Search)).Scopes(model.SortOrder(queryParams.SortOrder))

	err = db.Find(&fileLists).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return err, fileLists, total
}
