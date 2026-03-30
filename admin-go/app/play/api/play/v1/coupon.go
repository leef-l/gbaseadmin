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

// CouponCreateReq 创建ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨请求
type CouponCreateReq struct {
	g.Meta `path:"/coupon/create" method:"post" tags:"ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨" summary:"创建ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨"`
	Title string `json:"title" v:"required#ä¼˜æƒ åˆ¸åç§°不能为空" dc:"ä¼˜æƒ åˆ¸åç§°"`
	Type int `json:"type"  dc:"ä¼˜æƒ åˆ¸ç±»åž‹"`
	IsNewMember int `json:"isNewMember"  dc:"æ˜¯å¦æ–°äººä¸“äº«"`
	FaceValue int64 `json:"faceValue"  dc:"é¢å€¼ï¼ˆåˆ†ï¼‰"`
	MinAmount int64 `json:"minAmount"  dc:"æœ€ä½Žæ¶ˆè´¹é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	TotalNum int `json:"totalNum"  dc:"å‘æ”¾æ€»é‡ï¼ˆ0ä¸é™ï¼‰"`
	UsedNum int `json:"usedNum"  dc:"å·²ä½¿ç”¨æ•°é‡"`
	ClaimNum int `json:"claimNum"  dc:"å·²é¢†å–æ•°é‡"`
	PerLimit int `json:"perLimit"  dc:"æ¯äººé™é¢†å¼ æ•°"`
	ValidStartAt *gtime.Time `json:"validStartAt" v:"required#æœ‰æ•ˆæœŸå¼€å§‹æ—¶é—´不能为空" dc:"æœ‰æ•ˆæœŸå¼€å§‹æ—¶é—´"`
	ValidEndAt *gtime.Time `json:"validEndAt" v:"required#æœ‰æ•ˆæœŸç»“æŸæ—¶é—´不能为空" dc:"æœ‰æ•ˆæœŸç»“æŸæ—¶é—´"`
	Sort int `json:"sort"  dc:"æŽ’åº"`
	Status int `json:"status"  dc:"çŠ¶æ€"`
}

// CouponCreateRes 创建ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨响应
type CouponCreateRes struct {
	g.Meta `mime:"application/json"`
}

// CouponUpdateReq 更新ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨请求
type CouponUpdateReq struct {
	g.Meta `path:"/coupon/update" method:"put" tags:"ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨" summary:"更新ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨ID"`
	Title string `json:"title" dc:"ä¼˜æƒ åˆ¸åç§°"`
	Type int `json:"type" dc:"ä¼˜æƒ åˆ¸ç±»åž‹"`
	IsNewMember int `json:"isNewMember" dc:"æ˜¯å¦æ–°äººä¸“äº«"`
	FaceValue int64 `json:"faceValue" dc:"é¢å€¼ï¼ˆåˆ†ï¼‰"`
	MinAmount int64 `json:"minAmount" dc:"æœ€ä½Žæ¶ˆè´¹é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	TotalNum int `json:"totalNum" dc:"å‘æ”¾æ€»é‡ï¼ˆ0ä¸é™ï¼‰"`
	UsedNum int `json:"usedNum" dc:"å·²ä½¿ç”¨æ•°é‡"`
	ClaimNum int `json:"claimNum" dc:"å·²é¢†å–æ•°é‡"`
	PerLimit int `json:"perLimit" dc:"æ¯äººé™é¢†å¼ æ•°"`
	ValidStartAt *gtime.Time `json:"validStartAt" dc:"æœ‰æ•ˆæœŸå¼€å§‹æ—¶é—´"`
	ValidEndAt *gtime.Time `json:"validEndAt" dc:"æœ‰æ•ˆæœŸç»“æŸæ—¶é—´"`
	Sort int `json:"sort" dc:"æŽ’åº"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// CouponUpdateRes 更新ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨响应
type CouponUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// CouponDeleteReq 删除ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨请求
type CouponDeleteReq struct {
	g.Meta `path:"/coupon/delete" method:"delete" tags:"ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨" summary:"删除ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨ID"`
}

// CouponDeleteRes 删除ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨响应
type CouponDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// CouponDetailReq 获取ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨详情请求
type CouponDetailReq struct {
	g.Meta `path:"/coupon/detail" method:"get" tags:"ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨" summary:"获取ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨ID"`
}

// CouponDetailRes 获取ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨详情响应
type CouponDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.CouponDetailOutput
}

// CouponListReq 获取ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨列表请求
type CouponListReq struct {
	g.Meta   `path:"/coupon/list" method:"get" tags:"ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨" summary:"获取ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Type int `json:"type" dc:"ä¼˜æƒ åˆ¸ç±»åž‹"`
	IsNewMember int `json:"isNewMember" dc:"æ˜¯å¦æ–°äººä¸“äº«"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// CouponListRes 获取ä¼˜æƒ åˆ¸æ¨¡æ¿è¡¨列表响应
type CouponListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.CouponListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

