package response

import system "gin-myboot/modules/system/model"

type SysRoleResponse struct {
	Role system.SysRole `json:"role"`
}

type SysRoleCopyResponse struct {
	Role      system.SysRole `json:"role"`
	OldRoleId uint              `json:"oldRoleId"` // 旧角色ID
}
