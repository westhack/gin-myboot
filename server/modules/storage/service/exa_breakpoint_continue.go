package service

import (
	"errors"
	"gin-myboot/global"
	storage "gin-myboot/modules/storage/model"
	"gorm.io/gorm"
)

type UploadFileService struct {
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: FindOrCreateFile
//@description: 上传文件时检测当前文件属性，如果没有文件则创建，有则返回文件的当前切片
//@param: fileMd5 string, fileName string, chunkTotal int
//@return: err error, file model.ExaFile

func (e *UploadFileService) FindOrCreateFile(fileMd5 string, fileName string, chunkTotal int) (err error, file storage.ExaFile) {
	var cfile storage.ExaFile
	cfile.FileMd5 = fileMd5
	cfile.FileName = fileName
	cfile.ChunkTotal = chunkTotal

	if errors.Is(global.GormDB.Where("file_md5 = ? AND is_finish = ?", fileMd5, true).First(&file).Error, gorm.ErrRecordNotFound) {
		err = global.GormDB.Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).Preload("ExaFileChunk").FirstOrCreate(&file, cfile).Error
		return err, file
	}
	cfile.IsFinish = true
	cfile.FilePath = file.FilePath
	err = global.GormDB.Create(&cfile).Error
	return err, cfile
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateFileChunk
//@description: 创建文件切片记录
//@param: id uint, fileChunkPath string, fileChunkNumber int
//@return: error

func (e *UploadFileService) CreateFileChunk(id uint64, fileChunkPath string, fileChunkNumber int) error {
	var chunk storage.ExaFileChunk
	chunk.FileChunkPath = fileChunkPath
	chunk.ExaFileID = id
	chunk.FileChunkNumber = fileChunkNumber
	err := global.GormDB.Create(&chunk).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除文件切片记录
//@param: fileMd5 string, fileName string, filePath string
//@return: error

func (e *UploadFileService) DeleteFileChunk(fileMd5 string, fileName string, filePath string) error {
	var chunks []storage.ExaFileChunk
	var file storage.ExaFile
	err := global.GormDB.Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).First(&file).Update("IsFinish", true).Update("file_path", filePath).Error
	err = global.GormDB.Where("exa_file_id = ?", file.ID).Delete(&chunks).Unscoped().Error
	return err
}
