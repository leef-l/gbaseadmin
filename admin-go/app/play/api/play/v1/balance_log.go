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

// BalanceLogCreateReq 创建余额流水表请求
type BalanceLogCreateReq struct {
	g.Meta `path:"/balance_log/create" method:"post" tags:"余额流水表" summary:"创建余额流水表"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#会员ID不能为空" dc:"会员ID"`
	BizType int `json:"bizType" v:"required#业务类型不能为空" dc:"业务类型"`
	BizID snowflake.JsonInt64 `json:"bizID"  dc:"关联业务ID（订单ID/充值订单ID/活动ID）"`
	ChangeAmount int64 `json:"changeAmount" v:"required#变动金额（分，正数增加负数减少）不能为空" dc:"变动金额（分，正数增加负数减少）"`
	BeforeBalance int64 `json:"beforeBalance" v:"required#变动前余额（分）不能为空" dc:"变动前余额（分）"`
	AfterBalance int64 `json:"afterBalance" v:"required#变动后余额（分）不能为空" dc:"变动后余额（分）"`
	Remark string `json:"remark"  dc:"备注说明"`
}

// BalanceLogCreateRes 创建余额流水表响应
type BalanceLogCreateRes struct {
	g.Meta `mime:"application/json"`
}

// BalanceLogUpdateReq 更新余额流水表请求
type BalanceLogUpdateReq struct {
	g.Meta `path:"/balance_log/update" method:"put" tags:"余额流水表" summary:"更新余额流水表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"余额流水表ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"会员ID"`
	BizType int `json:"bizType" dc:"业务类型"`
	BizID snowflake.JsonInt64 `json:"bizID" dc:"关联业务ID（订单ID/充值订单ID/活动ID）"`
	ChangeAmount int64 `json:"changeAmount" dc:"变动金额（分，正数增加负数减少）"`
	BeforeBalance int64 `json:"beforeBalance" dc:"变动前余额（分）"`
	AfterBalance int64 `json:"afterBalance" dc:"变动后余额（分）"`
	Remark string `json:"remark" dc:"备注说明"`
}

// BalanceLogUpdateRes 更新余额流水表响应
type BalanceLogUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// BalanceLogDeleteReq 删除余额流水表请求
type BalanceLogDeleteReq struct {
	g.Meta `path:"/balance_log/delete" method:"delete" tags:"余额流水表" summary:"删除余额流水表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"余额流水表ID"`
}

// BalanceLogDeleteRes 删除余额流水表响应
type BalanceLogDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// BalanceLogDetailReq 获取余额流水表详情请求
type BalanceLogDetailReq struct {
	g.Meta `path:"/balance_log/detail" method:"get" tags:"余额流水表" summary:"获取余额流水表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"余额流水表ID"`
}

// BalanceLogDetailRes 获取余额流水表详情响应
type BalanceLogDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.BalanceLogDetailOutput
}

// BalanceLogListReq 获取余额流水表列表请求
type BalanceLogListReq struct {
	g.Meta   `path:"/balance_log/list" method:"get" tags:"余额流水表" summary:"获取余额流水表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	BizType int `json:"bizType" dc:"业务类型"`
}

// BalanceLogListRes 获取余额流水表列表响应
type BalanceLogListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.BalanceLogListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

