package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// Coupon API

// CouponCreateReq 创建优惠券模板表请求
type CouponCreateReq struct {
	g.Meta `path:"/coupon/create" method:"post" tags:"优惠券模板表" summary:"创建优惠券模板表"`
	Title string `json:"title" v:"required#优惠券名称不能为空" dc:"优惠券名称"`
	Type int `json:"type"  dc:"优惠券类型"`
	IsNewMember int `json:"isNewMember"  dc:"是否新人专享"`
	FaceValue int64 `json:"faceValue"  dc:"面值（分，满减/无门槛时为抵扣额，折扣时为折扣值如 85=8.5折）"`
	MinAmount int64 `json:"minAmount"  dc:"最低消费金额（分，0表示无门槛）"`
	TotalNum int `json:"totalNum"  dc:"发放总量（0表示不限）"`
	UsedNum int `json:"usedNum"  dc:"已使用数量"`
	ClaimNum int `json:"claimNum"  dc:"已领取数量"`
	PerLimit int `json:"perLimit"  dc:"每人限领张数"`
	ValidStartAt *gtime.Time `json:"validStartAt" v:"required#有效期开始时间不能为空" dc:"有效期开始时间"`
	ValidEndAt *gtime.Time `json:"validEndAt" v:"required#有效期结束时间不能为空" dc:"有效期结束时间"`
	Sort int `json:"sort"  dc:"排序（升序）"`
	Status int `json:"status"  dc:"状态"`
}

// CouponCreateRes 创建优惠券模板表响应
type CouponCreateRes struct {
	g.Meta `mime:"application/json"`
}

// CouponUpdateReq 更新优惠券模板表请求
type CouponUpdateReq struct {
	g.Meta `path:"/coupon/update" method:"put" tags:"优惠券模板表" summary:"更新优惠券模板表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"优惠券模板表ID"`
	Title string `json:"title" dc:"优惠券名称"`
	Type int `json:"type" dc:"优惠券类型"`
	IsNewMember int `json:"isNewMember" dc:"是否新人专享"`
	FaceValue int64 `json:"faceValue" dc:"面值（分，满减/无门槛时为抵扣额，折扣时为折扣值如 85=8.5折）"`
	MinAmount int64 `json:"minAmount" dc:"最低消费金额（分，0表示无门槛）"`
	TotalNum int `json:"totalNum" dc:"发放总量（0表示不限）"`
	UsedNum int `json:"usedNum" dc:"已使用数量"`
	ClaimNum int `json:"claimNum" dc:"已领取数量"`
	PerLimit int `json:"perLimit" dc:"每人限领张数"`
	ValidStartAt *gtime.Time `json:"validStartAt" dc:"有效期开始时间"`
	ValidEndAt *gtime.Time `json:"validEndAt" dc:"有效期结束时间"`
	Sort int `json:"sort" dc:"排序（升序）"`
	Status int `json:"status" dc:"状态"`
}

// CouponUpdateRes 更新优惠券模板表响应
type CouponUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// CouponDeleteReq 删除优惠券模板表请求
type CouponDeleteReq struct {
	g.Meta `path:"/coupon/delete" method:"delete" tags:"优惠券模板表" summary:"删除优惠券模板表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"优惠券模板表ID"`
}

// CouponDeleteRes 删除优惠券模板表响应
type CouponDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// CouponDetailReq 获取优惠券模板表详情请求
type CouponDetailReq struct {
	g.Meta `path:"/coupon/detail" method:"get" tags:"优惠券模板表" summary:"获取优惠券模板表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"优惠券模板表ID"`
}

// CouponDetailRes 获取优惠券模板表详情响应
type CouponDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.CouponDetailOutput
}

// CouponListReq 获取优惠券模板表列表请求
type CouponListReq struct {
	g.Meta   `path:"/coupon/list" method:"get" tags:"优惠券模板表" summary:"获取优惠券模板表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Type int `json:"type" dc:"优惠券类型"`
	IsNewMember int `json:"isNewMember" dc:"是否新人专享"`
	Status int `json:"status" dc:"状态"`
}

// CouponListRes 获取优惠券模板表列表响应
type CouponListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.CouponListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

