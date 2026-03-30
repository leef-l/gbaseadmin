package recharge_plan

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
	service.RegisterRechargePlan(New())
}

func New() *sRechargePlan {
	return &sRechargePlan{}
}

type sRechargePlan struct{}

// Create 创建充值方案表
func (s *sRechargePlan) Create(ctx context.Context, in *model.RechargePlanCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayRechargePlan.Ctx(ctx).Data(g.Map{
		dao.PlayRechargePlan.Columns().Id:        id,
		dao.PlayRechargePlan.Columns().Title: in.Title,
		dao.PlayRechargePlan.Columns().Amount: in.Amount,
		dao.PlayRechargePlan.Columns().GiftAmount: in.GiftAmount,
		dao.PlayRechargePlan.Columns().CoverImage: in.CoverImage,
		dao.PlayRechargePlan.Columns().Sort: in.Sort,
		dao.PlayRechargePlan.Columns().Status: in.Status,
		dao.PlayRechargePlan.Columns().CreatedAt: gtime.Now(),
		dao.PlayRechargePlan.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新充值方案表
func (s *sRechargePlan) Update(ctx context.Context, in *model.RechargePlanUpdateInput) error {
	data := g.Map{
		dao.PlayRechargePlan.Columns().Title: in.Title,
		dao.PlayRechargePlan.Columns().Amount: in.Amount,
		dao.PlayRechargePlan.Columns().GiftAmount: in.GiftAmount,
		dao.PlayRechargePlan.Columns().CoverImage: in.CoverImage,
		dao.PlayRechargePlan.Columns().Sort: in.Sort,
		dao.PlayRechargePlan.Columns().Status: in.Status,
		dao.PlayRechargePlan.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayRechargePlan.Ctx(ctx).Where(dao.PlayRechargePlan.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除充值方案表
func (s *sRechargePlan) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayRechargePlan.Ctx(ctx).Where(dao.PlayRechargePlan.Columns().Id, id).Data(g.Map{
		dao.PlayRechargePlan.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取充值方案表详情
func (s *sRechargePlan) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.RechargePlanDetailOutput, err error) {
	out = &model.RechargePlanDetailOutput{}
	err = dao.PlayRechargePlan.Ctx(ctx).Where(dao.PlayRechargePlan.Columns().Id, id).Where(dao.PlayRechargePlan.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	return
}

// List 获取充值方案表列表
func (s *sRechargePlan) List(ctx context.Context, in *model.RechargePlanListInput) (list []*model.RechargePlanListOutput, total int, err error) {
	m := dao.PlayRechargePlan.Ctx(ctx).Where(dao.PlayRechargePlan.Columns().DeletedAt, nil)
	if in.Status > 0 {
		m = m.Where(dao.PlayRechargePlan.Columns().Status, in.Status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayRechargePlan.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	return
}

