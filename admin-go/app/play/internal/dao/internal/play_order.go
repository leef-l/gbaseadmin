// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayOrderDao is the data access object for the table play_order.
type PlayOrderDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PlayOrderColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PlayOrderColumns defines and stores column names for the table play_order.
type PlayOrderColumns struct {
	Id             string // 订单ID（Snowflake）
	OrderNo        string // 订单编号
	MemberId       string // 下单会员ID
	CoachId        string // 陪玩师ID
	ShopId         string // 店铺ID（0表示无店铺）
	GoodsId        string // 商品ID
	GoodsTitle     string // 商品名称（冗余）
	GoodsPrice     string // 商品单价（分，下单时快照）
	Quantity       string // 数量
	TotalAmount    string // 订单总额（分）
	DiscountAmount string // 会员折扣金额（分）
	CouponAmount   string // 优惠券抵扣金额（分）
	PayAmount      string // 实付金额（分）
	CouponMemberId string // 使用的优惠券领取记录ID
	PayType        string // 支付方式:0=未支付,1=微信支付,2=支付宝支付,3=余额支付
	OrderStatus    string // 订单状态:0=待支付,1=已支付,2=进行中,3=已完成,4=已取消,5=退款中,6=已退款
	PayAt          string // 支付时间
	StartAt        string // 服务开始时间
	FinishAt       string // 服务完成时间
	CancelAt       string // 取消时间
	CancelReason   string // 取消原因
	Remark         string // 订单备注
	CreatedBy      string // 创建人ID
	DeptId         string // 所属部门ID
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	DeletedAt      string // 软删除时间
}

// playOrderColumns holds the columns for the table play_order.
var playOrderColumns = PlayOrderColumns{
	Id:             "id",
	OrderNo:        "order_no",
	MemberId:       "member_id",
	CoachId:        "coach_id",
	ShopId:         "shop_id",
	GoodsId:        "goods_id",
	GoodsTitle:     "goods_title",
	GoodsPrice:     "goods_price",
	Quantity:       "quantity",
	TotalAmount:    "total_amount",
	DiscountAmount: "discount_amount",
	CouponAmount:   "coupon_amount",
	PayAmount:      "pay_amount",
	CouponMemberId: "coupon_member_id",
	PayType:        "pay_type",
	OrderStatus:    "order_status",
	PayAt:          "pay_at",
	StartAt:        "start_at",
	FinishAt:       "finish_at",
	CancelAt:       "cancel_at",
	CancelReason:   "cancel_reason",
	Remark:         "remark",
	CreatedBy:      "created_by",
	DeptId:         "dept_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewPlayOrderDao creates and returns a new DAO object for table data access.
func NewPlayOrderDao(handlers ...gdb.ModelHandler) *PlayOrderDao {
	return &PlayOrderDao{
		group:    "default",
		table:    "play_order",
		columns:  playOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayOrderDao) Columns() PlayOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayOrderDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PlayOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
