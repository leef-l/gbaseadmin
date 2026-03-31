package playapi

import (
	"context"
	"strconv"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/utility/snowflake"
)

type sCoupon struct{}

// Available 可领取优惠券列表
func (s *sCoupon) Available(ctx context.Context, page, pageSize int) (list []v1.CouponAvailableItem, total int, err error) {
	m := dao.PlayCoupon.Ctx(ctx).
		Where(dao.PlayCoupon.Columns().Status, 1).
		WhereGTE(dao.PlayCoupon.Columns().ValidEndAt, gtime.Now()).
		WhereLTE(dao.PlayCoupon.Columns().ValidStartAt, gtime.Now()).
		Where(dao.PlayCoupon.Columns().DeletedAt, nil)
	// 排除已领完的: total_num>0 AND claim_num>=total_num
	m = m.Where("total_num=0 OR claim_num<total_num")

	total, err = m.Count()
	if err != nil {
		return
	}

	var records []struct {
		Id           uint64      `json:"id"`
		Title        string      `json:"title"`
		Type         int         `json:"type"`
		FaceValue    int64       `json:"face_value"`
		MinAmount    int64       `json:"min_amount"`
		ValidStartAt *gtime.Time `json:"valid_start_at"`
		ValidEndAt   *gtime.Time `json:"valid_end_at"`
		TotalNum     int         `json:"total_num"`
		ClaimNum     int         `json:"claim_num"`
	}
	err = m.Page(page, pageSize).
		OrderAsc(dao.PlayCoupon.Columns().Sort).
		OrderDesc(dao.PlayCoupon.Columns().CreatedAt).
		Scan(&records)
	if err != nil {
		return
	}
	list = make([]v1.CouponAvailableItem, 0, len(records))
	for _, r := range records {
		item := v1.CouponAvailableItem{
			CouponID:  strconv.FormatUint(r.Id, 10),
			Title:     r.Title,
			Type:      r.Type,
			FaceValue: r.FaceValue,
			MinAmount: r.MinAmount,
			TotalNum:  r.TotalNum,
			ClaimNum:  r.ClaimNum,
		}
		if r.ValidStartAt != nil {
			item.StartTime = r.ValidStartAt.String()
		}
		if r.ValidEndAt != nil {
			item.EndTime = r.ValidEndAt.String()
		}
		list = append(list, item)
	}
	return
}

// Receive 领取优惠券
func (s *sCoupon) Receive(ctx context.Context, memberID int64, couponID string) error {
	cid, _ := strconv.ParseUint(couponID, 10, 64)
	// 查询优惠券
	coupon, err := dao.PlayCoupon.Ctx(ctx).
		Where(dao.PlayCoupon.Columns().Id, cid).
		Where(dao.PlayCoupon.Columns().Status, 1).
		WhereGTE(dao.PlayCoupon.Columns().ValidEndAt, gtime.Now()).
		Where(dao.PlayCoupon.Columns().DeletedAt, nil).
		One()
	if err != nil {
		return err
	}
	if coupon.IsEmpty() {
		return gerror.New("优惠券不存在或已过期")
	}
	// 校验是否已领完
	totalNum := coupon[dao.PlayCoupon.Columns().TotalNum].Int()
	claimNum := coupon[dao.PlayCoupon.Columns().ClaimNum].Int()
	if totalNum > 0 && claimNum >= totalNum {
		return gerror.New("优惠券已领完")
	}
	// 校验每人限领
	perLimit := coupon[dao.PlayCoupon.Columns().PerLimit].Int()
	if perLimit > 0 {
		cnt, err := dao.PlayCouponMember.Ctx(ctx).
			Where(dao.PlayCouponMember.Columns().CouponId, cid).
			Where(dao.PlayCouponMember.Columns().MemberId, memberID).
			Where(dao.PlayCouponMember.Columns().DeletedAt, nil).
			Count()
		if err != nil {
			return err
		}
		if cnt >= perLimit {
			return gerror.New("您已领取过该优惠券")
		}
	}
	// 事务：插入领取记录 + 更新已领数量
	return dao.PlayCouponMember.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		id := snowflake.Generate()
		_, err := dao.PlayCouponMember.Ctx(ctx).Data(g.Map{
			dao.PlayCouponMember.Columns().Id:        id,
			dao.PlayCouponMember.Columns().CouponId:  cid,
			dao.PlayCouponMember.Columns().MemberId:  memberID,
			dao.PlayCouponMember.Columns().UseStatus: 0,
			dao.PlayCouponMember.Columns().ClaimAt:   gtime.Now(),
			dao.PlayCouponMember.Columns().ExpireAt:  coupon[dao.PlayCoupon.Columns().ValidEndAt].GTime(),
			dao.PlayCouponMember.Columns().CreatedAt: gtime.Now(),
			dao.PlayCouponMember.Columns().UpdatedAt: gtime.Now(),
		}).Insert()
		if err != nil {
			return err
		}
		_, err = dao.PlayCoupon.Ctx(ctx).
			Where(dao.PlayCoupon.Columns().Id, cid).
			Increment(dao.PlayCoupon.Columns().ClaimNum, 1)
		return err
	})
}

