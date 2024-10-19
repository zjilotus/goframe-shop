package controller

import (
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/api/frontend"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
	"golang.org/x/net/context"
)

var User = cUser{}

type cUser struct{}

func (c *cUser) Register(ctx context.Context, req *frontend.RegisterReq) (res *frontend.RegisterRes, err error) {
	data := model.RegisterInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.User().Register(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.RegisterRes{Id: out.Id}, nil
}

func (c *cUser) Info(ctx context.Context, req *frontend.UserInfoReq) (res *frontend.UserInfoRes, err error) {
	res = &frontend.UserInfoRes{}
	res.Name = gconv.String(ctx.Value(consts.CtxUserName))
	res.Avatar = gconv.String(ctx.Value(consts.CtxUserAvatar))
	res.Id = gconv.Uint(ctx.Value(consts.CtxUserId))
	res.Sex = gconv.Uint8(ctx.Value(consts.CtxUserSex))
	res.Sign = gconv.String(ctx.Value(consts.CtxUserSign))
	res.Status = gconv.Uint8(ctx.Value(consts.CtxUserStatus))
	return res, nil
}
