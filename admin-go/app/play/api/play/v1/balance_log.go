package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// BalanceLog API

// BalanceLogCreateReq 创建ä½™é¢æµæ°´è¡¨请求
type BalanceLogCreateReq struct {
	g.Meta `path:"/balance_log/create" method:"post" tags:"ä½™é¢æµæ°´è¡¨" summary:"创建ä½™é¢æµæ°´è¡¨"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#ä¼šå‘˜ID不能为空" dc:"ä¼šå‘˜ID"`
	BizType int `json:"bizType" v:"required#ä¸šåŠ¡ç±»åž‹不能为空" dc:"ä¸šåŠ¡ç±»åž‹"`
	BizID snowflake.JsonInt64 `json:"bizID"  dc:"å…³è”ä¸šåŠ¡ID"`
	ChangeAmount int64 `json:"changeAmount" v:"required#å˜åŠ¨é‡‘é¢ï¼ˆåˆ†ï¼‰不能为空" dc:"å˜åŠ¨é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	BeforeBalance int64 `json:"beforeBalance" v:"required#å˜åŠ¨å‰ä½™é¢ï¼ˆåˆ†ï¼‰不能为空" dc:"å˜åŠ¨å‰ä½™é¢ï¼ˆåˆ†ï¼‰"`
	AfterBalance int64 `json:"afterBalance" v:"required#å˜åŠ¨åŽä½™é¢ï¼ˆåˆ†ï¼‰不能为空" dc:"å˜åŠ¨åŽä½™é¢ï¼ˆåˆ†ï¼‰"`
	Remark string `json:"remark"  dc:"å¤‡æ³¨è¯´æ˜Ž"`
}

// BalanceLogCreateRes 创建ä½™é¢æµæ°´è¡¨响应
type BalanceLogCreateRes struct {
	g.Meta `mime:"application/json"`
}

// BalanceLogUpdateReq 更新ä½™é¢æµæ°´è¡¨请求
type BalanceLogUpdateReq struct {
	g.Meta `path:"/balance_log/update" method:"put" tags:"ä½™é¢æµæ°´è¡¨" summary:"更新ä½™é¢æµæ°´è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"ä½™é¢æµæ°´è¡¨ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"ä¼šå‘˜ID"`
	BizType int `json:"bizType" dc:"ä¸šåŠ¡ç±»åž‹"`
	BizID snowflake.JsonInt64 `json:"bizID" dc:"å…³è”ä¸šåŠ¡ID"`
	ChangeAmount int64 `json:"changeAmount" dc:"å˜åŠ¨é‡‘é¢ï¼ˆåˆ†ï¼‰"`
	BeforeBalance int64 `json:"beforeBalance" dc:"å˜åŠ¨å‰ä½™é¢ï¼ˆåˆ†ï¼‰"`
	AfterBalance int64 `json:"afterBalance" dc:"å˜åŠ¨åŽä½™é¢ï¼ˆåˆ†ï¼‰"`
	Remark string `json:"remark" dc:"å¤‡æ³¨è¯´æ˜Ž"`
}

// BalanceLogUpdateRes 更新ä½™é¢æµæ°´è¡¨响应
type BalanceLogUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// BalanceLogDeleteReq 删除ä½™é¢æµæ°´è¡¨请求
type BalanceLogDeleteReq struct {
	g.Meta `path:"/balance_log/delete" method:"delete" tags:"ä½™é¢æµæ°´è¡¨" summary:"删除ä½™é¢æµæ°´è¡¨"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"ä½™é¢æµæ°´è¡¨ID"`
}

// BalanceLogDeleteRes 删除ä½™é¢æµæ°´è¡¨响应
type BalanceLogDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// BalanceLogDetailReq 获取ä½™é¢æµæ°´è¡¨详情请求
type BalanceLogDetailReq struct {
	g.Meta `path:"/balance_log/detail" method:"get" tags:"ä½™é¢æµæ°´è¡¨" summary:"获取ä½™é¢æµæ°´è¡¨详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"ä½™é¢æµæ°´è¡¨ID"`
}

// BalanceLogDetailRes 获取ä½™é¢æµæ°´è¡¨详情响应
type BalanceLogDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.BalanceLogDetailOutput
}

// BalanceLogListReq 获取ä½™é¢æµæ°´è¡¨列表请求
type BalanceLogListReq struct {
	g.Meta   `path:"/balance_log/list" method:"get" tags:"ä½™é¢æµæ°´è¡¨" summary:"获取ä½™é¢æµæ°´è¡¨列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	BizType int `json:"bizType" dc:"ä¸šåŠ¡ç±»åž‹"`
}

// BalanceLogListRes 获取ä½™é¢æµæ°´è¡¨列表响应
type BalanceLogListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.BalanceLogListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

