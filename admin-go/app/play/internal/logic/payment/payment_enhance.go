package payment

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
	"gbaseadmin/utility/snowflake"
)

// BalancePay 余额支付
func (s *sPayment) BalancePay(ctx context.Context, in *model.BalancePayInput) error {
	// 1. 扣减余额（通过 AddLog 统一入口）
	err := service.BalanceLogEnhance().AddLog(ctx, &model.AddBalanceLogInput{
		MemberID:     in.MemberID,
		BizType:      2, // 消费
		BizID:        in.OrderID,
		ChangeAmount: -in.PayAmount, // 负数表示扣减
		Remark:       fmt.Sprintf("订单%s余额支付", in.OrderNo),
	})
	if err != nil {
		return err
	}

	// 2. 创建支付记录
	paymentID := snowflake.Generate()
	paymentNo := fmt.Sprintf("PAY%d", paymentID)
	_, err = dao.PlayPayment.Ctx(ctx).Data(g.Map{
		dao.PlayPayment.Columns().Id:        paymentID,
		dao.PlayPayment.Columns().OrderId:   in.OrderID,
		dao.PlayPayment.Columns().MemberId:  in.MemberID,
		dao.PlayPayment.Columns().PaymentNo: paymentNo,
		dao.PlayPayment.Columns().PayType:   3, // 余额支付
		dao.PlayPayment.Columns().PayAmount: in.PayAmount,
		dao.PlayPayment.Columns().PayStatus: 1, // 已支付
		dao.PlayPayment.Columns().PayAt:     gtime.Now(),
		dao.PlayPayment.Columns().CreatedAt: gtime.Now(),
		dao.PlayPayment.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	if err != nil {
		return err
	}

	// 3. 更新订单状态为已支付
	_, err = dao.PlayOrder.Ctx(ctx).Where(dao.PlayOrder.Columns().Id, in.OrderID).Data(g.Map{
		dao.PlayOrder.Columns().OrderStatus: 1,
		dao.PlayOrder.Columns().PayType:     3, // 余额支付
		dao.PlayOrder.Columns().PayAt:       gtime.Now(),
		dao.PlayOrder.Columns().UpdatedAt:   gtime.Now(),
	}).Update()
	return err
}
