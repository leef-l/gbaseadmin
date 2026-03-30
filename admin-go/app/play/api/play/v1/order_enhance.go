package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gbaseadmin/utility/snowflake"
)

// OrderChangeStatusReq 变更订单状态请求
type OrderChangeStatusReq struct {
	g.Meta       `path:"/order/change-status" method:"post" tags:"订单表" summary:"变更订单状态"`
	ID           snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"订单ID"`
	OrderStatus  int                 `json:"orderStatus" v:"required|in:1,2,3,4,5,6#状态不能为空|状态值无效" dc:"目标状态"`
	CancelReason string              `json:"cancelReason" dc:"取消原因（取消时必填）"`
}

// OrderChangeStatusRes 变更订单状态响应
type OrderChangeStatusRes struct {
	g.Meta `mime:"application/json"`
}
