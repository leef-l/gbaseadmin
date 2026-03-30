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
	Id              string // æ”¯ä»˜è®°å½•IDï¼ˆSnowflakeï¼‰
	OrderId         string // è®¢å•ID
	MemberId        string // ä¼šå‘˜ID
	PaymentNo       string // æ”¯ä»˜æµæ°´å·
	TradeNo         string // ç¬¬ä¸‰æ–¹äº¤æ˜“å·
	PayType         string // æ”¯ä»˜æ–¹å¼:1=å¾®ä¿¡æ”¯ä»˜,2=æ”¯ä»˜å®æ”¯ä»˜,3=ä½™é¢æ”¯ä»˜
	PayAmount       string // æ”¯ä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰
	PayStatus       string // æ”¯ä»˜çŠ¶æ€:0=å¾…æ”¯ä»˜,1=æ”¯ä»˜æˆåŠŸ,2=æ”¯ä»˜å¤±è´¥,3=å·²é€€æ¬¾
	PayAt           string // æ”¯ä»˜æˆåŠŸæ—¶é—´
	RefundAt        string // é€€æ¬¾æ—¶é—´
	RefundAmount    string // é€€æ¬¾é‡‘é¢ï¼ˆåˆ†ï¼‰
	CallbackContent string // å›žè°ƒæŠ¥æ–‡
	CreatedBy       string // åˆ›å»ºäººID
	DeptId          string // æ‰€å±žéƒ¨é—¨ID
	CreatedAt       string // åˆ›å»ºæ—¶é—´
	UpdatedAt       string // æ›´æ–°æ—¶é—´
	DeletedAt       string // è½¯åˆ é™¤æ—¶é—´
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
