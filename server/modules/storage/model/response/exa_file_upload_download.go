package response

import (
	storage "gin-myboot/modules/storage/model"
)

type ExaFileResponse struct {
	File storage.SysUploadFile `json:"file"`
}
