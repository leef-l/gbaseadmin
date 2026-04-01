// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayWithdraw is the golang structure of table play_withdraw for DAO operations like Where/Data.
type PlayWithdraw struct {
	g.Meta    `orm:"table:play_withdraw, do:true"`
	Id        any         // 提现ID
	CoachId   any         // 陪玩师ID
	MemberId  any         // 会员ID
	Amount    any         // 提现金额(分)
	Status    any         // 状态 0=待审核 1=已打款 2=已拒绝
	Reason    any         // 拒绝原因
	AuditedAt *gtime.Time // 审核时间
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
