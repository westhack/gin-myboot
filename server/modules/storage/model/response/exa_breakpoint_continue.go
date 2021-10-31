package response

import (
	storage "gin-myboot/modules/storage/model"
)

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File storage.ExaFile `json:"file"`
}
