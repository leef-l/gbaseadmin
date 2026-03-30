// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayRechargePlanDao is the data access object for the table play_recharge_plan.
type PlayRechargePlanDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  PlayRechargePlanColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// PlayRechargePlanColumns defines and stores column names for the table play_recharge_plan.
type PlayRechargePlanColumns struct {
	Id         string // æ–¹æ¡ˆIDï¼ˆSnowflakeï¼‰
	Title      string // æ–¹æ¡ˆåç§°
	Amount     string // å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰
	GiftAmount string // èµ é€é‡‘é¢ï¼ˆåˆ†ï¼‰
	CoverImage string // æ–¹æ¡ˆå°é¢å›¾
	Sort       string // æŽ’åºï¼ˆå‡åºï¼‰
	Status     string // çŠ¶æ€:0=å…³é—­,1=å¼€å¯
	CreatedBy  string // åˆ›å»ºäººID
	DeptId     string // æ‰€å±žéƒ¨é—¨ID
	CreatedAt  string // åˆ›å»ºæ—¶é—´
	UpdatedAt  string // æ›´æ–°æ—¶é—´
	DeletedAt  string // è½¯åˆ é™¤æ—¶é—´
}

// playRechargePlanColumns holds the columns for the table play_recharge_plan.
var playRechargePlanColumns = PlayRechargePlanColumns{
	Id:         "id",
	Title:      "title",
	Amount:     "amount",
	GiftAmount: "gift_amount",
	CoverImage: "cover_image",
	Sort:       "sort",
	Status:     "status",
	CreatedBy:  "created_by",
	DeptId:     "dept_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewPlayRechargePlanDao creates and returns a new DAO object for table data access.
func NewPlayRechargePlanDao(handlers ...gdb.ModelHandler) *PlayRechargePlanDao {
	return &PlayRechargePlanDao{
		group:    "default",
		table:    "play_recharge_plan",
		columns:  playRechargePlanColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayRechargePlanDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayRechargePlanDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayRechargePlanDao) Columns() PlayRechargePlanColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayRechargePlanDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayRechargePlanDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayRechargePlanDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