// Mine 我的优惠券
func (s *sCoupon) Mine(ctx context.Context, memberID int64, status *int, page, pageSize int) (list []v1.CouponMineItem, total int, err error) {
	m := dao.PlayCouponMember.Ctx(ctx).
		Where(dao.PlayCouponMember.Columns().MemberId, memberID).
		Where(dao.PlayCouponMember.Columns().DeletedAt, nil)
	if status != nil {
		m = m.Where(dao.PlayCouponMember.Columns().UseStatus, *status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	var records []struct {
		Id        uint64      `json:"id"`
		CouponId  uint64      `json:"coupon_id"`
		UseStatus int         `json:"use_status"`
		ClaimAt   *gtime.Time `json:"claim_at"`
		ExpireAt  *gtime.Time `json:"expire_at"`
	}
	err = m.Page(page, pageSize).
		OrderAsc(dao.PlayCouponMember.Columns().UseStatus).
		OrderDesc(dao.PlayCouponMember.Columns().ClaimAt).
		Scan(&records)
	if err != nil {
		return
	}
	list = make([]v1.CouponMineItem, 0, len(records))
	for _, r := range records {
		// 查优惠券模板信息
		coupon, e := dao.PlayCoupon.Ctx(ctx).Where(dao.PlayCoupon.Columns().Id, r.CouponId).One()
		if e != nil || coupon.IsEmpty() {
			continue
		}
		useStatus := r.UseStatus
		// 自动标记过期
		if useStatus == 0 && r.ExpireAt != nil && r.ExpireAt.Before(gtime.Now()) {
			useStatus = 2
		}
		item := v1.CouponMineItem{
			CouponMemberID: strconv.FormatUint(r.Id, 10),
			CouponID:       strconv.FormatUint(r.CouponId, 10),
			Title:          coupon[dao.PlayCoupon.Columns().Title].String(),
			Type:           coupon[dao.PlayCoupon.Columns().Type].Int(),
			FaceValue:      coupon[dao.PlayCoupon.Columns().FaceValue].Int64(),
			MinAmount:      coupon[dao.PlayCoupon.Columns().MinAmount].Int64(),
			UseStatus:      useStatus,
		}
		if coupon[dao.PlayCoupon.Columns().ValidStartAt].GTime() != nil {
			item.StartTime = coupon[dao.PlayCoupon.Columns().ValidStartAt].String()
		}
		if coupon[dao.PlayCoupon.Columns().ValidEndAt].GTime() != nil {
			item.EndTime = coupon[dao.PlayCoupon.Columns().ValidEndAt].String()
		}
		if r.ClaimAt != nil {
			item.ClaimAt = r.ClaimAt.String()
		}
		list = append(list, item)
	}
	return
}

// Usable 下单可用优惠券
func (s *sCoupon) Usable(ctx context.Context, memberID int64, orderAmount int64) (list []v1.CouponMineItem, err error) {
	var records []struct {
		Id       uint64 `json:"id"`
		CouponId uint64 `json:"coupon_id"`
		ClaimAt  *gtime.Time `json:"claim_at"`
	}
	err = dao.PlayCouponMember.Ctx(ctx).
		Where(dao.PlayCouponMember.Columns().MemberId, memberID).
		Where(dao.PlayCouponMember.Columns().UseStatus, 0).
		WhereGTE(dao.PlayCouponMember.Columns().ExpireAt, gtime.Now()).
		Where(dao.PlayCouponMember.Columns().DeletedAt, nil).
		Scan(&records)
	if err != nil {
		return
	}
	list = make([]v1.CouponMineItem, 0)
	for _, r := range records {
		coupon, e := dao.PlayCoupon.Ctx(ctx).Where(dao.PlayCoupon.Columns().Id, r.CouponId).One()
		if e != nil || coupon.IsEmpty() {
			continue
		}
		minAmount := coupon[dao.PlayCoupon.Columns().MinAmount].Int64()
		if minAmount > 0 && orderAmount < minAmount {
			continue
		}
		item := v1.CouponMineItem{
			CouponMemberID: strconv.FormatUint(r.Id, 10),
			CouponID:       strconv.FormatUint(r.CouponId, 10),
			Title:          coupon[dao.PlayCoupon.Columns().Title].String(),
			Type:           coupon[dao.PlayCoupon.Columns().Type].Int(),
			FaceValue:      coupon[dao.PlayCoupon.Columns().FaceValue].Int64(),
			MinAmount:      minAmount,
			UseStatus:      0,
		}
		if coupon[dao.PlayCoupon.Columns().ValidStartAt].GTime() != nil {
			item.StartTime = coupon[dao.PlayCoupon.Columns().ValidStartAt].String()
		}
		if coupon[dao.PlayCoupon.Columns().ValidEndAt].GTime() != nil {
			item.EndTime = coupon[dao.PlayCoupon.Columns().ValidEndAt].String()
		}
		if r.ClaimAt != nil {
			item.ClaimAt = r.ClaimAt.String()
		}
		list = append(list, item)
	}
	return
}
