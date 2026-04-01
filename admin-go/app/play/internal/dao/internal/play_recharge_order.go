// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayRechargeOrderDao is the data access object for the table play_recharge_order.
type PlayRechargeOrderDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  PlayRechargeOrderColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// PlayRechargeOrderColumns defines and stores column names for the table play_recharge_order.
type PlayRechargeOrderColumns struct {
	Id             string // 充值订单ID（Snowflake）
	OrderNo        string // 充值订单号
	MemberId       string // 会员ID
	RechargePlanId string // 充值方案ID
	Amount         string // 充值金额（分）
	GiftAmount     string // 赠送金额（分）
	PayType        string // 支付方式:1=微信支付,2=支付宝支付
	TradeNo        string // 第三方交易号
	PayStatus      string // 支付状态:0=待支付,1=支付成功,2=支付失败
	PayAt          string // 支付时间
	CreatedBy      string // 创建人ID
	DeptId         string // 所属部门ID
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	DeletedAt      string // 软删除时间
}

// playRechargeOrderColumns holds the columns for the table play_recharge_order.
var playRechargeOrderColumns = PlayRechargeOrderColumns{
	Id:             "id",
	OrderNo:        "order_no",
	MemberId:       "member_id",
	RechargePlanId: "recharge_plan_id",
	Amount:         "amount",
	GiftAmount:     "gift_amount",
	PayType:        "pay_type",
	TradeNo:        "trade_no",
	PayStatus:      "pay_status",
	PayAt:          "pay_at",
	CreatedBy:      "created_by",
	DeptId:         "dept_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewPlayRechargeOrderDao creates and returns a new DAO object for table data access.
func NewPlayRechargeOrderDao(handlers ...gdb.ModelHandler) *PlayRechargeOrderDao {
	return &PlayRechargeOrderDao{
		group:    "default",
		table:    "play_recharge_order",
		columns:  playRechargeOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayRechargeOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayRechargeOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayRechargeOrderDao) Columns() PlayRechargeOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayRechargeOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayRechargeOrderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayRechargeOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
