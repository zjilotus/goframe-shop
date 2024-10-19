package user

import (
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/do"
	"goframe-shop/internal/service"
	"goframe-shop/utility"
	"golang.org/x/net/context"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) Register(ctx context.Context, in model.RegisterInput) (out model.RegisterOutput, err error) {
	//处理加密盐和密码的逻辑
	UserSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt
	//插入数据返回id
	lastInsertID, err := dao.UserInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RegisterOutput{Id: uint(lastInsertID)}, err
}

func (s *sUser) UpdatePassword(ctx context.Context, in model.UpdatePasswordInput) (out model.UpdatePasswordOutput, err error) {
	userInfo := do.UserInfo{}
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	err = dao.UserInfo.Ctx(ctx).WherePri(userId).Scan(&userInfo)
	if err != nil {
		return out, err
	}
	if gconv.String(userInfo.SecretAnswer) == in.SecretAnswer {
		in.UserSalt = grand.S(10)
		in.Password = utility.EncryptPassword(in.Password, in.UserSalt)
		_, err := dao.UserInfo.Ctx(ctx).Data(in).WherePri(userId).Update()
		if err != nil {
			return out, err
		}
	} else {
		g.Dump("userInfo.SecretAnswer:", userInfo.SecretAnswer)
		g.Dump("in.SecretAnswer:", in.SecretAnswer)
		return out, errors.New(consts.ErrSecretAnswerMsg)
	}
	return model.UpdatePasswordOutput{Id: userId}, nil
}
