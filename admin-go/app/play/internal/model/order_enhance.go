package model

import "gbaseadmin/utility/snowflake"

// OrderChangeStatusInput 变更订单状态输入
type OrderChangeStatusInput struct {
	ID           snowflake.JsonInt64 `json:"id"`
	OrderStatus  int                 `json:"orderStatus"`
	CancelReason string              `json:"cancelReason"`
}
