package model

import "github.com/gogf/gf/v2/os/gtime"

type UserCouponCreateUpdateBase struct {
	UserId   uint
	CouponId uint
	Status   uint8
}

type UserCouponCreateInput struct {
	UserCouponCreateUpdateBase
}

type UserCouponCreateOutput struct {
	Id uint `json:"id"`
}

type UserCouponUpdateInput struct {
	UserCouponCreateUpdateBase
	Id uint
}

// UserCouponGetListInput 获取内容列表
type UserCouponGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// UserCouponGetListOutput 查询列表结果
type UserCouponGetListOutput struct {
	List  []UserCouponGetListOutputItem `json:"list" description:"列表"`
	Page  int                           `json:"page" description:"分页码"`
	Size  int                           `json:"size" description:"分页数量"`
	Total int                           `json:"total" description:"数据总数"`
}
type UserCouponGetListOutputItem struct {
	Id        uint        `json:"id"` // 自增ID
	UserId    uint        `json:"user_id"`
	CouponId  uint        `json:"coupon_id"`
	Status    uint8       `json:"status"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
