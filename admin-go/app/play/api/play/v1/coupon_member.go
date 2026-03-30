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

// CouponMemberCreateReq 创建ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨请求
type CouponMemberCreateReq struct {
	g.Meta `path:"/coupon_member/create" method:"post" tags:"ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨" summary:"创建ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨"`
	CouponID snowflake.JsonInt64 `json:"couponID" v:"required#ä¼˜æƒ åˆ¸æ¨¡æ¿ID不能为空" dc:"ä¼˜æƒ åˆ¸æ¨¡æ¿ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#ä¼šå‘˜ID不能为空" dc:"ä¼šå‘˜ID"`
	OrderID snowflake.JsonInt64 `json:"orderID"  dc:"ä½¿ç”¨çš„è®¢å•ID"`
	UseStatus int `json:"useStatus"  dc:"ä½¿ç”¨çŠ¶æ€"`
	ClaimAt *gtime.Time `json:"claimAt"  dc:"é¢†å–æ—¶é—´"`
	UseAt *gtime.Time `json:"useAt"  dc:"ä½¿ç”¨æ—¶é—´"`
	ExpireAt *gtime.Time `json:"expireAt"  dc:"è¿‡æœŸæ—¶é—´"`
}

// CouponMemberCreateRes 创建ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨响应
type CouponMemberCreateRes struct {
	g.Meta `mime:"application/json"`
}

// CouponMemberUpdateReq 更新ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨请求
type CouponMemberUpdateReq struct {
	g.Meta `path:"/coupon_member/update" method:"put" tags:"ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨" summary:"更新ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨ID"`
	CouponID snowflake.JsonInt64 `json:"couponID" dc:"ä¼˜æƒ åˆ¸æ¨¡æ¿ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"ä¼šå‘˜ID"`
	OrderID snowflake.JsonInt64 `json:"orderID" dc:"ä½¿ç”¨çš„è®¢å•ID"`
	UseStatus int `json:"useStatus" dc:"ä½¿ç”¨çŠ¶æ€"`
	ClaimAt *gtime.Time `json:"claimAt" dc:"é¢†å–æ—¶é—´"`
	UseAt *gtime.Time `json:"useAt" dc:"ä½¿ç”¨æ—¶é—´"`
	ExpireAt *gtime.Time `json:"expireAt" dc:"è¿‡æœŸæ—¶é—´"`
}

// CouponMemberUpdateRes 更新ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨响应
type CouponMemberUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// CouponMemberDeleteReq 删除ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨请求
type CouponMemberDeleteReq struct {
	g.Meta `path:"/coupon_member/delete" method:"delete" tags:"ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨" summary:"删除ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨ID"`
}

// CouponMemberDeleteRes 删除ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨响应
type CouponMemberDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// CouponMemberDetailReq 获取ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨详情请求
type CouponMemberDetailReq struct {
	g.Meta `path:"/coupon_member/detail" method:"get" tags:"ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨" summary:"获取ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨ID"`
}

// CouponMemberDetailRes 获取ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨详情响应
type CouponMemberDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.CouponMemberDetailOutput
}

// CouponMemberListReq 获取ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨列表请求
type CouponMemberListReq struct {
	g.Meta   `path:"/coupon_member/list" method:"get" tags:"ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨" summary:"获取ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	UseStatus int `json:"useStatus" dc:"ä½¿ç”¨çŠ¶æ€"`
}

// CouponMemberListRes 获取ä¼šå‘˜ä¼˜æƒ åˆ¸è¡¨列表响应
type CouponMemberListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.CouponMemberListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

