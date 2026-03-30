package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gbaseadmin/utility/snowflake"
)

// CoachApplyAuditReq 审核陪玩师申请请求
type CoachApplyAuditReq struct {
	g.Meta      `path:"/coach_apply/audit" method:"post" tags:"陪玩师申请表" summary:"审核陪玩师申请"`
	ID          snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"申请ID"`
	AuditStatus int                 `json:"auditStatus" v:"required|in:1,2#审核状态不能为空|审核状态只能是1(通过)或2(拒绝)" dc:"审核状态:1通过2拒绝"`
	AuditRemark string              `json:"auditRemark" dc:"审核备注"`
}

// CoachApplyAuditRes 审核陪玩师申请响应
type CoachApplyAuditRes struct {
	g.Meta `mime:"application/json"`
}
