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
	Id             string // 流水ID（Snowflake）
	OrderId        string // 订单ID
	OrderNo        string // 订单编号
	PayAmount      string // 实付金额（分）
	CoachId        string // 陪玩师ID
	ShopId         string // 店铺ID（0表示无店铺）
	PlatformRate   string // 平台抽成比例（百分比）
	PlatformAmount string // 平台抽成金额（分）
	ShopRate       string // 店铺抽成比例（百分比）
	ShopAmount     string // 店铺抽成金额（分）
	CoachAmount    string // 陪玩师收入（分）
	SettleStatus   string // 结算状态:0=待结算,1=已结算
	SettleAt       string // 结算时间
	CreatedBy      string // 创建人ID
	DeptId         string // 所属部门ID
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	DeletedAt      string // 软删除时间
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
