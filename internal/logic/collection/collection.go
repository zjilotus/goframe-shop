package collection

import (
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
	"golang.org/x/net/context"
)

type sCollection struct{}

func init() {
	service.RegisterCollection(New())
}

func New() *sCollection {
	return &sCollection{}
}

func (s *sCollection) AddCollection(ctx context.Context, in model.AddCollectionInput) (out *model.AddCollectionOutput, err error) {
	id, err := dao.CollectionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.AddCollectionOutput{Id: uint(id)}, nil
}

// 兼容处理：优先根据收藏id删除，收藏id为0；再根据对象id和type删除
func (s *sCollection) DeleteCollection(ctx context.Context, in model.DeleteCollectionInput) (out *model.DeleteCollectionOutput, err error) {
	if in.Id != 0 {
		_, err := dao.CollectionInfo.Ctx(ctx).WherePri(in.Id).Delete()
		if err != nil {
			return nil, err
		}
	} else {
		_, err := dao.CollectionInfo.Ctx(ctx).OmitEmpty(). // 过滤空值
									Where(in).Delete()
		if err != nil {
			return nil, err
		}
	}
	return &model.DeleteCollectionOutput{Id: in.Id}, nil
}

func (s *sCollection) GetList(ctx context.Context, in model.CollectionListInput) (out *model.CollectionListOutput, err error) {
	m := dao.CollectionInfo.Ctx(ctx)
	out = &model.CollectionListOutput{
		Page: in.Page,
		Size: in.Size,
		List: []model.CollectionListOutputItem{}, // 数据为空时放回 [] 而不是 null
	}
	// 翻页查询
	listModel := m.Page(in.Page, in.Size)
	// 条件查询
	if in.Type != 0 {
		listModel = listModel.Where(dao.CollectionInfo.Columns().Type, in.Type)
	}

	out.Total, err = listModel.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}

	if in.Type == consts.CollectionTypeGoods {
		err = listModel.With(model.GoodsItem{}).Scan(&out.List)
	} else if in.Type == consts.CollectionTypeArticle {
		err = listModel.With(model.ArticleItem{}).Scan(&out.List)
	} else {
		err = listModel.WithAll().Scan(&out.List)
	}
	if err != nil {
		return nil, err
	}
	return
}
