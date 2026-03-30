// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayActivityStepDao is the data access object for the table play_activity_step.
type PlayActivityStepDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  PlayActivityStepColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// PlayActivityStepColumns defines and stores column names for the table play_activity_step.
type PlayActivityStepColumns struct {
	Id          string // æ­¥éª¤IDï¼ˆSnowflakeï¼‰
	ActivityId  string // æ´»åŠ¨ID
	StepNum     string // æ­¥éª¤åºå·
	Title       string // æ­¥éª¤æ ‡é¢˜
	DescContent string // æ­¥éª¤è¯´æ˜Ž
	StepImage   string // æ­¥éª¤ç¤ºä¾‹å›¾ç‰‡
	Sort        string // æŽ’åº
	CreatedBy   string // åˆ›å»ºäººID
	DeptId      string // æ‰€å±žéƒ¨é—¨ID
	CreatedAt   string // åˆ›å»ºæ—¶é—´
	UpdatedAt   string // æ›´æ–°æ—¶é—´
	DeletedAt   string // è½¯åˆ é™¤æ—¶é—´
}

// playActivityStepColumns holds the columns for the table play_activity_step.
var playActivityStepColumns = PlayActivityStepColumns{
	Id:          "id",
	ActivityId:  "activity_id",
	StepNum:     "step_num",
	Title:       "title",
	DescContent: "desc_content",
	StepImage:   "step_image",
	Sort:        "sort",
	CreatedBy:   "created_by",
	DeptId:      "dept_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewPlayActivityStepDao creates and returns a new DAO object for table data access.
func NewPlayActivityStepDao(handlers ...gdb.ModelHandler) *PlayActivityStepDao {
	return &PlayActivityStepDao{
		group:    "default",
		table:    "play_activity_step",
		columns:  playActivityStepColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayActivityStepDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayActivityStepDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayActivityStepDao) Columns() PlayActivityStepColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayActivityStepDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayActivityStepDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayActivityStepDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
