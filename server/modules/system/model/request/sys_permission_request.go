package request

// 权限搜索条件

type SearchPermissionParams struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

type PermissionFormRequest struct {
	Id          uint64   `json:"id"`
	ParentId    uint64   `json:"parentId"`
	Title       string   `json:"title" binding:"required" required_err:"标题不能为空"`
	Name        string   `json:"name" binding:"required" required_err:"名称不能为空"`
	Icon        string   `json:"icon"`
	IsButton    bool     `json:"isButton"`
	Component   string   `json:"component"`
	Hidden      bool     `json:"hidden"`
	Path        string   `json:"path"`
	DefaultPath string   `json:"defaultPath"`
	Redirect    string   `json:"redirect"`
	Api         string   `json:"api"`
	SortOrder   float64  `json:"sortOrder"`
	KeepAlive   bool     `json:"keepAlive"`
	Status      bool     `json:"status"`
	Buttons     []string `json:"buttons"`
}
