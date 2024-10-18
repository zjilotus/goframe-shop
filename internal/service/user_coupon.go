// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"goframe-shop/internal/model"

	"golang.org/x/net/context"
)

type (
	IUserCoupon interface {
		Create(ctx context.Context, in model.UserCouponCreateInput) (out model.UserCouponCreateOutput, err error)
		Delete(ctx context.Context, id uint) (err error)
		Update(ctx context.Context, in model.UserCouponUpdateInput) (err error)
		// GetList 查询分类列表
		GetList(ctx context.Context, in model.UserCouponGetListInput) (out *model.UserCouponGetListOutput, err error)
	}
)

var (
	localUserCoupon IUserCoupon
)

func UserCoupon() IUserCoupon {
	if localUserCoupon == nil {
		panic("implement not found for interface IUserCoupon, forgot register?")
	}
	return localUserCoupon
}

func RegisterUserCoupon(i IUserCoupon) {
	localUserCoupon = i
}
