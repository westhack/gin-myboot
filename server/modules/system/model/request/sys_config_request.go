package request

// 权限搜索条件

type ConfigSearch struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

type UpdateConfigRequest struct {
	Id          int    `json:"id" binding:"required" required_err:"请选择要修改的记录"`
	Title       string `json:"title" binding:"required" required_err:"名称不能为空"`
	Name        string `json:"name" binding:"required" required_err:"角色名不能为空"`
	ParentId    int    `json:"parentId"`
	Permissions []int  `json:"permissions"`
}

type CreateConfigRequest struct {
	Title       string `json:"title" binding:"required" required_err:"名称不能为空"`
	Name        string `json:"name" binding:"required" required_err:"角色名不能为空"`
	ParentId    int    `json:"parentId"`
	Permissions []int  `json:"permissions"`
}

type ConfigValueRequest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
