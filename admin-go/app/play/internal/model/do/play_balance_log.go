// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayBalanceLog is the golang structure of table play_balance_log for DAO operations like Where/Data.
type PlayBalanceLog struct {
	g.Meta        `orm:"table:play_balance_log, do:true"`
	Id            any         // 流水ID（Snowflake）
	MemberId      any         // 会员ID
	BizType       any         // 业务类型:1=充值,2=消费,3=退款,4=活动赠送,5=提现
	BizId         any         // 关联业务ID（订单ID/充值订单ID/活动ID）
	ChangeAmount  any         // 变动金额（分，正数增加负数减少）
	BeforeBalance any         // 变动前余额（分）
	AfterBalance  any         // 变动后余额（分）
	Remark        any         // 备注说明
	CreatedBy     any         // 创建人ID
	DeptId        any         // 所属部门ID
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 软删除时间
}
