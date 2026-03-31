package playapi

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/utility/snowflake"
)

type sRecharge struct{}

// Plans 充值方案列表
func (s *sRecharge) Plans(ctx context.Context) (list []v1.RechargePlanItem, err error) {
	var records []struct {
		Id         uint64 `json:"id"`
		Title      string `json:"title"`
		Amount     int64  `json:"amount"`
		GiftAmount int64  `json:"gift_amount"`
		Sort       int    `json:"sort"`
	}
	err = dao.PlayRechargePlan.Ctx(ctx).
		Where(dao.PlayRechargePlan.Columns().Status, 1).
		OrderAsc(dao.PlayRechargePlan.Columns().Sort).
		Scan(&records)
	if err != nil {
		return
	}
	list = make([]v1.RechargePlanItem, 0, len(records))
	for _, r := range records {
		list = append(list, v1.RechargePlanItem{
			PlanID:     strconv.FormatUint(r.Id, 10),
			Title:      r.Title,
			Amount:     r.Amount,
			GiveAmount: r.GiftAmount,
			Sort:       r.Sort,
		})
	}
	return
}

// Create 创建充值订单
func (s *sRecharge) Create(ctx context.Context, memberID int64, planID string, payType string) (orderID string, payParams string, err error) {
	pid, _ := strconv.ParseUint(planID, 10, 64)
	// 查询充值方案
	plan, err := dao.PlayRechargePlan.Ctx(ctx).
		Where(dao.PlayRechargePlan.Columns().Id, pid).
		Where(dao.PlayRechargePlan.Columns().Status, 1).
		Where(dao.PlayRechargePlan.Columns().DeletedAt, nil).
		One()
	if err != nil {
		return
	}
	if plan.IsEmpty() {
		err = gerror.New("充值方案不存在")
		return
	}
	// 生成订单
	id := snowflake.Generate()
	orderNo := fmt.Sprintf("RC%s%06d", gtime.Now().Format("YmdHis"), id%1000000)
	payTypeInt := 1
	if payType == "alipay" {
		payTypeInt = 2
	}
	_, err = dao.PlayRechargeOrder.Ctx(ctx).Data(g.Map{
		dao.PlayRechargeOrder.Columns().Id:             id,
		dao.PlayRechargeOrder.Columns().OrderNo:        orderNo,
		dao.PlayRechargeOrder.Columns().MemberId:       memberID,
		dao.PlayRechargeOrder.Columns().RechargePlanId: pid,
		dao.PlayRechargeOrder.Columns().Amount:         plan[dao.PlayRechargePlan.Columns().Amount].Int64(),
		dao.PlayRechargeOrder.Columns().GiftAmount:     plan[dao.PlayRechargePlan.Columns().GiftAmount].Int64(),
		dao.PlayRechargeOrder.Columns().PayType:        payTypeInt,
		dao.PlayRechargeOrder.Columns().PayStatus:      0,
		dao.PlayRechargeOrder.Columns().CreatedAt:      gtime.Now(),
		dao.PlayRechargeOrder.Columns().UpdatedAt:      gtime.Now(),
	}).Insert()
	if err != nil {
		return
	}
	orderID = strconv.FormatInt(int64(id), 10)
	// payParams 占位，实际对接第三方支付后填充
	payParams = "{}"
	return
}

// WxNotify 充值微信回调（占位实现）
func (s *sRecharge) WxNotify(ctx context.Context) error {
	g.Log().Info(ctx, "收到充值微信回调，待对接支付SDK")
	return nil
}

// AlipayNotify 充值支付宝回调（占位实现）
func (s *sRecharge) AlipayNotify(ctx context.Context) error {
	g.Log().Info(ctx, "收到充值支付宝回调，待对接支付SDK")
	return nil
}
