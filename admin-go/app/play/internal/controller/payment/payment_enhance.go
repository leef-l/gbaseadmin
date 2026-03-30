package payment

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

// PayByBalance 余额支付
func (c *cPayment) PayByBalance(ctx context.Context, req *v1.PayByBalanceReq) (res *v1.PayByBalanceRes, err error) {
	err = service.PaymentEnhance().BalancePay(ctx, &model.BalancePayInput{
		OrderID:  req.OrderID,
		MemberID: req.MemberID,
	})
	return
}
