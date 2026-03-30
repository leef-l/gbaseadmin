package shop

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
	service.RegisterShop(New())
}

func New() *sShop {
	return &sShop{}
}

type sShop struct{}

// Create 创建店铺表
func (s *sShop) Create(ctx context.Context, in *model.ShopCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayShop.Ctx(ctx).Data(g.Map{
		dao.PlayShop.Columns().Id:        id,
		dao.PlayShop.Columns().Title: in.Title,
		dao.PlayShop.Columns().LogoImage: in.LogoImage,
		dao.PlayShop.Columns().CoverImage: in.CoverImage,
		dao.PlayShop.Columns().ContactName: in.ContactName,
		dao.PlayShop.Columns().ContactPhone: in.ContactPhone,
		dao.PlayShop.Columns().Intro: in.Intro,
		dao.PlayShop.Columns().CommissionRate: in.CommissionRate,
		dao.PlayShop.Columns().CoachNum: in.CoachNum,
		dao.PlayShop.Columns().Sort: in.Sort,
		dao.PlayShop.Columns().Status: in.Status,
		dao.PlayShop.Columns().CreatedAt: gtime.Now(),
		dao.PlayShop.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新店铺表
func (s *sShop) Update(ctx context.Context, in *model.ShopUpdateInput) error {
	data := g.Map{
		dao.PlayShop.Columns().Title: in.Title,
		dao.PlayShop.Columns().LogoImage: in.LogoImage,
		dao.PlayShop.Columns().CoverImage: in.CoverImage,
		dao.PlayShop.Columns().ContactName: in.ContactName,
		dao.PlayShop.Columns().ContactPhone: in.ContactPhone,
		dao.PlayShop.Columns().Intro: in.Intro,
		dao.PlayShop.Columns().CommissionRate: in.CommissionRate,
		dao.PlayShop.Columns().CoachNum: in.CoachNum,
		dao.PlayShop.Columns().Sort: in.Sort,
		dao.PlayShop.Columns().Status: in.Status,
		dao.PlayShop.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayShop.Ctx(ctx).Where(dao.PlayShop.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除店铺表
func (s *sShop) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayShop.Ctx(ctx).Where(dao.PlayShop.Columns().Id, id).Data(g.Map{
		dao.PlayShop.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取店铺表详情
func (s *sShop) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.ShopDetailOutput, err error) {
	out = &model.ShopDetailOutput{}
	err = dao.PlayShop.Ctx(ctx).Where(dao.PlayShop.Columns().Id, id).Where(dao.PlayShop.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	return
}

// List 获取店铺表列表
func (s *sShop) List(ctx context.Context, in *model.ShopListInput) (list []*model.ShopListOutput, total int, err error) {
	m := dao.PlayShop.Ctx(ctx).Where(dao.PlayShop.Columns().DeletedAt, nil)
	if in.Status > 0 {
		m = m.Where(dao.PlayShop.Columns().Status, in.Status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayShop.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	return
}

