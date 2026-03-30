package recharge_order

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
	service.RegisterRechargeOrder(New())
}

func New() *sRechargeOrder {
	return &sRechargeOrder{}
}

type sRechargeOrder struct{}

// Create 创建充值订单表
func (s *sRechargeOrder) Create(ctx context.Context, in *model.RechargeOrderCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayRechargeOrder.Ctx(ctx).Data(g.Map{
		dao.PlayRechargeOrder.Columns().Id:        id,
		dao.PlayRechargeOrder.Columns().OrderNo: in.OrderNo,
		dao.PlayRechargeOrder.Columns().MemberId: in.MemberID,
		dao.PlayRechargeOrder.Columns().RechargePlanId: in.RechargePlanID,
		dao.PlayRechargeOrder.Columns().Amount: in.Amount,
		dao.PlayRechargeOrder.Columns().GiftAmount: in.GiftAmount,
		dao.PlayRechargeOrder.Columns().PayType: in.PayType,
		dao.PlayRechargeOrder.Columns().TradeNo: in.TradeNo,
		dao.PlayRechargeOrder.Columns().PayStatus: in.PayStatus,
		dao.PlayRechargeOrder.Columns().PayAt: in.PayAt,
		dao.PlayRechargeOrder.Columns().CreatedAt: gtime.Now(),
		dao.PlayRechargeOrder.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新充值订单表
func (s *sRechargeOrder) Update(ctx context.Context, in *model.RechargeOrderUpdateInput) error {
	data := g.Map{
		dao.PlayRechargeOrder.Columns().OrderNo: in.OrderNo,
		dao.PlayRechargeOrder.Columns().MemberId: in.MemberID,
		dao.PlayRechargeOrder.Columns().RechargePlanId: in.RechargePlanID,
		dao.PlayRechargeOrder.Columns().Amount: in.Amount,
		dao.PlayRechargeOrder.Columns().GiftAmount: in.GiftAmount,
		dao.PlayRechargeOrder.Columns().PayType: in.PayType,
		dao.PlayRechargeOrder.Columns().TradeNo: in.TradeNo,
		dao.PlayRechargeOrder.Columns().PayStatus: in.PayStatus,
		dao.PlayRechargeOrder.Columns().PayAt: in.PayAt,
		dao.PlayRechargeOrder.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayRechargeOrder.Ctx(ctx).Where(dao.PlayRechargeOrder.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除充值订单表
func (s *sRechargeOrder) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayRechargeOrder.Ctx(ctx).Where(dao.PlayRechargeOrder.Columns().Id, id).Data(g.Map{
		dao.PlayRechargeOrder.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取充值订单表详情
func (s *sRechargeOrder) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.RechargeOrderDetailOutput, err error) {
	out = &model.RechargeOrderDetailOutput{}
	err = dao.PlayRechargeOrder.Ctx(ctx).Where(dao.PlayRechargeOrder.Columns().Id, id).Where(dao.PlayRechargeOrder.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	// 查询充值方案ID关联显示
	if out.RechargePlanID != 0 {
		val, err := g.DB().Ctx(ctx).Model("play_recharge_plan").Where("id", out.RechargePlanID).Where("deleted_at", nil).Value("title")
		if err == nil {
			out.RechargePlanTitle = val.String()
		}
	}
	return
}

// List 获取充值订单表列表
func (s *sRechargeOrder) List(ctx context.Context, in *model.RechargeOrderListInput) (list []*model.RechargeOrderListOutput, total int, err error) {
	m := dao.PlayRechargeOrder.Ctx(ctx).Where(dao.PlayRechargeOrder.Columns().DeletedAt, nil)
	if in.PayType > 0 {
		m = m.Where(dao.PlayRechargeOrder.Columns().PayType, in.PayType)
	}
	if in.PayStatus > 0 {
		m = m.Where(dao.PlayRechargeOrder.Columns().PayStatus, in.PayStatus)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayRechargeOrder.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	// 填充关联显示字段
	for _, item := range list {
		if item.RechargePlanID != 0 {
			val, err := g.DB().Ctx(ctx).Model("play_recharge_plan").Where("id", item.RechargePlanID).Where("deleted_at", nil).Value("title")
			if err == nil {
				item.RechargePlanTitle = val.String()
			}
		}
	}
	return
}

