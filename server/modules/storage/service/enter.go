package service

type ServiceGroup struct {
	UploadFileService
	CustomerService
	ExcelService
	SimpleUploaderService
}

var StorageServiceGroup = new(ServiceGroup)
