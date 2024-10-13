package frontend

import "github.com/gogf/gf/v2/frame/g"

type RotationGetListCommonReq struct {
	g.Meta `path:"/frontend/rotation/list" method:"get" tags:"轮播图" summary:"轮播图接口"`
	Sort   int `json:"sort" in:"query" dc:"排序"`
	CommonPaginationReq
}

type RotationGetListCommonRes struct {
	List  interface{} `json:"list" dc:"列表"`
	Page  int         `json:"page" dc:"分页码"`
	Size  int         `json:"size" dc:"分页数量"`
	Total int         `json:"total" dc:"数据总数"`
}
