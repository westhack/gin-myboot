package response

import "gin-myboot/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
