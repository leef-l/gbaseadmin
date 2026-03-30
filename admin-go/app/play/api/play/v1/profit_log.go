package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// ProfitLog API

// ProfitLogCreateReq 创建åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨请求
type ProfitLogCreateReq struct {
	g.Meta `path:"/profit_log/create" method:"post" tags:"åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨" summary:"创建åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨"`
	OrderID snowflake.JsonInt64 `json:"orderID" v:"required#è®¢å•ID不能为空" dc:"è®¢å•ID"`
	OrderNo string `json:"orderNo" v:"required#è®¢å•ç¼–å·不能为空" dc:"è®¢å•ç¼–å·"`
	PayAmount int64 `json:"payAmount"  dc:"å®žä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	CoachID snowflake.JsonInt64 `json:"coachID" v:"required#é™ªçŽ©å¸ˆID不能为空" dc:"é™ªçŽ©å¸ˆID"`
	ShopID snowflake.JsonInt64 `json:"shopID"  dc:"åº—é“ºID"`
	PlatformRate int `json:"platformRate"  dc:"å¹³å°æŠ½æˆæ¯”ä¾‹"`
	PlatformAmount int64 `json:"platformAmount"  dc:"å¹³å°æŠ½æˆé‡‘é¢ï¼ˆåˆ†ï¼‰"`
	ShopRate int `json:"shopRate"  dc:"åº—é“ºæŠ½æˆæ¯”ä¾‹"`
	ShopAmount int64 `json:"shopAmount"  dc:"åº—é“ºæŠ½æˆé‡‘é¢ï¼ˆåˆ†ï¼‰"`
	CoachAmount int64 `json:"coachAmount"  dc:"é™ªçŽ©å¸ˆæ”¶å…¥ï¼ˆåˆ†ï¼‰"`
	SettleStatus int `json:"settleStatus"  dc:"ç»“ç®—çŠ¶æ€"`
	SettleAt *gtime.Time `json:"settleAt"  dc:"ç»“ç®—æ—¶é—´"`
}

// ProfitLogCreateRes 创建åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨响应
type ProfitLogCreateRes struct {
	g.Meta `mime:"application/json"`
}

// ProfitLogUpdateReq 更新åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨请求
type ProfitLogUpdateReq struct {
	g.Meta `path:"/profit_log/update" method:"put" tags:"åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨" summary:"更新åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨ID"`
	OrderID snowflake.JsonInt64 `json:"orderID" dc:"è®¢å•ID"`
	OrderNo string `json:"orderNo" dc:"è®¢å•ç¼–å·"`
	PayAmount int64 `json:"payAmount" dc:"å®žä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	CoachID snowflake.JsonInt64 `json:"coachID" dc:"é™ªçŽ©å¸ˆID"`
	ShopID snowflake.JsonInt64 `json:"shopID" dc:"åº—é“ºID"`
	PlatformRate int `json:"platformRate" dc:"å¹³å°æŠ½æˆæ¯”ä¾‹"`
	PlatformAmount int64 `json:"platformAmount" dc:"å¹³å°æŠ½æˆé‡‘é¢ï¼ˆåˆ†ï¼‰"`
	ShopRate int `json:"shopRate" dc:"åº—é“ºæŠ½æˆæ¯”ä¾‹"`
	ShopAmount int64 `json:"shopAmount" dc:"åº—é“ºæŠ½æˆé‡‘é¢ï¼ˆåˆ†ï¼‰"`
	CoachAmount int64 `json:"coachAmount" dc:"é™ªçŽ©å¸ˆæ”¶å…¥ï¼ˆåˆ†ï¼‰"`
	SettleStatus int `json:"settleStatus" dc:"ç»“ç®—çŠ¶æ€"`
	SettleAt *gtime.Time `json:"settleAt" dc:"ç»“ç®—æ—¶é—´"`
}

// ProfitLogUpdateRes 更新åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨响应
type ProfitLogUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// ProfitLogDeleteReq 删除åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨请求
type ProfitLogDeleteReq struct {
	g.Meta `path:"/profit_log/delete" method:"delete" tags:"åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨" summary:"删除åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨ID"`
}

// ProfitLogDeleteRes 删除åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨响应
type ProfitLogDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ProfitLogDetailReq 获取åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨详情请求
type ProfitLogDetailReq struct {
	g.Meta `path:"/profit_log/detail" method:"get" tags:"åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨" summary:"获取åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨ID"`
}

// ProfitLogDetailRes 获取åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨详情响应
type ProfitLogDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.ProfitLogDetailOutput
}

// ProfitLogListReq 获取åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨列表请求
type ProfitLogListReq struct {
	g.Meta   `path:"/profit_log/list" method:"get" tags:"åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨" summary:"获取åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	SettleStatus int `json:"settleStatus" dc:"ç»“ç®—çŠ¶æ€"`
}

// ProfitLogListRes 获取åˆ©æ¶¦åˆ†æˆæµæ°´è¡¨列表响应
type ProfitLogListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.ProfitLogListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

