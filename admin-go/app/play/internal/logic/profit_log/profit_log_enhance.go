package profit_log

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// SettleOrder 订单完成时利润结算
func (s *sProfitLog) SettleOrder(ctx context.Context, in *model.SettleOrderInput) error {
	// 1. 查询订单信息
	order, err := dao.PlayOrder.Ctx(ctx).Where(dao.PlayOrder.Columns().Id, in.OrderID).One()
	if err != nil {
		return err
	}
	if order.IsEmpty() {
		return gerror.New("订单不存在")
	}
	if order[dao.PlayOrder.Columns().OrderStatus].Int() != 3 {
		return gerror.New("订单未完成，不能结算")
	}

	payAmount := order[dao.PlayOrder.Columns().PayAmount].Int64()
	coachID := order[dao.PlayOrder.Columns().CoachId].Int64()
	shopID := order[dao.PlayOrder.Columns().ShopId].Int64()
	orderNo := order[dao.PlayOrder.Columns().OrderNo].String()

	// 2. 查询店铺抽成比例（commission_rate 即平台抽成比例）
	shopVal, err := dao.PlayShop.Ctx(ctx).Where(dao.PlayShop.Columns().Id, shopID).One()
	if err != nil {
		return err
	}
	if shopVal.IsEmpty() {
		return gerror.New("店铺不存在")
	}
	commissionRate := shopVal[dao.PlayShop.Columns().CommissionRate].Int()

	// 3. 计算分成
	platformAmount := payAmount * int64(commissionRate) / 100
	coachAmount := payAmount - platformAmount

	// 4. 写入 profit_log
	profitID := snowflake.Generate()
	_, err = dao.PlayProfitLog.Ctx(ctx).Data(g.Map{
		dao.PlayProfitLog.Columns().Id:             profitID,
		dao.PlayProfitLog.Columns().OrderId:        in.OrderID,
		dao.PlayProfitLog.Columns().OrderNo:        orderNo,
		dao.PlayProfitLog.Columns().PayAmount:      payAmount,
		dao.PlayProfitLog.Columns().CoachId:        coachID,
		dao.PlayProfitLog.Columns().ShopId:         shopID,
		dao.PlayProfitLog.Columns().PlatformRate:   commissionRate,
		dao.PlayProfitLog.Columns().PlatformAmount: platformAmount,
		dao.PlayProfitLog.Columns().ShopRate:       100 - commissionRate,
		dao.PlayProfitLog.Columns().ShopAmount:     coachAmount,
		dao.PlayProfitLog.Columns().CoachAmount:    coachAmount,
		dao.PlayProfitLog.Columns().SettleStatus:   1,
		dao.PlayProfitLog.Columns().SettleAt:       gtime.Now(),
		dao.PlayProfitLog.Columns().CreatedAt:      gtime.Now(),
		dao.PlayProfitLog.Columns().UpdatedAt:      gtime.Now(),
	}).Insert()
	if err != nil {
		return err
	}

	// 5. 更新陪玩师收入
	_, err = dao.PlayCoach.Ctx(ctx).Where(dao.PlayCoach.Columns().Id, coachID).Data(g.Map{
		dao.PlayCoach.Columns().IncomeTotal:   gdb.Raw(fmt.Sprintf("`income_total` + %d", coachAmount)),
		dao.PlayCoach.Columns().IncomeBalance: gdb.Raw(fmt.Sprintf("`income_balance` + %d", coachAmount)),
		dao.PlayCoach.Columns().UpdatedAt:     gtime.Now(),
	}).Update()
	return err
}
