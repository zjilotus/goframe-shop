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
	ICollection interface {
		AddCollection(ctx context.Context, in model.AddCollectionInput) (out *model.AddCollectionOutput, err error)
		// 兼容处理：优先根据收藏id删除，收藏id为0；再根据对象id和type删除
		DeleteCollection(ctx context.Context, in model.DeleteCollectionInput) (out *model.DeleteCollectionOutput, err error)
		GetList(ctx context.Context, in model.CollectionListInput) (out *model.CollectionListOutput, err error)
	}
)

var (
	localCollection ICollection
)

func Collection() ICollection {
	if localCollection == nil {
		panic("implement not found for interface ICollection, forgot register?")
	}
	return localCollection
}

func RegisterCollection(i ICollection) {
	localCollection = i
}
