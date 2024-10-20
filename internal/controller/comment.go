package controller

import (
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/frontend"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
	"golang.org/x/net/context"
)

var Comment = cComment{}

type cComment struct{}

func (a *cComment) Add(ctx context.Context, req *frontend.AddCommentReq) (res *frontend.AddCommentRes, err error) {
	data := model.AddCommentInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	data.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	out, err := service.Comment().AddComment(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.AddCommentRes{Id: out.Id}, nil
}

func (a *cComment) Delete(ctx context.Context, req *frontend.DeleteCommentReq) (res *frontend.DeleteCommentRes, err error) {
	data := model.DeleteCommentInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	collection, err := service.Comment().DeleteComment(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.DeleteCommentRes{Id: collection.Id}, nil
}

func (a *cComment) List(ctx context.Context, req *frontend.ListCommentReq) (res *frontend.ListCommentRes, err error) {
	getListRes, err := service.Comment().GetList(ctx, model.CommentListInput{
		Page: req.Page,
		Size: req.Size,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.ListCommentRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
