package request

// 权限搜索条件

type SearchRoleParams struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

type RoleFormRequest struct {
	Id          uint64   `json:"id"`
	Title       string   `json:"title" binding:"required" required_err:"名称不能为空"`
	Name        string   `json:"name" binding:"required" required_err:"角色名不能为空"`
	Description string   `json:"description"`
	DefaultRole *bool    `json:"defaultRole"`
	ParentId    uint64   `json:"parentId"`
	Permissions []uint64 `json:"permissions"`
	Departments []uint64 `json:"departments"`
}
