package model

import "github.com/gogf/gf/v2/os/gtime"

type AdminCreateUpdateBase struct {
	Name     string
	Password string
	RoleIds  string
	UserSalt string
	IsAmin   int
}

type AdminCreateInput struct {
	AdminCreateUpdateBase
}

type AdminCreateOutput struct {
	AdminId int `json:"admin_id"`
}

type AdminUpdateInput struct {
	AdminCreateUpdateBase
	Id uint
}

type AdminGetListInput struct {
	Page int
	Size int
	Sort int
}

type AdminGetListOutput struct {
	List  []AdminGetListOutputItem `json:"list" description:"列表"`
	Page  int                      `json:"page" description:"分页码"`
	Size  int                      `json:"size" description:"分页数量"`
	Total int                      `json:"total" description:"数据总数"`
}

//// AdminSearchInput 搜索列表
//type AdminSearchInput struct {
//	Key        string // 关键字
//	Type       string // 内容模型
//	CategoryId uint   // 栏目ID
//	Page       int    // 分页号码
//	Size       int    // 分页数量，最大50
//	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
//}
//
//// AdminSearchOutput 搜索列表结果
//type AdminSearchOutput struct {
//	List  []AdminSearchOutputItem `json:"list"`  // 列表
//	Stats map[string]int          `json:"stats"` // 搜索统计
//	Page  int                     `json:"page"`  // 分页码
//	Size  int                     `json:"size"`  // 分页数量
//	Total int                     `json:"total"` // 数据总数
//}

type AdminGetListOutputItem struct {
	//Admin *AdminListItem `json:"admin"`
	Id        uint        `json:"id"`
	Name      string      `json:"name"`
	RoleIds   string      `json:"role_ids"`
	IsAdmin   int         `json:"is_admin"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}

type AdminSearchOutputItem struct {
	AdminGetListOutputItem
}

//type AdminListItem struct {
//	Id        uint        `json:"id"`
//	PicUrl    string      `json:"pic_url"`
//	Link      string      `json:"link"`
//	Sort      int         `json:"sort"`
//	CreatedAt *gtime.Time `json:"created_at"`
//	UpdatedAt *gtime.Time `json:"updated_at"`
//}
