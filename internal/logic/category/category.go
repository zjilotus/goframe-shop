package category

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/entity"
	"goframe-shop/internal/service"
	"golang.org/x/net/context"
)

type sCategory struct {
}

func init() {
	service.RegisterCategory(New())
}

func New() *sCategory {
	return &sCategory{}
}

func (s *sCategory) Create(ctx context.Context, in model.CategoryCreateInput) (out model.CategoryCreateOutput, err error) {
	lastInsertId, err := dao.CategoryInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.CategoryCreateOutput{CategoryId: int(lastInsertId)}, nil
}

func (s *sCategory) Delete(ctx context.Context, id uint) (err error) {
	_, err = dao.CategoryInfo.Ctx(ctx).Where(g.Map{
		dao.CategoryInfo.Columns().Id: id,
	}).Delete()
	if err != nil {
		return err
	}
	return
}

func (s *sCategory) Update(ctx context.Context, in model.CategoryUpdateInput) error {
	_, err := dao.CategoryInfo.Ctx(ctx).
		Data(in).
		FieldsEx().
		Where(dao.CategoryInfo.Columns().Id, in.Id).
		Update()
	if err != nil {
		return err
	}
	return err
}

// GetList 查询分类列表
func (s *sCategory) GetList(ctx context.Context, in model.CategoryGetListInput) (out *model.CategoryGetListOutput, err error) {
	var (
		m = dao.CategoryInfo.Ctx(ctx)
	)
	out = &model.CategoryGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	listModel = listModel.OrderDesc(dao.CategoryInfo.Columns().Sort)

	// 执行查询
	var list []*entity.CategoryInfo
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
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
