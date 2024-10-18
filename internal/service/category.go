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
	ICategory interface {
		Create(ctx context.Context, in model.CategoryCreateInput) (out model.CategoryCreateOutput, err error)
		Delete(ctx context.Context, id uint) (err error)
		Update(ctx context.Context, in model.CategoryUpdateInput) error
		// GetList 查询分类列表
		GetList(ctx context.Context, in model.CategoryGetListInput) (out *model.CategoryGetListOutput, err error)
		GetListAll(ctx context.Context, in model.CategoryGetListInput) (out *model.CategoryGetListOutput, err error)
	}
)

var (
	localCategory ICategory
)

func Category() ICategory {
	if localCategory == nil {
		panic("implement not found for interface ICategory, forgot register?")
	}
	return localCategory
}

func RegisterCategory(i ICategory) {
	localCategory = i
}
