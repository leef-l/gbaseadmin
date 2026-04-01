// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayWithdraw is the golang structure for table play_withdraw.
type PlayWithdraw struct {
	Id        int64       `orm:"id"         description:"提现ID"`                 // 提现ID
	CoachId   int64       `orm:"coach_id"   description:"陪玩师ID"`                // 陪玩师ID
	MemberId  int64       `orm:"member_id"  description:"会员ID"`                 // 会员ID
	Amount    int         `orm:"amount"     description:"提现金额(分)"`              // 提现金额(分)
	Status    int         `orm:"status"     description:"状态 0=待审核 1=已打款 2=已拒绝"` // 状态 0=待审核 1=已打款 2=已拒绝
	Reason    string      `orm:"reason"     description:"拒绝原因"`                 // 拒绝原因
	AuditedAt *gtime.Time `orm:"audited_at" description:"审核时间"`                 // 审核时间
	CreatedAt *gtime.Time `orm:"created_at" description:""`                     //
	UpdatedAt *gtime.Time `orm:"updated_at" description:""`                     //
	DeletedAt *gtime.Time `orm:"deleted_at" description:""`                     //
}
