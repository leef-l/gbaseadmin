package order

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model"
)

// ChangeStatus 变更订单状态
func (s *sOrder) ChangeStatus(ctx context.Context, in *model.OrderChangeStatusInput) error {
	data := g.Map{
		dao.PlayOrder.Columns().OrderStatus: in.OrderStatus,
		dao.PlayOrder.Columns().UpdatedAt:   gtime.Now(),
	}
	switch in.OrderStatus {
	case 1: // 已支付
		data[dao.PlayOrder.Columns().PayAt] = gtime.Now()
	case 2: // 进行中
		data[dao.PlayOrder.Columns().StartAt] = gtime.Now()
	case 3: // 已完成
		data[dao.PlayOrder.Columns().FinishAt] = gtime.Now()
	case 4: // 已取消
		data[dao.PlayOrder.Columns().CancelAt] = gtime.Now()
		data[dao.PlayOrder.Columns().CancelReason] = in.CancelReason
	}
	_, err := dao.PlayOrder.Ctx(ctx).Where(dao.PlayOrder.Columns().Id, in.ID).Data(data).Update()
	return err
}
