package router

import (
	"gin-myboot/modules/storage/api"
	"github.com/gin-gonic/gin"
)

type ExcelRouter struct {
}

func (e *ExcelRouter) InitExcelRouter(Router *gin.RouterGroup) {
	excelRouter := Router.Group("excel")
	var exaExcelApi = api.StorageApiGroupApp.ExcelApi
	{
		excelRouter.POST("/importExcel", exaExcelApi.ImportExcel)          // 导入Excel
		excelRouter.GET("/loadExcel", exaExcelApi.LoadExcel)               // 加载Excel数据
		excelRouter.POST("/exportExcel", exaExcelApi.ExportExcel)          // 导出Excel
		excelRouter.GET("/downloadTemplate", exaExcelApi.DownloadTemplate) // 下载模板文件
	}
}
