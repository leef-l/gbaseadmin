package order

import (
	"context"

	v1 "gbaseadmin/app/play/api/play/v1"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
)

// ChangeStatus 变更订单状态
func (c *cOrder) ChangeStatus(ctx context.Context, req *v1.OrderChangeStatusReq) (res *v1.OrderChangeStatusRes, err error) {
	err = service.OrderEnhance().ChangeStatus(ctx, &model.OrderChangeStatusInput{
		ID:           req.ID,
		OrderStatus:  req.OrderStatus,
		CancelReason: req.CancelReason,
	})
	return
}
