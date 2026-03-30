package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/utility/snowflake"
)

// CoachApply API

// CoachApplyCreateReq 创建陪玩师申请表请求
type CoachApplyCreateReq struct {
	g.Meta `path:"/coach_apply/create" method:"post" tags:"陪玩师申请表" summary:"创建陪玩师申请表"`
	MemberID snowflake.JsonInt64 `json:"memberID" v:"required#会员ID不能为空" dc:"会员ID"`
	RealName string `json:"realName" v:"required#真实姓名不能为空" dc:"真实姓名"`
	IDCard string `json:"idCard" v:"required#身份证号不能为空" dc:"身份证号"`
	IDCardFrontImage string `json:"idCardFrontImage" v:"required#身份证正面照不能为空" dc:"身份证正面照"`
	IDCardBackImage string `json:"idCardBackImage" v:"required#身份证反面照不能为空" dc:"身份证反面照"`
	SkillDesc string `json:"skillDesc"  dc:"技能描述"`
	AuditStatus int `json:"auditStatus"  dc:"审核状态"`
	AuditRemark string `json:"auditRemark"  dc:"审核备注"`
	AuditAt *gtime.Time `json:"auditAt"  dc:"审核时间"`
}

// CoachApplyCreateRes 创建陪玩师申请表响应
type CoachApplyCreateRes struct {
	g.Meta `mime:"application/json"`
}

// CoachApplyUpdateReq 更新陪玩师申请表请求
type CoachApplyUpdateReq struct {
	g.Meta `path:"/coach_apply/update" method:"put" tags:"陪玩师申请表" summary:"更新陪玩师申请表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"陪玩师申请表ID"`
	MemberID snowflake.JsonInt64 `json:"memberID" dc:"会员ID"`
	RealName string `json:"realName" dc:"真实姓名"`
	IDCard string `json:"idCard" dc:"身份证号"`
	IDCardFrontImage string `json:"idCardFrontImage" dc:"身份证正面照"`
	IDCardBackImage string `json:"idCardBackImage" dc:"身份证反面照"`
	SkillDesc string `json:"skillDesc" dc:"技能描述"`
	AuditStatus int `json:"auditStatus" dc:"审核状态"`
	AuditRemark string `json:"auditRemark" dc:"审核备注"`
	AuditAt *gtime.Time `json:"auditAt" dc:"审核时间"`
}

// CoachApplyUpdateRes 更新陪玩师申请表响应
type CoachApplyUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// CoachApplyDeleteReq 删除陪玩师申请表请求
type CoachApplyDeleteReq struct {
	g.Meta `path:"/coach_apply/delete" method:"delete" tags:"陪玩师申请表" summary:"删除陪玩师申请表"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"陪玩师申请表ID"`
}

// CoachApplyDeleteRes 删除陪玩师申请表响应
type CoachApplyDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// CoachApplyDetailReq 获取陪玩师申请表详情请求
type CoachApplyDetailReq struct {
	g.Meta `path:"/coach_apply/detail" method:"get" tags:"陪玩师申请表" summary:"获取陪玩师申请表详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"陪玩师申请表ID"`
}

// CoachApplyDetailRes 获取陪玩师申请表详情响应
type CoachApplyDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.CoachApplyDetailOutput
}

// CoachApplyListReq 获取陪玩师申请表列表请求
type CoachApplyListReq struct {
	g.Meta   `path:"/coach_apply/list" method:"get" tags:"陪玩师申请表" summary:"获取陪玩师申请表列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	AuditStatus int `json:"auditStatus" dc:"审核状态"`
}

// CoachApplyListRes 获取陪玩师申请表列表响应
type CoachApplyListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.CoachApplyListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

