// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayActivityJoinDao is the data access object for the table play_activity_join.
type PlayActivityJoinDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  PlayActivityJoinColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// PlayActivityJoinColumns defines and stores column names for the table play_activity_join.
type PlayActivityJoinColumns struct {
	Id          string // è®°å½•IDï¼ˆSnowflakeï¼‰
	ActivityId  string // æ´»åŠ¨ID
	MemberId    string // ä¼šå‘˜ID
	JoinStatus  string // å‚ä¸ŽçŠ¶æ€:0=å·²æŠ¥å,1=è¿›è¡Œä¸­,2=å·²å®Œæˆ,3=å·²é¢†å¥–
	CurrentStep string // å½“å‰å®Œæˆåˆ°ç¬¬å‡ æ­¥
	FinishAt    string // å®Œæˆæ—¶é—´
	RewardAt    string // é¢†å¥–æ—¶é—´
	Remark      string // å¤‡æ³¨
	CreatedBy   string // åˆ›å»ºäººID
	DeptId      string // æ‰€å±žéƒ¨é—¨ID
	CreatedAt   string // åˆ›å»ºæ—¶é—´
	UpdatedAt   string // æ›´æ–°æ—¶é—´
	DeletedAt   string // è½¯åˆ é™¤æ—¶é—´
}

// playActivityJoinColumns holds the columns for the table play_activity_join.
var playActivityJoinColumns = PlayActivityJoinColumns{
	Id:          "id",
	ActivityId:  "activity_id",
	MemberId:    "member_id",
	JoinStatus:  "join_status",
	CurrentStep: "current_step",
	FinishAt:    "finish_at",
	RewardAt:    "reward_at",
	Remark:      "remark",
	CreatedBy:   "created_by",
	DeptId:      "dept_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewPlayActivityJoinDao creates and returns a new DAO object for table data access.
func NewPlayActivityJoinDao(handlers ...gdb.ModelHandler) *PlayActivityJoinDao {
	return &PlayActivityJoinDao{
		group:    "default",
		table:    "play_activity_join",
		columns:  playActivityJoinColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayActivityJoinDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayActivityJoinDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayActivityJoinDao) Columns() PlayActivityJoinColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayActivityJoinDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayActivityJoinDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayActivityJoinDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
