package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// RechargeOrder API

// RechargeOrderCreateReq 创建å……å€¼è®¢å•è¡¨请求
type RechargeOrderCreateReq struct {
	g.Meta `path:"/recharge_order/create" method:"post" tags:"å……å€¼è®¢å•è¡¨" summary:"创建å……å€¼è®¢å•è¡¨"`
	OrderNo string `json:"orderNo" v:"required#å……å€¼è®¢å•å·不能为空" dc:"å……å€¼è®¢å•å·"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#ä¼šå‘˜ID不能为空" dc:"ä¼šå‘˜ID"`
	RechargePlanID snowflake.JsonInt64 `json:"rechargePlanID" v:"required#å……å€¼æ–¹æ¡ˆID不能为空" dc:"å……å€¼æ–¹æ¡ˆID"`
	Amount int64 `json:"amount" v:"required#å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰不能为空" dc:"å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	GiftAmount int64 `json:"giftAmount"  dc:"èµ é€é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	PayType int `json:"payType"  dc:"æ”¯ä»˜æ–¹å¼"`
	TradeNo string `json:"tradeNo"  dc:"ç¬¬ä¸‰æ–¹äº¤æ˜“å·"`
	PayStatus int `json:"payStatus"  dc:"æ”¯ä»˜çŠ¶æ€"`
	PayAt *gtime.Time `json:"payAt"  dc:"æ”¯ä»˜æ—¶é—´"`
}

// RechargeOrderCreateRes 创建å……å€¼è®¢å•è¡¨响应
type RechargeOrderCreateRes struct {
	g.Meta `mime:"application/json"`
}

// RechargeOrderUpdateReq 更新å……å€¼è®¢å•è¡¨请求
type RechargeOrderUpdateReq struct {
	g.Meta `path:"/recharge_order/update" method:"put" tags:"å……å€¼è®¢å•è¡¨" summary:"更新å……å€¼è®¢å•è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"å……å€¼è®¢å•è¡¨ID"`
	OrderNo string `json:"orderNo" dc:"å……å€¼è®¢å•å·"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"ä¼šå‘˜ID"`
	RechargePlanID snowflake.JsonInt64 `json:"rechargePlanID" dc:"å……å€¼æ–¹æ¡ˆID"`
	Amount int64 `json:"amount" dc:"å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	GiftAmount int64 `json:"giftAmount" dc:"èµ é€é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	PayType int `json:"payType" dc:"æ”¯ä»˜æ–¹å¼"`
	TradeNo string `json:"tradeNo" dc:"ç¬¬ä¸‰æ–¹äº¤æ˜“å·"`
	PayStatus int `json:"payStatus" dc:"æ”¯ä»˜çŠ¶æ€"`
	PayAt *gtime.Time `json:"payAt" dc:"æ”¯ä»˜æ—¶é—´"`
}

// RechargeOrderUpdateRes 更新å……å€¼è®¢å•è¡¨响应
type RechargeOrderUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// RechargeOrderDeleteReq 删除å……å€¼è®¢å•è¡¨请求
type RechargeOrderDeleteReq struct {
	g.Meta `path:"/recharge_order/delete" method:"delete" tags:"å……å€¼è®¢å•è¡¨" summary:"删除å……å€¼è®¢å•è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"å……å€¼è®¢å•è¡¨ID"`
}

// RechargeOrderDeleteRes 删除å……å€¼è®¢å•è¡¨响应
type RechargeOrderDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// RechargeOrderDetailReq 获取å……å€¼è®¢å•è¡¨详情请求
type RechargeOrderDetailReq struct {
	g.Meta `path:"/recharge_order/detail" method:"get" tags:"å……å€¼è®¢å•è¡¨" summary:"获取å……å€¼è®¢å•è¡¨详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"å……å€¼è®¢å•è¡¨ID"`
}

// RechargeOrderDetailRes 获取å……å€¼è®¢å•è¡¨详情响应
type RechargeOrderDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.RechargeOrderDetailOutput
}

// RechargeOrderListReq 获取å……å€¼è®¢å•è¡¨列表请求
type RechargeOrderListReq struct {
	g.Meta   `path:"/recharge_order/list" method:"get" tags:"å……å€¼è®¢å•è¡¨" summary:"获取å……å€¼è®¢å•è¡¨列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	PayType int `json:"payType" dc:"æ”¯ä»˜æ–¹å¼"`
	PayStatus int `json:"payStatus" dc:"æ”¯ä»˜çŠ¶æ€"`
}

// RechargeOrderListRes 获取å……å€¼è®¢å•è¡¨列表响应
type RechargeOrderListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.RechargeOrderListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

