package api

import (
	"gin-myboot/global"
	"gin-myboot/modules/common/model/request"
	"gin-myboot/modules/common/model/response"
	storage "gin-myboot/modules/storage/model"
	storageRes "gin-myboot/modules/storage/model/response"
	"gin-myboot/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UploadFileApi struct {
}

// UploadFile
// @Tags SysUploadFile
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /v1/storage/upload/file [post]
func (u *UploadFileApi) UploadFile(c *gin.Context) {
	var file storage.SysUploadFile
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.Logger.Error("接收文件失败!", zap.Any("err", err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	userId := utils.GetUserID(c)
	err, file = uploadFileService.UploadFile(header, noSave, userId) // 文件上传后拿到文件路径
	if err != nil {
		global.Logger.Error("修改数据库链接失败!", zap.Any("err", err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(storageRes.ExaFileResponse{File: file}, "上传成功", c)
}

// DeleteFile 
// @Tags SysUploadFile
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body storage.SysUploadFile true "传入文件里面id即可"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v1/storage/upload/delete [post]
func (u *UploadFileApi) DeleteFile(c *gin.Context) {
	var file storage.SysUploadFile
	_ = c.ShouldBindJSON(&file)
	if err := uploadFileService.DeleteFile(file); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteFileByIds
// @Tags SysUploadFile
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body storage.SysUploadFile true "传入文件里面id即可"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /upload/deleteByIds [post]
func (u *UploadFileApi) DeleteFileByIds(c *gin.Context) {
	var req request.GetByIds
	_ = c.ShouldBindJSON(&req)
	if err := uploadFileService.DeleteFileByIds(req); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetFileList
// @Tags SysUploadFile
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/storage/upload/getList [post]
func (u *UploadFileApi) GetFileList(c *gin.Context) {
	var pageInfo request.QueryParams
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := uploadFileService.GetFileList(pageInfo)
	if err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}


// GetUserFileList
// @Tags SysUploadFile
// @Summary 分页获取用户上传文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.QueryParams true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/storage/upload/getUserFileList [post]
func (u *UploadFileApi) GetUserFileList(c *gin.Context) {
	var pageInfo request.QueryParams
	_ = c.ShouldBindJSON(&pageInfo)

	userId := utils.GetUserID(c)

	err, list, total := uploadFileService.GetUserFileList(pageInfo, userId)
	if err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}


// CreateFile
// @Tags SysUploadFile
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body storage.SysUploadFile true "文件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /v1/storage/upload/create [post]
func (u *UploadFileApi) CreateFile(c *gin.Context) {
	var file storage.SysUploadFile
	_ = c.ShouldBindJSON(&file)
	if err := uploadFileService.CreateFile(file); err != nil {
		global.Logger.Error("添加失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}


// UpdateFile
// @Tags SysUploadFile
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body storage.SysUploadFile true "文件信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /v1/storage/upload/update [post]
func (u *UploadFileApi) UpdateFile(c *gin.Context) {
	var file storage.SysUploadFile
	_ = c.ShouldBindJSON(&file)
	if err := uploadFileService.UpdateFile(file); err != nil {
		global.Logger.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}