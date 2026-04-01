package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// Withdraw API

// WithdrawCreateReq 创建陪玩师提现记录请求
type WithdrawCreateReq struct {
	g.Meta `path:"/withdraw/create" method:"post" tags:"陪玩师提现记录" summary:"创建陪玩师提现记录"`
	CoachID snowflake.JsonInt64 `json:"coachID" v:"required#陪玩师ID不能为空" dc:"陪玩师ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#会员ID不能为空" dc:"会员ID"`
	Amount int `json:"amount"  dc:"提现金额(分)"`
	Status int `json:"status"  dc:"状态 0=待审核 1=已打款 2=已拒绝"`
	Reason string `json:"reason"  dc:"拒绝原因"`
	AuditedAt *gtime.Time `json:"auditedAt"  dc:"审核时间"`
}

// WithdrawCreateRes 创建陪玩师提现记录响应
type WithdrawCreateRes struct {
	g.Meta `mime:"application/json"`
}

// WithdrawUpdateReq 更新陪玩师提现记录请求
type WithdrawUpdateReq struct {
	g.Meta `path:"/withdraw/update" method:"put" tags:"陪玩师提现记录" summary:"更新陪玩师提现记录"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"陪玩师提现记录ID"`
	CoachID snowflake.JsonInt64 `json:"coachID" dc:"陪玩师ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"会员ID"`
	Amount int `json:"amount" dc:"提现金额(分)"`
	Status int `json:"status" dc:"状态 0=待审核 1=已打款 2=已拒绝"`
	Reason string `json:"reason" dc:"拒绝原因"`
	AuditedAt *gtime.Time `json:"auditedAt" dc:"审核时间"`
}

// WithdrawUpdateRes 更新陪玩师提现记录响应
type WithdrawUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// WithdrawDeleteReq 删除陪玩师提现记录请求
type WithdrawDeleteReq struct {
	g.Meta `path:"/withdraw/delete" method:"delete" tags:"陪玩师提现记录" summary:"删除陪玩师提现记录"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"陪玩师提现记录ID"`
}

// WithdrawDeleteRes 删除陪玩师提现记录响应
type WithdrawDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// WithdrawBatchDeleteReq 批量删除陪玩师提现记录请求
type WithdrawBatchDeleteReq struct {
	g.Meta `path:"/withdraw/batch-delete" method:"delete" tags:"陪玩师提现记录" summary:"批量删除陪玩师提现记录"`
	IDs    []snowflake.JsonInt64 `json:"ids" v:"required#ID列表不能为空" dc:"陪玩师提现记录ID列表"`
}

// WithdrawBatchDeleteRes 批量删除陪玩师提现记录响应
type WithdrawBatchDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// WithdrawDetailReq 获取陪玩师提现记录详情请求
type WithdrawDetailReq struct {
	g.Meta `path:"/withdraw/detail" method:"get" tags:"陪玩师提现记录" summary:"获取陪玩师提现记录详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"陪玩师提现记录ID"`
}

// WithdrawDetailRes 获取陪玩师提现记录详情响应
type WithdrawDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.WithdrawDetailOutput
}

// WithdrawListReq 获取陪玩师提现记录列表请求
type WithdrawListReq struct {
	g.Meta    `path:"/withdraw/list" method:"get" tags:"陪玩师提现记录" summary:"获取陪玩师提现记录列表"`
	PageNum   int    `json:"pageNum" d:"1" dc:"页码"`
	PageSize  int    `json:"pageSize" d:"10" dc:"每页数量"`
	OrderBy   string `json:"orderBy" dc:"排序字段"`
	OrderDir  string `json:"orderDir" d:"asc" dc:"排序方向:asc/desc"`
	StartTime string `json:"startTime" dc:"开始时间"`
	EndTime   string `json:"endTime" dc:"结束时间"`
}

// WithdrawListRes 获取陪玩师提现记录列表响应
type WithdrawListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.WithdrawListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}
// WithdrawExportReq 导出陪玩师提现记录请求
type WithdrawExportReq struct {
	g.Meta    `path:"/withdraw/export" method:"get" tags:"陪玩师提现记录" summary:"导出陪玩师提现记录"`
	StartTime string `json:"startTime" dc:"开始时间"`
	EndTime   string `json:"endTime" dc:"结束时间"`
}

// WithdrawExportRes 导出陪玩师提现记录响应
type WithdrawExportRes struct {
	g.Meta `mime:"text/csv"`
}


