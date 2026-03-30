// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayActivityRewardDao is the data access object for the table play_activity_reward.
type PlayActivityRewardDao struct {
	table    string                    // table is the underlying table name of the DAO.
	group    string                    // group is the database configuration group name of the current DAO.
	columns  PlayActivityRewardColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler        // handlers for customized model modification.
}

// PlayActivityRewardColumns defines and stores column names for the table play_activity_reward.
type PlayActivityRewardColumns struct {
	Id          string // å¥–åŠ±IDï¼ˆSnowflakeï¼‰
	ActivityId  string // æ´»åŠ¨ID
	RewardType  string // å¥–åŠ±ç±»åž‹:1=ä½™é¢,2=ä¼˜æƒ åˆ¸,3=ç»éªŒå€¼,4=ä¼šå‘˜ç­‰çº§å¤©æ•°
	RewardValue string // å¥–åŠ±æ•°å€¼
	RewardName  string // å¥–åŠ±åç§°
	Sort        string // æŽ’åº
	CreatedBy   string // åˆ›å»ºäººID
	DeptId      string // æ‰€å±žéƒ¨é—¨ID
	CreatedAt   string // åˆ›å»ºæ—¶é—´
	UpdatedAt   string // æ›´æ–°æ—¶é—´
	DeletedAt   string // è½¯åˆ é™¤æ—¶é—´
}

// playActivityRewardColumns holds the columns for the table play_activity_reward.
var playActivityRewardColumns = PlayActivityRewardColumns{
	Id:          "id",
	ActivityId:  "activity_id",
	RewardType:  "reward_type",
	RewardValue: "reward_value",
	RewardName:  "reward_name",
	Sort:        "sort",
	CreatedBy:   "created_by",
	DeptId:      "dept_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewPlayActivityRewardDao creates and returns a new DAO object for table data access.
func NewPlayActivityRewardDao(handlers ...gdb.ModelHandler) *PlayActivityRewardDao {
	return &PlayActivityRewardDao{
		group:    "default",
		table:    "play_activity_reward",
		columns:  playActivityRewardColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayActivityRewardDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayActivityRewardDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayActivityRewardDao) Columns() PlayActivityRewardColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayActivityRewardDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayActivityRewardDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayActivityRewardDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
