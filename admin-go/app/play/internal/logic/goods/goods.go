package goods

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
	"gbaseadmin/utility/snowflake"
)

func init() {
	service.RegisterGoods(New())
}

func New() *sGoods {
	return &sGoods{}
}

type sGoods struct{}

// Create 创建å•†å“è¡¨
func (s *sGoods) Create(ctx context.Context, in *model.GoodsCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayGoods.Ctx(ctx).Data(g.Map{
		dao.PlayGoods.Columns().Id:        id,
		dao.PlayGoods.Columns().CategoryId: in.CategoryID,
		dao.PlayGoods.Columns().CoachId: in.CoachID,
		dao.PlayGoods.Columns().Title: in.Title,
		dao.PlayGoods.Columns().CoverImage: in.CoverImage,
		dao.PlayGoods.Columns().DescContent: in.DescContent,
		dao.PlayGoods.Columns().Price: in.Price,
		dao.PlayGoods.Columns().Unit: in.Unit,
		dao.PlayGoods.Columns().SalesNum: in.SalesNum,
		dao.PlayGoods.Columns().Sort: in.Sort,
		dao.PlayGoods.Columns().Status: in.Status,
		dao.PlayGoods.Columns().CreatedAt: gtime.Now(),
		dao.PlayGoods.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新å•†å“è¡¨
func (s *sGoods) Update(ctx context.Context, in *model.GoodsUpdateInput) error {
	data := g.Map{
		dao.PlayGoods.Columns().CategoryId: in.CategoryID,
		dao.PlayGoods.Columns().CoachId: in.CoachID,
		dao.PlayGoods.Columns().Title: in.Title,
		dao.PlayGoods.Columns().CoverImage: in.CoverImage,
		dao.PlayGoods.Columns().DescContent: in.DescContent,
		dao.PlayGoods.Columns().Price: in.Price,
		dao.PlayGoods.Columns().Unit: in.Unit,
		dao.PlayGoods.Columns().SalesNum: in.SalesNum,
		dao.PlayGoods.Columns().Sort: in.Sort,
		dao.PlayGoods.Columns().Status: in.Status,
		dao.PlayGoods.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayGoods.Ctx(ctx).Where(dao.PlayGoods.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除å•†å“è¡¨
func (s *sGoods) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayGoods.Ctx(ctx).Where(dao.PlayGoods.Columns().Id, id).Data(g.Map{
		dao.PlayGoods.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取å•†å“è¡¨详情
func (s *sGoods) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.GoodsDetailOutput, err error) {
	out = &model.GoodsDetailOutput{}
	err = dao.PlayGoods.Ctx(ctx).Where(dao.PlayGoods.Columns().Id, id).Where(dao.PlayGoods.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	// 查询åˆ†ç±»ID关联显示
	if out.CategoryID != 0 {
		val, _ := g.DB().Ctx(ctx).Model("play_category").Where("id", out.CategoryID).Where("deleted_at", nil).Value("title")
		out.CategoryTitle = val.String()
	}
	return
}

// List 获取å•†å“è¡¨列表
func (s *sGoods) List(ctx context.Context, in *model.GoodsListInput) (list []*model.GoodsListOutput, total int, err error) {
	m := dao.PlayGoods.Ctx(ctx).Where(dao.PlayGoods.Columns().DeletedAt, nil)
	if in.Status > 0 {
		m = m.Where(dao.PlayGoods.Columns().Status, in.Status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayGoods.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	// 填充关联显示字段
	for _, item := range list {
		if item.CategoryID != 0 {
			val, _ := g.DB().Ctx(ctx).Model("play_category").Where("id", item.CategoryID).Where("deleted_at", nil).Value("title")
			item.CategoryTitle = val.String()
		}
	}
	return
}

