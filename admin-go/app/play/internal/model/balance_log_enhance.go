package model

import "gbaseadmin/utility/snowflake"

// AddBalanceLogInput 余额变动统一入口输入
type AddBalanceLogInput struct {
	MemberID     snowflake.JsonInt64 `json:"memberID"`
	BizType      int                 `json:"bizType"`      // 1充值 2消费 3退款 4活动赠送 5提现
	BizID        snowflake.JsonInt64 `json:"bizID"`        // 关联业务ID（订单ID等）
	ChangeAmount int64               `json:"changeAmount"` // 变动金额（分，正数增加负数减少）
	Remark       string              `json:"remark"`
}
