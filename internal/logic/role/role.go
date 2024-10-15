package role

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/entity"
	"goframe-shop/internal/service"
)

type sRole struct{}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

func (s *sRole) Create(ctx context.Context, in model.RoleCreateInput) (out model.RoleCreateOutput, err error) {
	lastInsertID, err := dao.RoleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RoleCreateOutput{RoleId: uint(int(lastInsertID))}, err
}

func (s *sRole) Delete(ctx context.Context, id uint) error {
	_, err := dao.RoleInfo.Ctx(ctx).Where(g.Map{
		dao.RoleInfo.Columns().Id: id,
	}).Unscoped().Delete()
	return err
}

func (s *sRole) Update(ctx context.Context, in model.RoleUpdateInput) error {
	_, err := dao.RoleInfo.Ctx(ctx).Data(in).FieldsEx(dao.RoleInfo.Columns().Id, in.Id).Update()
	return err
}

func (s *sRole) GetList(ctx context.Context, in model.RoleGetListInput) (out *model.RoleGetListOutput, err error) {
	var (
		m = dao.RoleInfo.Ctx(ctx)
	)
	out = &model.RoleGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	var list []*entity.RoleInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}

	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}

	// 不指定item的键名用：Scan
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

func (s *sRole) AddPermission(ctx context.Context, in model.RoleAddPermissionInput) (out model.RoleAddPermissionOutput, err error) {
	id, err := dao.RolePermissionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return model.RoleAddPermissionOutput{}, err
	}
	return model.RoleAddPermissionOutput{
		Id: uint(id),
	}, err
}

func (s *sRole) DeletePermission(ctx context.Context, in model.RoleDeletePermissionInput) error {
	_, err := dao.RolePermissionInfo.Ctx(ctx).Where(g.Map{
		dao.RolePermissionInfo.Columns().RoleId:       in.RoleId,
		dao.RolePermissionInfo.Columns().PermissionId: in.PermissionId,
	}).Delete()
	if err != nil {
		return err
	}
	return err
}
