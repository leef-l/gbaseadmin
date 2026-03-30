package profit_log

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
	service.RegisterProfitLog(New())
}

func New() *sProfitLog {
	return &sProfitLog{}
}

type sProfitLog struct{}

// Create 创建利润分成流水表
func (s *sProfitLog) Create(ctx context.Context, in *model.ProfitLogCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayProfitLog.Ctx(ctx).Data(g.Map{
		dao.PlayProfitLog.Columns().Id:        id,
		dao.PlayProfitLog.Columns().OrderId: in.OrderID,
		dao.PlayProfitLog.Columns().OrderNo: in.OrderNo,
		dao.PlayProfitLog.Columns().PayAmount: in.PayAmount,
		dao.PlayProfitLog.Columns().CoachId: in.CoachID,
		dao.PlayProfitLog.Columns().ShopId: in.ShopID,
		dao.PlayProfitLog.Columns().PlatformRate: in.PlatformRate,
		dao.PlayProfitLog.Columns().PlatformAmount: in.PlatformAmount,
		dao.PlayProfitLog.Columns().ShopRate: in.ShopRate,
		dao.PlayProfitLog.Columns().ShopAmount: in.ShopAmount,
		dao.PlayProfitLog.Columns().CoachAmount: in.CoachAmount,
		dao.PlayProfitLog.Columns().SettleStatus: in.SettleStatus,
		dao.PlayProfitLog.Columns().SettleAt: in.SettleAt,
		dao.PlayProfitLog.Columns().CreatedAt: gtime.Now(),
		dao.PlayProfitLog.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新利润分成流水表
func (s *sProfitLog) Update(ctx context.Context, in *model.ProfitLogUpdateInput) error {
	data := g.Map{
		dao.PlayProfitLog.Columns().OrderId: in.OrderID,
		dao.PlayProfitLog.Columns().OrderNo: in.OrderNo,
		dao.PlayProfitLog.Columns().PayAmount: in.PayAmount,
		dao.PlayProfitLog.Columns().CoachId: in.CoachID,
		dao.PlayProfitLog.Columns().ShopId: in.ShopID,
		dao.PlayProfitLog.Columns().PlatformRate: in.PlatformRate,
		dao.PlayProfitLog.Columns().PlatformAmount: in.PlatformAmount,
		dao.PlayProfitLog.Columns().ShopRate: in.ShopRate,
		dao.PlayProfitLog.Columns().ShopAmount: in.ShopAmount,
		dao.PlayProfitLog.Columns().CoachAmount: in.CoachAmount,
		dao.PlayProfitLog.Columns().SettleStatus: in.SettleStatus,
		dao.PlayProfitLog.Columns().SettleAt: in.SettleAt,
		dao.PlayProfitLog.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayProfitLog.Ctx(ctx).Where(dao.PlayProfitLog.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除利润分成流水表
func (s *sProfitLog) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayProfitLog.Ctx(ctx).Where(dao.PlayProfitLog.Columns().Id, id).Data(g.Map{
		dao.PlayProfitLog.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取利润分成流水表详情
func (s *sProfitLog) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.ProfitLogDetailOutput, err error) {
	out = &model.ProfitLogDetailOutput{}
	err = dao.PlayProfitLog.Ctx(ctx).Where(dao.PlayProfitLog.Columns().Id, id).Where(dao.PlayProfitLog.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	// 查询店铺ID（0表示无店铺）关联显示
	if out.ShopID != 0 {
		val, err := g.DB().Ctx(ctx).Model("play_shop").Where("id", out.ShopID).Where("deleted_at", nil).Value("title")
		if err == nil {
			out.ShopTitle = val.String()
		}
	}
	return
}

// List 获取利润分成流水表列表
func (s *sProfitLog) List(ctx context.Context, in *model.ProfitLogListInput) (list []*model.ProfitLogListOutput, total int, err error) {
	m := dao.PlayProfitLog.Ctx(ctx).Where(dao.PlayProfitLog.Columns().DeletedAt, nil)
	if in.SettleStatus > 0 {
		m = m.Where(dao.PlayProfitLog.Columns().SettleStatus, in.SettleStatus)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayProfitLog.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	// 填充关联显示字段
	for _, item := range list {
		if item.ShopID != 0 {
			val, err := g.DB().Ctx(ctx).Model("play_shop").Where("id", item.ShopID).Where("deleted_at", nil).Value("title")
			if err == nil {
				item.ShopTitle = val.String()
			}
		}
	}
	return
}

