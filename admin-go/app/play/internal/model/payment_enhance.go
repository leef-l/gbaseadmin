package model

import "gbaseadmin/utility/snowflake"

// BalancePayInput 余额支付输入
type BalancePayInput struct {
	OrderID  snowflake.JsonInt64 `json:"orderID"`
	MemberID snowflake.JsonInt64 `json:"memberID"`
}
