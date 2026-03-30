// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayProfitLogDao is the data access object for the table play_profit_log.
type PlayProfitLogDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  PlayProfitLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// PlayProfitLogColumns defines and stores column names for the table play_profit_log.
type PlayProfitLogColumns struct {
	Id             string // æµæ°´IDï¼ˆSnowflakeï¼‰
	OrderId        string // è®¢å•ID
	OrderNo        string // è®¢å•ç¼–å·
	PayAmount      string // å®žä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰
	CoachId        string // é™ªçŽ©å¸ˆID
	ShopId         string // åº—é“ºID
	PlatformRate   string // å¹³å°æŠ½æˆæ¯”ä¾‹
	PlatformAmount string // å¹³å°æŠ½æˆé‡‘é¢ï¼ˆåˆ†ï¼‰
	ShopRate       string // åº—é“ºæŠ½æˆæ¯”ä¾‹
	ShopAmount     string // åº—é“ºæŠ½æˆé‡‘é¢ï¼ˆåˆ†ï¼‰
	CoachAmount    string // é™ªçŽ©å¸ˆæ”¶å…¥ï¼ˆåˆ†ï¼‰
	SettleStatus   string // ç»“ç®—çŠ¶æ€:0=å¾…ç»“ç®—,1=å·²ç»“ç®—
	SettleAt       string // ç»“ç®—æ—¶é—´
	CreatedBy      string // åˆ›å»ºäººID
	DeptId         string // æ‰€å±žéƒ¨é—¨ID
	CreatedAt      string // åˆ›å»ºæ—¶é—´
	UpdatedAt      string // æ›´æ–°æ—¶é—´
	DeletedAt      string // è½¯åˆ é™¤æ—¶é—´
}

// playProfitLogColumns holds the columns for the table play_profit_log.
var playProfitLogColumns = PlayProfitLogColumns{
	Id:             "id",
	OrderId:        "order_id",
	OrderNo:        "order_no",
	PayAmount:      "pay_amount",
	CoachId:        "coach_id",
	ShopId:         "shop_id",
	PlatformRate:   "platform_rate",
	PlatformAmount: "platform_amount",
	ShopRate:       "shop_rate",
	ShopAmount:     "shop_amount",
	CoachAmount:    "coach_amount",
	SettleStatus:   "settle_status",
	SettleAt:       "settle_at",
	CreatedBy:      "created_by",
	DeptId:         "dept_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewPlayProfitLogDao creates and returns a new DAO object for table data access.
func NewPlayProfitLogDao(handlers ...gdb.ModelHandler) *PlayProfitLogDao {
	return &PlayProfitLogDao{
		group:    "default",
		table:    "play_profit_log",
		columns:  playProfitLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayProfitLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayProfitLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayProfitLogDao) Columns() PlayProfitLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayProfitLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayProfitLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayProfitLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
