package playapi

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	v1 "gbaseadmin/api/playapi/v1"
	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
	"gbaseadmin/utility/snowflake"
)

type sPlayapiOrder struct{}

func init() {
	service.RegisterPlayapiOrder(&sPlayapiOrder{})
}

func (s *sPlayapiOrder) Create(ctx context.Context, req *v1.OrderCreateReq) (res *v1.OrderCreateRes, err error) {
	r := g.RequestFromCtx(ctx)
	memberID := r.GetCtxVar("jwt_member_id").Int64()

	gc := dao.PlayGoods.Columns()
	goods, err := dao.PlayGoods.Ctx(ctx).Where(gc.Id, req.GoodsID).Where(gc.Status, 1).One()
	if err != nil {
		return
	}
	if goods.IsEmpty() {
		err = gerror.New("商品不存在或已下架")
		return
	}

	coachID := goods[gc.CoachId].Int64()
	cc := dao.PlayCoach.Columns()
	coach, err := dao.PlayCoach.Ctx(ctx).Where(cc.Id, coachID).Where(cc.Status, 1).One()
	if err != nil {
		return
	}
	if coach.IsEmpty() {
		err = gerror.New("陪玩师不可用")
		return
	}

	// 不能购买自己的商品
	coachMemberID := coach[cc.MemberId].Int64()
	if coachMemberID == memberID {
		err = gerror.New("不能购买自己的商品")
		return
	}

	// 计算金额
	goodsPrice := goods[gc.Price].Int64()
	totalAmount := goodsPrice * int64(req.Quantity)
	discountAmount := int64(0)
	couponAmount := int64(0)

	// 查会员等级折扣
	mc := dao.PlayMember.Columns()
	mlc := dao.PlayMemberLevel.Columns()
	memberInfo, _ := dao.PlayMember.Ctx(ctx).Where(mc.Id, memberID).One()
	if !memberInfo.IsEmpty() {
		levelID := memberInfo[mc.MemberLevelId]
		if levelID != nil {
			discountVal, _ := dao.PlayMemberLevel.Ctx(ctx).Where(mlc.Id, levelID).Where(mlc.Status, 1).Value(mlc.Discount)
			if discountVal != nil && discountVal.Int() > 0 && discountVal.Int() < 100 {
				discountAmount = totalAmount * int64(100-discountVal.Int()) / 100
			}
		}
	}

	// 处理优惠券
	couponMemberID := int64(0)
	if req.CouponMemberID != "" {
		couponMemberID, _ = strconv.ParseInt(req.CouponMemberID, 10, 64)
		if couponMemberID > 0 {
			cmc := dao.PlayCouponMember.Columns()
			cm, e := dao.PlayCouponMember.Ctx(ctx).Where(cmc.Id, couponMemberID).Where(cmc.MemberId, memberID).Where(cmc.UseStatus, 0).One()
			if e != nil || cm.IsEmpty() {
				err = gerror.New("优惠券不可用")
				return
			}
			// 简化：直接用优惠券关联的coupon金额
			couponC := dao.PlayCoupon.Columns()
			couponVal, _ := dao.PlayCoupon.Ctx(ctx).Where(couponC.Id, cm[cmc.CouponId]).Value(couponC.FaceValue)
			if couponVal != nil {
				couponAmount = couponVal.Int64()
			}
		}
	}

	payAmount := totalAmount - discountAmount - couponAmount
	if payAmount < 0 {
		payAmount = 0
	}

	// 生成订单
	orderID := snowflake.Generate()
	orderNo := fmt.Sprintf("PW%s%06d", time.Now().Format("20060102150405"), rand.Intn(1000000))

	err = dao.PlayOrder.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		oc := dao.PlayOrder.Columns()
		_, e := dao.PlayOrder.Ctx(ctx).Data(g.Map{
			oc.Id:             orderID,
			oc.OrderNo:        orderNo,
			oc.MemberId:       memberID,
			oc.CoachId:        coachID,
			oc.GoodsId:        req.GoodsID,
			oc.GoodsTitle:     goods[gc.Title].String(),
			oc.GoodsPrice:     goodsPrice,
			oc.Quantity:       req.Quantity,
			oc.TotalAmount:    totalAmount,
			oc.DiscountAmount: discountAmount,
			oc.CouponAmount:   couponAmount,
			oc.PayAmount:      payAmount,
			oc.CouponMemberId: couponMemberID,
			oc.OrderStatus:    0,
			oc.Remark:         req.Remark,
			oc.CreatedAt:      gtime.Now(),
			oc.UpdatedAt:      gtime.Now(),
		}).Insert()
		if e != nil {
			return e
		}
		// 锁定优惠券
		if couponMemberID > 0 {
			cmc := dao.PlayCouponMember.Columns()
			_, e = dao.PlayCouponMember.Ctx(ctx).Where(cmc.Id, couponMemberID).Data(g.Map{
				cmc.UseStatus:  1,
				cmc.OrderId:    orderID,
				cmc.UpdatedAt:  gtime.Now(),
			}).Update()
		}
		return e
	})
	if err != nil {
		return
	}

	res = &v1.OrderCreateRes{
		OrderID:   fmt.Sprintf("%d", orderID),
		OrderNo:   orderNo,
		PayAmount: payAmount,
	}
	return
}

