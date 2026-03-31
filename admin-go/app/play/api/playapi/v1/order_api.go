package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ========== MemberAuth 接口 ==========

type OrderCreateReq struct {
	g.Meta         `path:"/order/create" method:"post" tags:"C端订单" summary:"创建订单"`
	GoodsID        string `json:"goodsId" v:"required#商品ID不能为空" dc:"商品ID"`
	Quantity       int    `json:"quantity" v:"required|min:1|max:99#数量不能为空|数量最少1|数量最多99" dc:"购买数量"`
	CouponMemberID string `json:"couponMemberId" dc:"使用的优惠券ID（play_coupon_member.id）"`
	Remark         string `json:"remark" v:"max-length:200#备注最多200字" dc:"订单备注"`
}

type OrderCreateRes struct {
	g.Meta    `mime:"application/json"`
	OrderID   string `json:"orderId" dc:"订单ID"`
	OrderNo   string `json:"orderNo" dc:"订单编号"`
	PayAmount int64  `json:"payAmount" dc:"应付金额(分)"`
}

type OrderListReq struct {
	g.Meta   `path:"/order/list" method:"get" tags:"C端订单" summary:"我的订单列表"`
	Status   *int `json:"status" dc:"状态筛选:0=待支付,1=已支付,2=进行中,3=已完成,4=已取消,5=退款中,6=已退款"`
	Page     int  `json:"page" v:"min:1" dc:"页码" d:"1"`
	PageSize int  `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}

type OrderListItem struct {
	OrderID     string `json:"orderId" dc:"订单ID"`
	OrderNo     string `json:"orderNo" dc:"订单编号"`
	GoodsTitle  string `json:"goodsTitle" dc:"商品标题"`
	GoodsImage  string `json:"goodsImage" dc:"商品图片"`
	CoachID     string `json:"coachId" dc:"陪玩师ID"`
	CoachName   string `json:"coachName" dc:"陪玩师昵称"`
	CoachAvatar string `json:"coachAvatar" dc:"陪玩师头像"`
	Quantity    int    `json:"quantity" dc:"数量"`
	TotalAmount int64  `json:"totalAmount" dc:"总金额(分)"`
	PayAmount   int64  `json:"payAmount" dc:"实付金额(分)"`
	OrderStatus int    `json:"orderStatus" dc:"订单状态"`
	CreatedAt   string `json:"createdAt" dc:"下单时间"`
}

type OrderListRes struct {
	g.Meta `mime:"application/json"`
	Total  int             `json:"total" dc:"总数"`
	List   []OrderListItem `json:"list" dc:"订单列表"`
}

type OrderDetailReq struct {
	g.Meta  `path:"/order/detail" method:"get" tags:"C端订单" summary:"订单详情"`
	OrderID string `json:"orderId" v:"required#订单ID不能为空" dc:"订单ID"`
}

type OrderDetailRes struct {
	g.Meta         `mime:"application/json"`
	OrderID        string `json:"orderId" dc:"订单ID"`
	OrderNo        string `json:"orderNo" dc:"订单编号"`
	GoodsID        string `json:"goodsId" dc:"商品ID"`
	GoodsTitle     string `json:"goodsTitle" dc:"商品标题"`
	GoodsImage     string `json:"goodsImage" dc:"商品图片"`
	GoodsPrice     int64  `json:"goodsPrice" dc:"商品单价(分)"`
	CoachID        string `json:"coachId" dc:"陪玩师ID"`
	CoachName      string `json:"coachName" dc:"陪玩师昵称"`
	CoachAvatar    string `json:"coachAvatar" dc:"陪玩师头像"`
	Quantity       int    `json:"quantity" dc:"数量"`
	TotalAmount    int64  `json:"totalAmount" dc:"总金额(分)"`
	DiscountAmount int64  `json:"discountAmount" dc:"折扣金额(分)"`
	CouponAmount   int64  `json:"couponAmount" dc:"优惠券抵扣(分)"`
	PayAmount      int64  `json:"payAmount" dc:"实付金额(分)"`
	PayType        int    `json:"payType" dc:"支付方式"`
	OrderStatus    int    `json:"orderStatus" dc:"订单状态"`
	Remark         string `json:"remark" dc:"备注"`
	CreatedAt      string `json:"createdAt" dc:"下单时间"`
	PayAt          string `json:"payAt" dc:"支付时间"`
	StartAt        string `json:"startAt" dc:"开始时间"`
	FinishAt       string `json:"finishAt" dc:"完成时间"`
	CancelAt       string `json:"cancelAt" dc:"取消时间"`
	CancelReason   string `json:"cancelReason" dc:"取消原因"`
	HasReview      bool   `json:"hasReview" dc:"是否已评价"`
}

type OrderCancelReq struct {
	g.Meta       `path:"/order/cancel" method:"post" tags:"C端订单" summary:"取消订单"`
	OrderID      string `json:"orderId" v:"required#订单ID不能为空" dc:"订单ID"`
	CancelReason string `json:"cancelReason" v:"max-length:200#取消原因最多200字" dc:"取消原因"`
}

type OrderCancelRes struct {
	g.Meta `mime:"application/json"`
}

type OrderRefundReq struct {
	g.Meta       `path:"/order/refund" method:"post" tags:"C端订单" summary:"申请退款"`
	OrderID      string `json:"orderId" v:"required#订单ID不能为空" dc:"订单ID"`
	RefundReason string `json:"refundReason" v:"required|max-length:200#退款原因不能为空|退款原因最多200字" dc:"退款原因"`
}

type OrderRefundRes struct {
	g.Meta `mime:"application/json"`
}

// ========== CoachOnly 接口 ==========

type OrderAcceptReq struct {
	g.Meta  `path:"/order/accept" method:"post" tags:"C端订单" summary:"陪玩师接单"`
	OrderID string `json:"orderId" v:"required#订单ID不能为空" dc:"订单ID"`
}

type OrderAcceptRes struct {
	g.Meta `mime:"application/json"`
}

type OrderCompleteReq struct {
	g.Meta  `path:"/order/finish" method:"post" tags:"C端订单" summary:"陪玩师完成服务"`
	OrderID string `json:"orderId" v:"required#订单ID不能为空" dc:"订单ID"`
}

type OrderCompleteRes struct {
	g.Meta `mime:"application/json"`
}
