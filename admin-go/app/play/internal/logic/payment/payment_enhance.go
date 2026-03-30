package payment

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
	"gbaseadmin/utility/snowflake"
)

// BalancePay 余额支付（事务）
func (s *sPayment) BalancePay(ctx context.Context, in *model.BalancePayInput) error {
	return dao.PlayOrder.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 1. 查询订单信息并校验状态
		order, err := dao.PlayOrder.Ctx(ctx).Where(dao.PlayOrder.Columns().Id, in.OrderID).One()
		if err != nil {
			return err
		}
		if order.IsEmpty() {
			return gerror.New("订单不存在")
		}
		if order[dao.PlayOrder.Columns().OrderStatus].Int() != 0 {
			return gerror.New("订单状态不是待支付")
		}
		payAmount := order[dao.PlayOrder.Columns().PayAmount].Int64()
		orderNo := order[dao.PlayOrder.Columns().OrderNo].String()

		// 2. 校验会员余额
		balVal, err := dao.PlayMember.Ctx(ctx).Where(dao.PlayMember.Columns().Id, in.MemberID).Value(dao.PlayMember.Columns().Balance)
		if err != nil {
			return err
		}
		if balVal.Int64() < payAmount {
			return gerror.New("余额不足")
		}

		// 3. 扣减余额（通过 AddLog 统一入口）
		err = service.BalanceLogEnhance().AddLog(ctx, &model.AddBalanceLogInput{
			MemberID:     in.MemberID,
			BizType:      2, // 消费
			BizID:        in.OrderID,
			ChangeAmount: -payAmount,
			Remark:       fmt.Sprintf("订单%s余额支付", orderNo),
		})
		if err != nil {
			return err
		}

		// 4. 创建支付记录
		paymentID := snowflake.Generate()
		paymentNo := fmt.Sprintf("PAY%d", paymentID)
		_, err = dao.PlayPayment.Ctx(ctx).Data(g.Map{
			dao.PlayPayment.Columns().Id:        paymentID,
			dao.PlayPayment.Columns().OrderId:   in.OrderID,
			dao.PlayPayment.Columns().MemberId:  in.MemberID,
			dao.PlayPayment.Columns().PaymentNo: paymentNo,
			dao.PlayPayment.Columns().PayType:   3, // 余额支付
			dao.PlayPayment.Columns().PayAmount: payAmount,
			dao.PlayPayment.Columns().PayStatus: 1, // 已支付
			dao.PlayPayment.Columns().PayAt:     gtime.Now(),
			dao.PlayPayment.Columns().CreatedAt: gtime.Now(),
			dao.PlayPayment.Columns().UpdatedAt: gtime.Now(),
		}).Insert()
		if err != nil {
			return err
		}

		// 5. 更新订单状态为已支付
		_, err = dao.PlayOrder.Ctx(ctx).Where(dao.PlayOrder.Columns().Id, in.OrderID).Data(g.Map{
			dao.PlayOrder.Columns().OrderStatus: 1,
			dao.PlayOrder.Columns().PayType:     3,
			dao.PlayOrder.Columns().PayAt:       gtime.Now(),
			dao.PlayOrder.Columns().UpdatedAt:   gtime.Now(),
		}).Update()
		return err
	})
}
