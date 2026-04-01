// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayBalanceLog is the golang structure for table play_balance_log.
type PlayBalanceLog struct {
	Id            uint64      `orm:"id"             description:"流水ID（Snowflake）"`                 // 流水ID（Snowflake）
	MemberId      uint64      `orm:"member_id"      description:"会员ID"`                            // 会员ID
	BizType       int         `orm:"biz_type"       description:"业务类型:1=充值,2=消费,3=退款,4=活动赠送,5=提现"` // 业务类型:1=充值,2=消费,3=退款,4=活动赠送,5=提现
	BizId         uint64      `orm:"biz_id"         description:"关联业务ID（订单ID/充值订单ID/活动ID）"`        // 关联业务ID（订单ID/充值订单ID/活动ID）
	ChangeAmount  int64       `orm:"change_amount"  description:"变动金额（分，正数增加负数减少）"`                // 变动金额（分，正数增加负数减少）
	BeforeBalance int64       `orm:"before_balance" description:"变动前余额（分）"`                        // 变动前余额（分）
	AfterBalance  int64       `orm:"after_balance"  description:"变动后余额（分）"`                        // 变动后余额（分）
	Remark        string      `orm:"remark"         description:"备注说明"`                            // 备注说明
	CreatedBy     uint64      `orm:"created_by"     description:"创建人ID"`                           // 创建人ID
	DeptId        uint64      `orm:"dept_id"        description:"所属部门ID"`                          // 所属部门ID
	CreatedAt     *gtime.Time `orm:"created_at"     description:"创建时间"`                            // 创建时间
	UpdatedAt     *gtime.Time `orm:"updated_at"     description:"更新时间"`                            // 更新时间
	DeletedAt     *gtime.Time `orm:"deleted_at"     description:"软删除时间"`                           // 软删除时间
}
