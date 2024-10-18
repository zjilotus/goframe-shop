package controller

import (
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
	"golang.org/x/net/context"
)

var Goods = cGoods{}

type cGoods struct{}

func (a *cGoods) Create(ctx context.Context, req *backend.GoodsReq) (res *backend.GoodsRes, err error) {
	data := model.GoodsCreateInput{}
	err = gconv.Scan(req, &data)
	out, err := service.Goods().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.GoodsRes{Id: out.Id}, nil
}

func (a *cGoods) Delete(ctx context.Context, req *backend.GoodsDeleteReq) (res *backend.GoodsDeleteRes, err error) {
	err = service.Goods().Delete(ctx, req.Id)
	return
}

func (a *cGoods) Update(ctx context.Context, req *backend.GoodsUpdateReq) (res *backend.GoodsUpdateRes, err error) {
	data := model.GoodsUpdateInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	err = service.Goods().Update(ctx, data)
	return &backend.GoodsUpdateRes{Id: req.Id}, nil
}

func (a *cGoods) List(ctx context.Context, req *backend.GoodsGetListCommonReq) (res *backend.GoodsGetListCommonRes, err error) {
	getListRes, err := service.Goods().GetList(ctx, model.GoodsGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.GoodsGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
