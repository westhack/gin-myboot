package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

// GetById Find by id structure
type GetById struct {
	ID uint64 `json:"id" form:"id"` // 主键ID
}

type GetByIds struct {
	ID []uint64 `json:"id" form:"id"` // 主键ID
}

type IdsReq struct {
	Ids []uint64 `json:"ids" form:"ids"`
}

// GetRoleId Get role by id structure
type GetRoleId struct {
	RoleId string `json:"roleId" form:"roleId"` // 角色ID
}

type Empty struct{}

// QueryParams 通用搜索结构体
type QueryParams struct {
	PageInfo
	Search    []SearchParams  `json:"search"`
	SortOrder SortOrderParams `json:"sortOrder"`
}

// SearchParams 通用搜索结构体
type SearchParams struct {
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

// SortOrderParams 通用搜索结构体
type SortOrderParams struct {
	Column string `json:"column"`
	Order  string `json:"order"`
}

type ChangeStatus struct {
	ID     []uint64 `json:"id" form:"id"` // 主键ID
	Status int      `json:"status" form:"status"`
}