func (s *sPlayapiOrder) List(ctx context.Context, req *v1.OrderListReq) (list []v1.OrderListItem, total int, err error) {
	r := g.RequestFromCtx(ctx)
	memberID := r.GetCtxVar("jwt_member_id").Int64()
	oc := dao.PlayOrder.Columns()

	m := dao.PlayOrder.Ctx(ctx).Where(oc.MemberId, memberID)
	if req.Status != nil {
		m = m.Where(oc.OrderStatus, *req.Status)
	}

	total, err = m.Count()
	if err != nil {
		return
	}

	records, err := m.OrderDesc(oc.CreatedAt).Page(req.Page, req.PageSize).All()
	if err != nil {
		return
	}

	mc := dao.PlayMember.Columns()
	cc := dao.PlayCoach.Columns()
	gc := dao.PlayGoods.Columns()

	list = make([]v1.OrderListItem, 0, len(records))
	for _, o := range records {
		// 获取陪玩师信息
		coachName := ""
		coachAvatar := ""
		coach, _ := dao.PlayCoach.Ctx(ctx).Where(cc.Id, o[oc.CoachId]).One()
		if !coach.IsEmpty() {
			mem, _ := dao.PlayMember.Ctx(ctx).Where(mc.Id, coach[cc.MemberId]).One()
			if !mem.IsEmpty() {
				coachName = mem[mc.Nickname].String()
				coachAvatar = mem[mc.Avatar].String()
			}
		}
		goodsImg := ""
		imgVal, _ := dao.PlayGoods.Ctx(ctx).Where(gc.Id, o[oc.GoodsId]).Value(gc.CoverImage)
		if imgVal != nil {
			goodsImg = imgVal.String()
		}
		list = append(list, v1.OrderListItem{
			OrderID:     o[oc.Id].String(),
			OrderNo:     o[oc.OrderNo].String(),
			GoodsTitle:  o[oc.GoodsTitle].String(),
			GoodsImage:  goodsImg,
			CoachID:     o[oc.CoachId].String(),
			CoachName:   coachName,
			CoachAvatar: coachAvatar,
			Quantity:    o[oc.Quantity].Int(),
			TotalAmount: o[oc.TotalAmount].Int64(),
			PayAmount:   o[oc.PayAmount].Int64(),
			OrderStatus: o[oc.OrderStatus].Int(),
			CreatedAt:   o[oc.CreatedAt].String(),
		})
	}
	return
}

func (s *sPlayapiOrder) Detail(ctx context.Context, req *v1.OrderDetailReq) (res *v1.OrderDetailRes, err error) {
	r := g.RequestFromCtx(ctx)
	memberID := r.GetCtxVar("jwt_member_id").Int64()
	oc := dao.PlayOrder.Columns()

	order, err := dao.PlayOrder.Ctx(ctx).Where(oc.Id, req.OrderID).Where(oc.MemberId, memberID).One()
	if err != nil {
		return
	}
	if order.IsEmpty() {
		err = gerror.New("订单不存在")
		return
	}

	mc := dao.PlayMember.Columns()
	cc := dao.PlayCoach.Columns()
	gc := dao.PlayGoods.Columns()

	coachName := ""
	coachAvatar := ""
	coach, _ := dao.PlayCoach.Ctx(ctx).Where(cc.Id, order[oc.CoachId]).One()
	if !coach.IsEmpty() {
		mem, _ := dao.PlayMember.Ctx(ctx).Where(mc.Id, coach[cc.MemberId]).One()
		if !mem.IsEmpty() {
			coachName = mem[mc.Nickname].String()
			coachAvatar = mem[mc.Avatar].String()
		}
	}
	goodsImg := ""
	imgVal, _ := dao.PlayGoods.Ctx(ctx).Where(gc.Id, order[oc.GoodsId]).Value(gc.CoverImage)
	if imgVal != nil {
		goodsImg = imgVal.String()
	}

	// 是否已评价
	rc := dao.PlayReview.Columns()
	reviewCount, _ := dao.PlayReview.Ctx(ctx).Where(rc.OrderId, req.OrderID).Where(rc.MemberId, memberID).Count()

	res = &v1.OrderDetailRes{
		OrderID:        order[oc.Id].String(),
		OrderNo:        order[oc.OrderNo].String(),
		GoodsID:        order[oc.GoodsId].String(),
		GoodsTitle:     order[oc.GoodsTitle].String(),
		GoodsImage:     goodsImg,
		GoodsPrice:     order[oc.GoodsPrice].Int64(),
		CoachID:        order[oc.CoachId].String(),
		CoachName:      coachName,
		CoachAvatar:    coachAvatar,
		Quantity:       order[oc.Quantity].Int(),
		TotalAmount:    order[oc.TotalAmount].Int64(),
		DiscountAmount: order[oc.DiscountAmount].Int64(),
		CouponAmount:   order[oc.CouponAmount].Int64(),
		PayAmount:      order[oc.PayAmount].Int64(),
		PayType:        order[oc.PayType].Int(),
		OrderStatus:    order[oc.OrderStatus].Int(),
		Remark:         order[oc.Remark].String(),
		CreatedAt:      order[oc.CreatedAt].String(),
		PayAt:          order[oc.PayAt].String(),
		StartAt:        order[oc.StartAt].String(),
		FinishAt:       order[oc.FinishAt].String(),
		CancelAt:       order[oc.CancelAt].String(),
		CancelReason:   order[oc.CancelReason].String(),
		HasReview:      reviewCount > 0,
	}
	return
}

