// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayPaymentDao is the data access object for the table play_payment.
type PlayPaymentDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PlayPaymentColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PlayPaymentColumns defines and stores column names for the table play_payment.
type PlayPaymentColumns struct {
	Id              string // 支付记录ID（Snowflake）
	OrderId         string // 订单ID
	MemberId        string // 会员ID
	PaymentNo       string // 支付流水号（平台内部）
	TradeNo         string // 第三方交易号
	PayType         string // 支付方式:1=微信支付,2=支付宝支付,3=余额支付
	PayAmount       string // 支付金额（分）
	PayStatus       string // 支付状态:0=待支付,1=支付成功,2=支付失败,3=已退款
	PayAt           string // 支付成功时间
	RefundAt        string // 退款时间
	RefundAmount    string // 退款金额（分）
	CallbackContent string // 回调报文
	CreatedBy       string // 创建人ID
	DeptId          string // 所属部门ID
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
	DeletedAt       string // 软删除时间
}

// playPaymentColumns holds the columns for the table play_payment.
var playPaymentColumns = PlayPaymentColumns{
	Id:              "id",
	OrderId:         "order_id",
	MemberId:        "member_id",
	PaymentNo:       "payment_no",
	TradeNo:         "trade_no",
	PayType:         "pay_type",
	PayAmount:       "pay_amount",
	PayStatus:       "pay_status",
	PayAt:           "pay_at",
	RefundAt:        "refund_at",
	RefundAmount:    "refund_amount",
	CallbackContent: "callback_content",
	CreatedBy:       "created_by",
	DeptId:          "dept_id",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
}

// NewPlayPaymentDao creates and returns a new DAO object for table data access.
func NewPlayPaymentDao(handlers ...gdb.ModelHandler) *PlayPaymentDao {
	return &PlayPaymentDao{
		group:    "default",
		table:    "play_payment",
		columns:  playPaymentColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayPaymentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayPaymentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayPaymentDao) Columns() PlayPaymentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayPaymentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayPaymentDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayPaymentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
