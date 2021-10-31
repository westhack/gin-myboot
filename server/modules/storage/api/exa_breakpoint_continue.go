package api

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/response"
	storageRes "gin-myboot/modules/storage/model/response"
	"gin-myboot/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"strconv"
)

// @Tags SysUploadFile
// @Summary 断点续传到服务器
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "an example for breakpoint resume, 断点续传示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"切片创建成功"}"
// @Router /fileUploadAndDownload/breakpointContinue [post]
func (u *UploadFileApi) BreakpointContinue(c *gin.Context) {
	fileMd5 := c.Request.FormValue("fileMd5")
	fileName := c.Request.FormValue("fileName")
	chunkMd5 := c.Request.FormValue("chunkMd5")
	chunkNumber, _ := strconv.Atoi(c.Request.FormValue("chunkNumber"))
	chunkTotal, _ := strconv.Atoi(c.Request.FormValue("chunkTotal"))
	_, FileHeader, err := c.Request.FormFile("file")
	if err != nil {
		global.Logger.Error("接收文件失败!", zap.Any("err", err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	f, err := FileHeader.Open()
	if err != nil {
		global.Logger.Error("文件读取失败!", zap.Any("err", err))
		response.FailWithMessage("文件读取失败", c)
		return
	}
	defer f.Close()
	cen, _ := ioutil.ReadAll(f)
	if !utils.CheckMd5(cen, chunkMd5) {
		global.Logger.Error("检查md5失败!", zap.Any("err", err))
		response.FailWithMessage("检查md5失败", c)
		return
	}
	err, file := uploadFileService.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		global.Logger.Error("查找或创建记录失败!", zap.Any("err", err))
		response.FailWithMessage("查找或创建记录失败", c)
		return
	}
	err, pathc := utils.BreakPointContinue(cen, fileName, chunkNumber, chunkTotal, fileMd5)
	if err != nil {
		global.Logger.Error("断点续传失败!", zap.Any("err", err))
		response.FailWithMessage("断点续传失败", c)
		return
	}

	if err = uploadFileService.CreateFileChunk(file.ID, pathc, chunkNumber); err != nil {
		global.Logger.Error("创建文件记录失败!", zap.Any("err", err))
		response.FailWithMessage("创建文件记录失败", c)
		return
	}
	response.OkWithMessage("切片创建成功", c)
}

// @Tags SysUploadFile
// @Summary 查找文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "Find the file, 查找文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查找成功"}"
// @Router /fileUploadAndDownload/findFile [post]
func (u *UploadFileApi) FindFile(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	chunkTotal, _ := strconv.Atoi(c.Query("chunkTotal"))
	err, file := uploadFileService.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		global.Logger.Error("查找失败!", zap.Any("err", err))
		response.FailWithMessage("查找失败", c)
	} else {
		response.OkWithDetailed(storageRes.FileResponse{File: file}, "查找成功", c)
	}
}

// @Tags SysUploadFile
// @Summary 创建文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件完成"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"file uploaded, 文件创建成功"}"
// @Router /fileUploadAndDownload/findFile [post]
func (b *UploadFileApi) BreakpointContinueFinish(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	err, filePath := utils.MakeFile(fileName, fileMd5)
	if err != nil {
		global.Logger.Error("文件创建失败!", zap.Any("err", err))
		response.FailWithDetailed(storageRes.FilePathResponse{FilePath: filePath}, "文件创建失败", c)
	} else {
		response.OkWithDetailed(storageRes.FilePathResponse{FilePath: filePath}, "文件创建成功", c)
	}
}

// @Tags SysUploadFile
// @Summary 删除切片
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "删除缓存切片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"缓存切片删除成功"}"
// @Router /fileUploadAndDownload/removeChunk [post]
func (u *UploadFileApi) RemoveChunk(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	filePath := c.Query("filePath")
	err := utils.RemoveChunk(fileMd5)
	err = uploadFileService.DeleteFileChunk(fileMd5, fileName, filePath)
	if err != nil {
		global.Logger.Error("缓存切片删除失败!", zap.Any("err", err))
		response.FailWithDetailed(storageRes.FilePathResponse{FilePath: filePath}, "缓存切片删除失败", c)
	} else {
		response.OkWithDetailed(storageRes.FilePathResponse{FilePath: filePath}, "缓存切片删除成功", c)
	}
}
