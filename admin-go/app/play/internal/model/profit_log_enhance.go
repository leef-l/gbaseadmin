package model

import "gbaseadmin/utility/snowflake"

// SettleOrderInput 订单利润结算入参
type SettleOrderInput struct {
	OrderID snowflake.JsonInt64 `json:"orderID"`
}
