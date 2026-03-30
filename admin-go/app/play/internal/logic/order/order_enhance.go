package order

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model"
)

// 合法状态流转映射
var validTransitions = map[int][]int{
	0: {1, 4},    // 待支付 → 已支付 / 已取消
	1: {2, 5},    // 已支付 → 进行中 / 退款中
	2: {3, 5},    // 进行中 → 已完成 / 退款中
	5: {6, 2},    // 退款中 → 已退款 / 拒绝退款回到进行中
}

// ChangeStatus 变更订单状态（含状态机校验）
func (s *sOrder) ChangeStatus(ctx context.Context, in *model.OrderChangeStatusInput) error {
	// 查询当前订单状态
	val, err := dao.PlayOrder.Ctx(ctx).Where(dao.PlayOrder.Columns().Id, in.ID).Value(dao.PlayOrder.Columns().OrderStatus)
	if err != nil {
		return err
	}
	currentStatus := val.Int()

	// 校验状态流转合法性
	allowed, ok := validTransitions[currentStatus]
	if !ok {
		return gerror.Newf("当前状态(%d)不允许变更", currentStatus)
	}
	valid := false
	for _, s := range allowed {
		if s == in.OrderStatus {
			valid = true
			break
		}
	}
	if !valid {
		return gerror.Newf("不允许从状态(%d)变更到状态(%d)", currentStatus, in.OrderStatus)
	}

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
	_, err = dao.PlayOrder.Ctx(ctx).Where(dao.PlayOrder.Columns().Id, in.ID).Data(data).Update()
	return err
}
