package coach

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
	service.RegisterCoach(New())
}

func New() *sCoach {
	return &sCoach{}
}

type sCoach struct{}

// Create 创建陪玩师表
func (s *sCoach) Create(ctx context.Context, in *model.CoachCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayCoach.Ctx(ctx).Data(g.Map{
		dao.PlayCoach.Columns().Id:        id,
		dao.PlayCoach.Columns().MemberId: in.MemberID,
		dao.PlayCoach.Columns().CoachLevelId: in.CoachLevelID,
		dao.PlayCoach.Columns().ShopId: in.ShopID,
		dao.PlayCoach.Columns().RealName: in.RealName,
		dao.PlayCoach.Columns().Intro: in.Intro,
		dao.PlayCoach.Columns().CoverImage: in.CoverImage,
		dao.PlayCoach.Columns().TotalOrders: in.TotalOrders,
		dao.PlayCoach.Columns().TotalScore: in.TotalScore,
		dao.PlayCoach.Columns().ScoreNum: in.ScoreNum,
		dao.PlayCoach.Columns().IncomeTotal: in.IncomeTotal,
		dao.PlayCoach.Columns().IncomeBalance: in.IncomeBalance,
		dao.PlayCoach.Columns().IsOnline: in.IsOnline,
		dao.PlayCoach.Columns().Sort: in.Sort,
		dao.PlayCoach.Columns().Status: in.Status,
		dao.PlayCoach.Columns().CreatedAt: gtime.Now(),
		dao.PlayCoach.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新陪玩师表
func (s *sCoach) Update(ctx context.Context, in *model.CoachUpdateInput) error {
	data := g.Map{
		dao.PlayCoach.Columns().MemberId: in.MemberID,
		dao.PlayCoach.Columns().CoachLevelId: in.CoachLevelID,
		dao.PlayCoach.Columns().ShopId: in.ShopID,
		dao.PlayCoach.Columns().RealName: in.RealName,
		dao.PlayCoach.Columns().Intro: in.Intro,
		dao.PlayCoach.Columns().CoverImage: in.CoverImage,
		dao.PlayCoach.Columns().TotalOrders: in.TotalOrders,
		dao.PlayCoach.Columns().TotalScore: in.TotalScore,
		dao.PlayCoach.Columns().ScoreNum: in.ScoreNum,
		dao.PlayCoach.Columns().IncomeTotal: in.IncomeTotal,
		dao.PlayCoach.Columns().IncomeBalance: in.IncomeBalance,
		dao.PlayCoach.Columns().IsOnline: in.IsOnline,
		dao.PlayCoach.Columns().Sort: in.Sort,
		dao.PlayCoach.Columns().Status: in.Status,
		dao.PlayCoach.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayCoach.Ctx(ctx).Where(dao.PlayCoach.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除陪玩师表
func (s *sCoach) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayCoach.Ctx(ctx).Where(dao.PlayCoach.Columns().Id, id).Data(g.Map{
		dao.PlayCoach.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取陪玩师表详情
func (s *sCoach) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.CoachDetailOutput, err error) {
	out = &model.CoachDetailOutput{}
	err = dao.PlayCoach.Ctx(ctx).Where(dao.PlayCoach.Columns().Id, id).Where(dao.PlayCoach.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	// 查询陪玩师等级ID关联显示
	if out.CoachLevelID != 0 {
		val, _ := g.DB().Ctx(ctx).Model("play_coach_level").Where("id", out.CoachLevelID).Where("deleted_at", nil).Value("title")
		out.CoachLevelTitle = val.String()
	}
	// 查询所属店铺ID（0表示无店铺）关联显示
	if out.ShopID != 0 {
		val, _ := g.DB().Ctx(ctx).Model("play_shop").Where("id", out.ShopID).Where("deleted_at", nil).Value("title")
		out.ShopTitle = val.String()
	}
	return
}

// List 获取陪玩师表列表
func (s *sCoach) List(ctx context.Context, in *model.CoachListInput) (list []*model.CoachListOutput, total int, err error) {
	m := dao.PlayCoach.Ctx(ctx).Where(dao.PlayCoach.Columns().DeletedAt, nil)
	if in.IsOnline > 0 {
		m = m.Where(dao.PlayCoach.Columns().IsOnline, in.IsOnline)
	}
	if in.Status > 0 {
		m = m.Where(dao.PlayCoach.Columns().Status, in.Status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayCoach.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	// 填充关联显示字段
	for _, item := range list {
		if item.CoachLevelID != 0 {
			val, _ := g.DB().Ctx(ctx).Model("play_coach_level").Where("id", item.CoachLevelID).Where("deleted_at", nil).Value("title")
			item.CoachLevelTitle = val.String()
		}
		if item.ShopID != 0 {
			val, _ := g.DB().Ctx(ctx).Model("play_shop").Where("id", item.ShopID).Where("deleted_at", nil).Value("title")
			item.ShopTitle = val.String()
		}
	}
	return
}

