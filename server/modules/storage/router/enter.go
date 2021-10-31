package router

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	CustomerRouter
	ExcelRouter
	UploadFileAndRouter
	SimpleUploaderRouter
}

var StorageRouterGroupApp = new(RouterGroup)

func InitStorageRouter(Router *gin.RouterGroup) {
	StorageRouterGroupApp.InitCustomerRouter(Router)       // 客户路由
	StorageRouterGroupApp.InitExcelRouter(Router)          // 表格导入导出
	StorageRouterGroupApp.InitSimpleUploaderRouter(Router) // 断点续传（插件版）
	StorageRouterGroupApp.InitUploadFileAndRouter(Router)  // // 文件上传下载功能路由
}
