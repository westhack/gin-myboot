package api

import "gin-myboot/modules/storage/service"

type ApiGroup struct {
	CustomerApi
	ExcelApi
	UploadFileApi
	SimpleUploaderApi
}

var StorageApiGroupApp = new(ApiGroup)

var uploadFileService = service.StorageServiceGroup.UploadFileService
var customerService = service.StorageServiceGroup.CustomerService
var excelService = service.StorageServiceGroup.ExcelService
var simpleUploaderService = service.StorageServiceGroup.SimpleUploaderService
