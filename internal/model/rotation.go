package model

import "github.com/gogf/gf/v2/os/gtime"

type RotationCreateUpdateBase struct {
	PicUrl string
	Link   string
	Sort   int
}

type RotationCreateInput struct {
	RotationCreateUpdateBase
}

type RotationCreateOutput struct {
	RotationId int `json:"rotation_id"`
}

type RotationUpdateInput struct {
	RotationCreateUpdateBase
	Id uint
}

type RotationGetListInput struct {
	Page int
	Size int
	Sort int
}

type RotationGetListOutput struct {
	List  []RotationGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int                         `json:"total" description:"数据总数"`
}

// RotationSearchInput 搜索列表
type RotationSearchInput struct {
	Key        string // 关键字
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// RotationSearchOutput 搜索列表结果
type RotationSearchOutput struct {
	List  []RotationSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int             `json:"stats"` // 搜索统计
	Page  int                        `json:"page"`  // 分页码
	Size  int                        `json:"size"`  // 分页数量
	Total int                        `json:"total"` // 数据总数
}

type RotationGetListOutputItem struct {
	//Rotation *RotationListItem `json:"rotation"`
	Id        uint        `json:"id"`
	PicUrl    string      `json:"pic_url"`
	Link      string      `json:"link"`
	Sort      int         `json:"sort"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}

type RotationSearchOutputItem struct {
	RotationGetListOutputItem
}

//type RotationListItem struct {
//	Id        uint        `json:"id"`
//	PicUrl    string      `json:"pic_url"`
//	Link      string      `json:"link"`
//	Sort      int         `json:"sort"`
//	CreatedAt *gtime.Time `json:"created_at"`
//	UpdatedAt *gtime.Time `json:"updated_at"`
//}
