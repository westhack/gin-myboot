package router

import (
	"gin-myboot/middleware"
	"gin-myboot/modules/storage/api"
	"github.com/gin-gonic/gin"
)

type UploadFileAndRouter struct {
}

func (e *UploadFileAndRouter) InitUploadFileAndRouter(Router *gin.RouterGroup) {
	uploadFileRouter := Router.Group("v1/storage/upload").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	var uploadFileApi = api.StorageApiGroupApp.UploadFileApi
	{
		uploadFileRouter.POST("/file", uploadFileApi.UploadFile)                                   // 上传文件
		uploadFileRouter.POST("/getList", uploadFileApi.GetFileList)                               // 获取上传文件列表
		uploadFileRouter.POST("/getUserFileList", uploadFileApi.GetUserFileList)                   // 获取用户上传文件列表
		uploadFileRouter.POST("/delete", uploadFileApi.DeleteFile)                                 // 删除指定文件
		uploadFileRouter.POST("/deleteByIds", uploadFileApi.DeleteFileByIds)                       // 删除指定文件
		uploadFileRouter.POST("/create", uploadFileApi.CreateFile)                                 // 添加文件
		uploadFileRouter.POST("/update", uploadFileApi.UpdateFile)                                 // 修改文件
		uploadFileRouter.POST("/breakpointContinue", uploadFileApi.BreakpointContinue)             // 断点续传
		uploadFileRouter.GET("/findFile", uploadFileApi.FindFile)                                  // 查询当前文件成功的切片
		uploadFileRouter.POST("/breakpointContinueFinish", uploadFileApi.BreakpointContinueFinish) // 查询当前文件成功的切片
		uploadFileRouter.POST("/removeChunk", uploadFileApi.RemoveChunk)                           // 查询当前文件成功的切片
	}
}
