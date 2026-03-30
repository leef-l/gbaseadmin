package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// ProfitLog API

// ProfitLogCreateReq 创建利润分成流水表请求
type ProfitLogCreateReq struct {
	g.Meta `path:"/profit_log/create" method:"post" tags:"利润分成流水表" summary:"创建利润分成流水表"`
	OrderID snowflake.JsonInt64 `json:"orderID" v:"required#订单ID不能为空" dc:"订单ID"`
	OrderNo string `json:"orderNo" v:"required#订单编号不能为空" dc:"订单编号"`
	PayAmount int64 `json:"payAmount"  dc:"实付金额（分）"`
	CoachID snowflake.JsonInt64 `json:"coachID" v:"required#陪玩师ID不能为空" dc:"陪玩师ID"`
	ShopID snowflake.JsonInt64 `json:"shopID"  dc:"店铺ID（0表示无店铺）"`
	PlatformRate int `json:"platformRate"  dc:"平台抽成比例（百分比）"`
	PlatformAmount int64 `json:"platformAmount"  dc:"平台抽成金额（分）"`
	ShopRate int `json:"shopRate"  dc:"店铺抽成比例（百分比）"`
	ShopAmount int64 `json:"shopAmount"  dc:"店铺抽成金额（分）"`
	CoachAmount int64 `json:"coachAmount"  dc:"陪玩师收入（分）"`
	SettleStatus int `json:"settleStatus"  dc:"结算状态"`
	SettleAt *gtime.Time `json:"settleAt"  dc:"结算时间"`
}

// ProfitLogCreateRes 创建利润分成流水表响应
type ProfitLogCreateRes struct {
	g.Meta `mime:"application/json"`
}

// ProfitLogUpdateReq 更新利润分成流水表请求
type ProfitLogUpdateReq struct {
	g.Meta `path:"/profit_log/update" method:"put" tags:"利润分成流水表" summary:"更新利润分成流水表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"利润分成流水表ID"`
	OrderID snowflake.JsonInt64 `json:"orderID" dc:"订单ID"`
	OrderNo string `json:"orderNo" dc:"订单编号"`
	PayAmount int64 `json:"payAmount" dc:"实付金额（分）"`
	CoachID snowflake.JsonInt64 `json:"coachID" dc:"陪玩师ID"`
	ShopID snowflake.JsonInt64 `json:"shopID" dc:"店铺ID（0表示无店铺）"`
	PlatformRate int `json:"platformRate" dc:"平台抽成比例（百分比）"`
	PlatformAmount int64 `json:"platformAmount" dc:"平台抽成金额（分）"`
	ShopRate int `json:"shopRate" dc:"店铺抽成比例（百分比）"`
	ShopAmount int64 `json:"shopAmount" dc:"店铺抽成金额（分）"`
	CoachAmount int64 `json:"coachAmount" dc:"陪玩师收入（分）"`
	SettleStatus int `json:"settleStatus" dc:"结算状态"`
	SettleAt *gtime.Time `json:"settleAt" dc:"结算时间"`
}

// ProfitLogUpdateRes 更新利润分成流水表响应
type ProfitLogUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// ProfitLogDeleteReq 删除利润分成流水表请求
type ProfitLogDeleteReq struct {
	g.Meta `path:"/profit_log/delete" method:"delete" tags:"利润分成流水表" summary:"删除利润分成流水表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"利润分成流水表ID"`
}

// ProfitLogDeleteRes 删除利润分成流水表响应
type ProfitLogDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ProfitLogDetailReq 获取利润分成流水表详情请求
type ProfitLogDetailReq struct {
	g.Meta `path:"/profit_log/detail" method:"get" tags:"利润分成流水表" summary:"获取利润分成流水表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"利润分成流水表ID"`
}

// ProfitLogDetailRes 获取利润分成流水表详情响应
type ProfitLogDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.ProfitLogDetailOutput
}

// ProfitLogListReq 获取利润分成流水表列表请求
type ProfitLogListReq struct {
	g.Meta   `path:"/profit_log/list" method:"get" tags:"利润分成流水表" summary:"获取利润分成流水表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	SettleStatus int `json:"settleStatus" dc:"结算状态"`
}

// ProfitLogListRes 获取利润分成流水表列表响应
type ProfitLogListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.ProfitLogListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

