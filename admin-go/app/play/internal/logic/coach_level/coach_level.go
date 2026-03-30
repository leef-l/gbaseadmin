package coach_level

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
	service.RegisterCoachLevel(New())
}

func New() *sCoachLevel {
	return &sCoachLevel{}
}

type sCoachLevel struct{}

// Create 创建陪玩师等级表
func (s *sCoachLevel) Create(ctx context.Context, in *model.CoachLevelCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayCoachLevel.Ctx(ctx).Data(g.Map{
		dao.PlayCoachLevel.Columns().Id:        id,
		dao.PlayCoachLevel.Columns().Title: in.Title,
		dao.PlayCoachLevel.Columns().Level: in.Level,
		dao.PlayCoachLevel.Columns().Icon: in.Icon,
		dao.PlayCoachLevel.Columns().MinOrders: in.MinOrders,
		dao.PlayCoachLevel.Columns().MinScore: in.MinScore,
		dao.PlayCoachLevel.Columns().CommissionRate: in.CommissionRate,
		dao.PlayCoachLevel.Columns().Sort: in.Sort,
		dao.PlayCoachLevel.Columns().Status: in.Status,
		dao.PlayCoachLevel.Columns().CreatedAt: gtime.Now(),
		dao.PlayCoachLevel.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新陪玩师等级表
func (s *sCoachLevel) Update(ctx context.Context, in *model.CoachLevelUpdateInput) error {
	data := g.Map{
		dao.PlayCoachLevel.Columns().Title: in.Title,
		dao.PlayCoachLevel.Columns().Level: in.Level,
		dao.PlayCoachLevel.Columns().Icon: in.Icon,
		dao.PlayCoachLevel.Columns().MinOrders: in.MinOrders,
		dao.PlayCoachLevel.Columns().MinScore: in.MinScore,
		dao.PlayCoachLevel.Columns().CommissionRate: in.CommissionRate,
		dao.PlayCoachLevel.Columns().Sort: in.Sort,
		dao.PlayCoachLevel.Columns().Status: in.Status,
		dao.PlayCoachLevel.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayCoachLevel.Ctx(ctx).Where(dao.PlayCoachLevel.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除陪玩师等级表
func (s *sCoachLevel) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayCoachLevel.Ctx(ctx).Where(dao.PlayCoachLevel.Columns().Id, id).Data(g.Map{
		dao.PlayCoachLevel.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取陪玩师等级表详情
func (s *sCoachLevel) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.CoachLevelDetailOutput, err error) {
	out = &model.CoachLevelDetailOutput{}
	err = dao.PlayCoachLevel.Ctx(ctx).Where(dao.PlayCoachLevel.Columns().Id, id).Where(dao.PlayCoachLevel.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	return
}

// List 获取陪玩师等级表列表
func (s *sCoachLevel) List(ctx context.Context, in *model.CoachLevelListInput) (list []*model.CoachLevelListOutput, total int, err error) {
	m := dao.PlayCoachLevel.Ctx(ctx).Where(dao.PlayCoachLevel.Columns().DeletedAt, nil)
	if in.Level > 0 {
		m = m.Where(dao.PlayCoachLevel.Columns().Level, in.Level)
	}
	if in.Status > 0 {
		m = m.Where(dao.PlayCoachLevel.Columns().Status, in.Status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayCoachLevel.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	return
}

