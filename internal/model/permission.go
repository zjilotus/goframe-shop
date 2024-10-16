package model

import "github.com/gogf/gf/v2/os/gtime"

type PermissionCreateUpdateBase struct {
	Name string
	Path string
}

type PermissionCreateInput struct {
	PermissionCreateUpdateBase
}

type PermissionCreateOutput struct {
	PermissionId uint `json:"permission_id"`
}

type PermissionUpdateInput struct {
	PermissionCreateUpdateBase
	Id uint `json:"id"`
}

type PermissionGetListInput struct {
	Page int
	Size int
	Sort int
}

type PermissionGetListOutput struct {
	List  []PermissionGetListOutputItem `json:"list" description:"列表"`
	Page  int                           `json:"page" description:"分页码"`
	Size  int                           `json:"size" description:"分页数量"`
	Total int                           `json:"total" description:"数据总数"`
}

type PermissionGetListOutputItem struct {
	Id        uint        `json:"id"` // 自增ID
	Name      string      `json:"name"`
	Path      string      `json:"path"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
type PermissionSearchOutputItem struct {
	PermissionGetListOutputItem
}
