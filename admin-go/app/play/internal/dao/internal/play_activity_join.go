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
	Id          string // 记录ID（Snowflake）
	ActivityId  string // 活动ID
	MemberId    string // 会员ID
	JoinStatus  string // 参与状态:0=已报名,1=进行中,2=已完成,3=已领奖
	CurrentStep string // 当前完成到第几步（步骤活动用）
	FinishAt    string // 完成时间
	RewardAt    string // 领奖时间
	Remark      string // 备注
	CreatedBy   string // 创建人ID
	DeptId      string // 所属部门ID
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 软删除时间
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
