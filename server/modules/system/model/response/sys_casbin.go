package response

import (
	"gin-myboot/modules/system/model/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
