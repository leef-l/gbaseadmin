package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// ActivityStepLog API

// ActivityStepLogCreateReq 创建活动步骤提交记录请求
type ActivityStepLogCreateReq struct {
	g.Meta `path:"/activity_step_log/create" method:"post" tags:"活动步骤提交记录" summary:"创建活动步骤提交记录"`
	ActivityID snowflake.JsonInt64 `json:"activityID"  dc:"活动ID"`
	StepID snowflake.JsonInt64 `json:"stepID"  dc:"步骤ID"`
	JoinID snowflake.JsonInt64 `json:"joinID"  dc:"参与记录ID"`
	MemberID snowflake.JsonInt64 `json:"memberID"  dc:"会员ID"`
	StepType int `json:"stepType"  dc:"步骤类型"`
	SubmitText string `json:"submitText"  dc:"用户提交的文字或链接"`
	SubmitImage string `json:"submitImage"  dc:"用户提交的图片URL"`
	AuditStatus int `json:"auditStatus"  dc:"审核状态"`
	AuditRemark string `json:"auditRemark"  dc:"审核备注"`
	AuditBy int64 `json:"auditBy"  dc:"审核人ID"`
	AuditAt *gtime.Time `json:"auditAt"  dc:"审核时间"`
}

// ActivityStepLogCreateRes 创建活动步骤提交记录响应
type ActivityStepLogCreateRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityStepLogUpdateReq 更新活动步骤提交记录请求
type ActivityStepLogUpdateReq struct {
	g.Meta `path:"/activity_step_log/update" method:"put" tags:"活动步骤提交记录" summary:"更新活动步骤提交记录"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"活动步骤提交记录ID"`
	ActivityID snowflake.JsonInt64 `json:"activityID" dc:"活动ID"`
	StepID snowflake.JsonInt64 `json:"stepID" dc:"步骤ID"`
	JoinID snowflake.JsonInt64 `json:"joinID" dc:"参与记录ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"会员ID"`
	StepType int `json:"stepType" dc:"步骤类型"`
	SubmitText string `json:"submitText" dc:"用户提交的文字或链接"`
	SubmitImage string `json:"submitImage" dc:"用户提交的图片URL"`
	AuditStatus int `json:"auditStatus" dc:"审核状态"`
	AuditRemark string `json:"auditRemark" dc:"审核备注"`
	AuditBy int64 `json:"auditBy" dc:"审核人ID"`
	AuditAt *gtime.Time `json:"auditAt" dc:"审核时间"`
}

// ActivityStepLogUpdateRes 更新活动步骤提交记录响应
type ActivityStepLogUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityStepLogDeleteReq 删除活动步骤提交记录请求
type ActivityStepLogDeleteReq struct {
	g.Meta `path:"/activity_step_log/delete" method:"delete" tags:"活动步骤提交记录" summary:"删除活动步骤提交记录"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"活动步骤提交记录ID"`
}

// ActivityStepLogDeleteRes 删除活动步骤提交记录响应
type ActivityStepLogDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ActivityStepLogDetailReq 获取活动步骤提交记录详情请求
type ActivityStepLogDetailReq struct {
	g.Meta `path:"/activity_step_log/detail" method:"get" tags:"活动步骤提交记录" summary:"获取活动步骤提交记录详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"活动步骤提交记录ID"`
}

// ActivityStepLogDetailRes 获取活动步骤提交记录详情响应
type ActivityStepLogDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.ActivityStepLogDetailOutput
}

// ActivityStepLogListReq 获取活动步骤提交记录列表请求
type ActivityStepLogListReq struct {
	g.Meta   `path:"/activity_step_log/list" method:"get" tags:"活动步骤提交记录" summary:"获取活动步骤提交记录列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	StepType int `json:"stepType" dc:"步骤类型"`
	AuditStatus int `json:"auditStatus" dc:"审核状态"`
}

// ActivityStepLogListRes 获取活动步骤提交记录列表响应
type ActivityStepLogListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.ActivityStepLogListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

