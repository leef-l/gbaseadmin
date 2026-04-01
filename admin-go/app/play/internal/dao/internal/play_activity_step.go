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
	Id          string // 步骤ID（Snowflake）
	ActivityId  string // 活动ID
	StepNum     string // 步骤序号
	Title       string // 步骤标题
	StepType    string // 步骤类型：1=文字 2=链接 3=图片
	ExampleText string // 示例文字或链接URL
	DescContent string // 步骤说明（富文本，支持图文）
	StepImage   string // 步骤示例图片
	IsRequired  string // 是否需要填写:0=不需要,1=需要
	Sort        string // 排序（升序）
	CreatedBy   string // 创建人ID
	DeptId      string // 所属部门ID
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 软删除时间
}

// playActivityStepColumns holds the columns for the table play_activity_step.
var playActivityStepColumns = PlayActivityStepColumns{
	Id:          "id",
	ActivityId:  "activity_id",
	StepNum:     "step_num",
	Title:       "title",
	StepType:    "step_type",
	ExampleText: "example_text",
	DescContent: "desc_content",
	StepImage:   "step_image",
	IsRequired:  "is_required",
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
