package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 可领取优惠券列表（公开） ====================

type CouponAvailableReq struct {
	g.Meta   `path:"/coupon/available" method:"get" tags:"C端优惠券" summary:"可领取优惠券列表"`
	Page     int `json:"page" v:"min:1" dc:"页码" d:"1"`
	PageSize int `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}

type CouponAvailableRes struct {
	g.Meta `mime:"application/json"`
	Total  int                    `json:"total" dc:"总数"`
	List   []CouponAvailableItem  `json:"list" dc:"优惠券列表"`
}

type CouponAvailableItem struct {
	CouponID     string `json:"couponId" dc:"优惠券ID"`
	Title        string `json:"title" dc:"优惠券标题"`
	Type         int    `json:"type" dc:"类型:1=满减,2=折扣,3=无门槛"`
	FaceValue    int64  `json:"faceValue" dc:"面值(分)/折扣百分比"`
	MinAmount    int64  `json:"minAmount" dc:"最低使用金额(分)，0=无门槛"`
	StartTime    string `json:"startTime" dc:"有效期开始"`
	EndTime      string `json:"endTime" dc:"有效期结束"`
	TotalNum     int    `json:"totalNum" dc:"发放总量"`
	ClaimNum     int    `json:"claimNum" dc:"已领取数量"`
}

// ==================== 领取优惠券（MemberAuth） ====================

type CouponReceiveReq struct {
	g.Meta   `path:"/coupon/receive" method:"post" tags:"C端优惠券" summary:"领取优惠券"`
	CouponID string `json:"couponId" v:"required#优惠券ID不能为空" dc:"优惠券ID"`
}

type CouponReceiveRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 我的优惠券（MemberAuth） ====================

type CouponMineReq struct {
	g.Meta   `path:"/coupon/mine" method:"get" tags:"C端优惠券" summary:"我的优惠券"`
	Status   *int `json:"status" dc:"状态筛选:0=未使用,1=已使用,2=已过期"`
	Page     int  `json:"page" v:"min:1" dc:"页码" d:"1"`
	PageSize int  `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}

type CouponMineRes struct {
	g.Meta `mime:"application/json"`
	Total  int              `json:"total" dc:"总数"`
	List   []CouponMineItem `json:"list" dc:"我的优惠券列表"`
}

type CouponMineItem struct {
	CouponMemberID string `json:"couponMemberId" dc:"用户优惠券ID"`
	CouponID       string `json:"couponId" dc:"优惠券ID"`
	Title          string `json:"title" dc:"优惠券标题"`
	Type           int    `json:"type" dc:"类型:1=满减,2=折扣,3=无门槛"`
	FaceValue      int64  `json:"faceValue" dc:"面值(分)/折扣百分比"`
	MinAmount      int64  `json:"minAmount" dc:"最低使用金额(分)"`
	StartTime      string `json:"startTime" dc:"有效期开始"`
	EndTime        string `json:"endTime" dc:"有效期结束"`
	UseStatus      int    `json:"useStatus" dc:"状态:0=未使用,1=已使用,2=已过期"`
	ClaimAt        string `json:"claimAt" dc:"领取时间"`
}

// ==================== 下单可用优惠券（MemberAuth） ====================

type CouponUsableApiReq struct {
	g.Meta      `path:"/coupon/usable" method:"get" tags:"C端优惠券" summary:"下单可用优惠券列表"`
	OrderAmount int64 `json:"orderAmount" v:"required|min:1#订单金额不能为空|订单金额必须大于0" dc:"订单金额(分)"`
}

type CouponUsableApiRes struct {
	g.Meta `mime:"application/json"`
	List   []CouponMineItem `json:"list" dc:"可用优惠券列表"`
}
