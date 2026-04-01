// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayActivityStepLogDao is the data access object for the table play_activity_step_log.
type PlayActivityStepLogDao struct {
	table    string                     // table is the underlying table name of the DAO.
	group    string                     // group is the database configuration group name of the current DAO.
	columns  PlayActivityStepLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler         // handlers for customized model modification.
}

// PlayActivityStepLogColumns defines and stores column names for the table play_activity_step_log.
type PlayActivityStepLogColumns struct {
	Id          string // 记录ID（Snowflake）
	ActivityId  string // 活动ID
	StepId      string // 步骤ID
	JoinId      string // 参与记录ID
	MemberId    string // 会员ID
	StepType    string // 步骤类型：1=文字 2=链接 3=图片
	SubmitText  string // 用户提交的文字或链接
	SubmitImage string // 用户提交的图片URL
	AuditStatus string // 审核状态：0=待审核 1=通过 2=驳回
	AuditRemark string // 审核备注
	AuditBy     string // 审核人ID
	AuditAt     string // 审核时间
	CreatedBy   string // 创建人ID
	DeptId      string // 所属部门ID
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 软删除时间
}

// playActivityStepLogColumns holds the columns for the table play_activity_step_log.
var playActivityStepLogColumns = PlayActivityStepLogColumns{
	Id:          "id",
	ActivityId:  "activity_id",
	StepId:      "step_id",
	JoinId:      "join_id",
	MemberId:    "member_id",
	StepType:    "step_type",
	SubmitText:  "submit_text",
	SubmitImage: "submit_image",
	AuditStatus: "audit_status",
	AuditRemark: "audit_remark",
	AuditBy:     "audit_by",
	AuditAt:     "audit_at",
	CreatedBy:   "created_by",
	DeptId:      "dept_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewPlayActivityStepLogDao creates and returns a new DAO object for table data access.
func NewPlayActivityStepLogDao(handlers ...gdb.ModelHandler) *PlayActivityStepLogDao {
	return &PlayActivityStepLogDao{
		group:    "default",
		table:    "play_activity_step_log",
		columns:  playActivityStepLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayActivityStepLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayActivityStepLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayActivityStepLogDao) Columns() PlayActivityStepLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayActivityStepLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayActivityStepLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayActivityStepLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
