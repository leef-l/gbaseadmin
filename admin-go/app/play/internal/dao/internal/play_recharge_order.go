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
	Id             string // å……å€¼è®¢å•IDï¼ˆSnowflakeï¼‰
	OrderNo        string // å……å€¼è®¢å•å·
	MemberId       string // ä¼šå‘˜ID
	RechargePlanId string // å……å€¼æ–¹æ¡ˆID
	Amount         string // å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰
	GiftAmount     string // èµ é€é‡‘é¢ï¼ˆåˆ†ï¼‰
	PayType        string // æ”¯ä»˜æ–¹å¼:1=å¾®ä¿¡æ”¯ä»˜,2=æ”¯ä»˜å®æ”¯ä»˜
	TradeNo        string // ç¬¬ä¸‰æ–¹äº¤æ˜“å·
	PayStatus      string // æ”¯ä»˜çŠ¶æ€:0=å¾…æ”¯ä»˜,1=æ”¯ä»˜æˆåŠŸ,2=æ”¯ä»˜å¤±è´¥
	PayAt          string // æ”¯ä»˜æ—¶é—´
	CreatedBy      string // åˆ›å»ºäººID
	DeptId         string // æ‰€å±žéƒ¨é—¨ID
	CreatedAt      string // åˆ›å»ºæ—¶é—´
	UpdatedAt      string // æ›´æ–°æ—¶é—´
	DeletedAt      string // è½¯åˆ é™¤æ—¶é—´
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