func (s *sPlayapiOrder) Cancel(ctx context.Context, req *v1.OrderCancelReq) error {
	r := g.RequestFromCtx(ctx)
	memberID := r.GetCtxVar("jwt_member_id").Int64()
	oc := dao.PlayOrder.Columns()

	order, err := dao.PlayOrder.Ctx(ctx).Where(oc.Id, req.OrderID).Where(oc.MemberId, memberID).One()
	if err != nil {
		return err
	}
	if order.IsEmpty() {
		return gerror.New("订单不存在")
	}
	if order[oc.OrderStatus].Int() != 0 {
		return gerror.New("只有待支付订单可以取消")
	}

	return service.OrderEnhance().ChangeStatus(ctx, &model.OrderChangeStatusInput{
		ID:           snowflake.JsonInt64(order[oc.Id].Int64()),
		OrderStatus:  4,
		CancelReason: req.CancelReason,
	})
}

func (s *sPlayapiOrder) Refund(ctx context.Context, req *v1.OrderRefundReq) error {
	r := g.RequestFromCtx(ctx)
	memberID := r.GetCtxVar("jwt_member_id").Int64()
	oc := dao.PlayOrder.Columns()

	order, err := dao.PlayOrder.Ctx(ctx).Where(oc.Id, req.OrderID).Where(oc.MemberId, memberID).One()
	if err != nil {
		return err
	}
	if order.IsEmpty() {
		return gerror.New("订单不存在")
	}
	status := order[oc.OrderStatus].Int()
	if status != 1 && status != 2 {
		return gerror.New("当前状态不可申请退款")
	}

	return service.OrderEnhance().ChangeStatus(ctx, &model.OrderChangeStatusInput{
		ID:          snowflake.JsonInt64(order[oc.Id].Int64()),
		OrderStatus: 5,
	})
}

func (s *sPlayapiOrder) Accept(ctx context.Context, req *v1.OrderAcceptReq) error {
	r := g.RequestFromCtx(ctx)
	coachID := r.GetCtxVar("jwt_coach_id").Int64()
	oc := dao.PlayOrder.Columns()

	order, err := dao.PlayOrder.Ctx(ctx).Where(oc.Id, req.OrderID).Where(oc.CoachId, coachID).One()
	if err != nil {
		return err
	}
	if order.IsEmpty() {
		return gerror.New("订单不存在")
	}
	if order[oc.OrderStatus].Int() != 1 {
		return gerror.New("当前状态不可接单")
	}

	return service.OrderEnhance().ChangeStatus(ctx, &model.OrderChangeStatusInput{
		ID:          snowflake.JsonInt64(order[oc.Id].Int64()),
		OrderStatus: 2,
	})
}

func (s *sPlayapiOrder) Complete(ctx context.Context, req *v1.OrderCompleteReq) error {
	r := g.RequestFromCtx(ctx)
	coachID := r.GetCtxVar("jwt_coach_id").Int64()
	oc := dao.PlayOrder.Columns()

	order, err := dao.PlayOrder.Ctx(ctx).Where(oc.Id, req.OrderID).Where(oc.CoachId, coachID).One()
	if err != nil {
		return err
	}
	if order.IsEmpty() {
		return gerror.New("订单不存在")
	}
	if order[oc.OrderStatus].Int() != 2 {
		return gerror.New("当前状态不可完成")
	}

	return service.OrderEnhance().ChangeStatus(ctx, &model.OrderChangeStatusInput{
		ID:          snowflake.JsonInt64(order[oc.Id].Int64()),
		OrderStatus: 3,
	})
}
