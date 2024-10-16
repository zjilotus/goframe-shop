// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe-shop/internal/model"
)

type (
	IPermission interface {
		Create(ctx context.Context, in model.PermissionCreateInput) (out model.PermissionCreateOutput, err error)
		Delete(ctx context.Context, id uint) error
		Update(ctx context.Context, in model.PermissionUpdateInput) error
		GetList(ctx context.Context, in model.PermissionGetListInput) (out *model.PermissionGetListOutput, err error)
	}
)

var (
	localPermission IPermission
)

func Permission() IPermission {
	if localPermission == nil {
		panic("implement not found for interface IPermission, forgot register?")
	}
	return localPermission
}

func RegisterPermission(i IPermission) {
	localPermission = i
}
