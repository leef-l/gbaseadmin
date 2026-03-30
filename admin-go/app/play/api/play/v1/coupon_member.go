package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// CouponMember API

// CouponMemberCreateReq 创建会员优惠券表请求
type CouponMemberCreateReq struct {
	g.Meta `path:"/coupon_member/create" method:"post" tags:"会员优惠券表" summary:"创建会员优惠券表"`
	CouponID snowflake.JsonInt64 `json:"couponID" v:"required#优惠券模板ID不能为空" dc:"优惠券模板ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#会员ID不能为空" dc:"会员ID"`
	OrderID snowflake.JsonInt64 `json:"orderID"  dc:"使用的订单ID（0表示未使用）"`
	UseStatus int `json:"useStatus"  dc:"使用状态"`
	ClaimAt *gtime.Time `json:"claimAt"  dc:"领取时间"`
	UseAt *gtime.Time `json:"useAt"  dc:"使用时间"`
	ExpireAt *gtime.Time `json:"expireAt"  dc:"过期时间"`
}

// CouponMemberCreateRes 创建会员优惠券表响应
type CouponMemberCreateRes struct {
	g.Meta `mime:"application/json"`
}

// CouponMemberUpdateReq 更新会员优惠券表请求
type CouponMemberUpdateReq struct {
	g.Meta `path:"/coupon_member/update" method:"put" tags:"会员优惠券表" summary:"更新会员优惠券表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"会员优惠券表ID"`
	CouponID snowflake.JsonInt64 `json:"couponID" dc:"优惠券模板ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"会员ID"`
	OrderID snowflake.JsonInt64 `json:"orderID" dc:"使用的订单ID（0表示未使用）"`
	UseStatus int `json:"useStatus" dc:"使用状态"`
	ClaimAt *gtime.Time `json:"claimAt" dc:"领取时间"`
	UseAt *gtime.Time `json:"useAt" dc:"使用时间"`
	ExpireAt *gtime.Time `json:"expireAt" dc:"过期时间"`
}

// CouponMemberUpdateRes 更新会员优惠券表响应
type CouponMemberUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// CouponMemberDeleteReq 删除会员优惠券表请求
type CouponMemberDeleteReq struct {
	g.Meta `path:"/coupon_member/delete" method:"delete" tags:"会员优惠券表" summary:"删除会员优惠券表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"会员优惠券表ID"`
}

// CouponMemberDeleteRes 删除会员优惠券表响应
type CouponMemberDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// CouponMemberDetailReq 获取会员优惠券表详情请求
type CouponMemberDetailReq struct {
	g.Meta `path:"/coupon_member/detail" method:"get" tags:"会员优惠券表" summary:"获取会员优惠券表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"会员优惠券表ID"`
}

// CouponMemberDetailRes 获取会员优惠券表详情响应
type CouponMemberDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.CouponMemberDetailOutput
}

// CouponMemberListReq 获取会员优惠券表列表请求
type CouponMemberListReq struct {
	g.Meta   `path:"/coupon_member/list" method:"get" tags:"会员优惠券表" summary:"获取会员优惠券表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	UseStatus int `json:"useStatus" dc:"使用状态"`
}

// CouponMemberListRes 获取会员优惠券表列表响应
type CouponMemberListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.CouponMemberListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

