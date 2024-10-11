package model

import "github.com/gogf/gf/v2/os/gtime"

type PositionCreateUpdateBase struct {
	PicUrl    string
	Link      string
	Sort      int
	GoodsName string
	GoodsId   uint
}

type PositionCreateInput struct {
	PositionCreateUpdateBase
}

type PositionCreateOutput struct {
	PositionId int `json:"position_id"`
}

type PositionUpdateInput struct {
	PositionCreateUpdateBase
	Id uint
}

type PositionGetListInput struct {
	Page int
	Size int
	Sort int
}

type PositionGetListOutput struct {
	List  []PositionGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int                         `json:"total" description:"数据总数"`
}

// PositionSearchInput 搜索列表
type PositionSearchInput struct {
	Key        string // 关键字
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// PositionSearchOutput 搜索列表结果
type PositionSearchOutput struct {
	List  []PositionSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int             `json:"stats"` // 搜索统计
	Page  int                        `json:"page"`  // 分页码
	Size  int                        `json:"size"`  // 分页数量
	Total int                        `json:"total"` // 数据总数
}

type PositionGetListOutputItem struct {
	//Position *PositionListItem `json:"position"`
	Id        uint        `json:"id"`
	PicUrl    string      `json:"pic_url"`
	Link      string      `json:"link"`
	GoodsName string      `json:"goods_name"`
	GoodsId   string      `json:"goods_id"`
	Sort      int         `json:"sort"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}

type PositionSearchOutputItem struct {
	PositionGetListOutputItem
}

//type PositionListItem struct {
//	Id        uint        `json:"id"`
//	PicUrl    string      `json:"pic_url"`
//	Link      string      `json:"link"`
//	Sort      int         `json:"sort"`
//	CreatedAt *gtime.Time `json:"created_at"`
//	UpdatedAt *gtime.Time `json:"updated_at"`
//}
