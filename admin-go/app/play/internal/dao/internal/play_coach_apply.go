// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayCoachApplyDao is the data access object for the table play_coach_apply.
type PlayCoachApplyDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  PlayCoachApplyColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// PlayCoachApplyColumns defines and stores column names for the table play_coach_apply.
type PlayCoachApplyColumns struct {
	Id               string // 申请ID（Snowflake）
	MemberId         string // 会员ID
	RealName         string // 真实姓名
	IdCard           string // 身份证号
	IdCardFrontImage string // 身份证正面照
	IdCardBackImage  string // 身份证反面照
	SkillDesc        string // 技能描述
	AuditStatus      string // 审核状态:0=待审核,1=通过,2=拒绝
	AuditRemark      string // 审核备注
	AuditAt          string // 审核时间
	CreatedBy        string // 创建人ID
	DeptId           string // 所属部门ID
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
	DeletedAt        string // 软删除时间
}

// playCoachApplyColumns holds the columns for the table play_coach_apply.
var playCoachApplyColumns = PlayCoachApplyColumns{
	Id:               "id",
	MemberId:         "member_id",
	RealName:         "real_name",
	IdCard:           "id_card",
	IdCardFrontImage: "id_card_front_image",
	IdCardBackImage:  "id_card_back_image",
	SkillDesc:        "skill_desc",
	AuditStatus:      "audit_status",
	AuditRemark:      "audit_remark",
	AuditAt:          "audit_at",
	CreatedBy:        "created_by",
	DeptId:           "dept_id",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
}

// NewPlayCoachApplyDao creates and returns a new DAO object for table data access.
func NewPlayCoachApplyDao(handlers ...gdb.ModelHandler) *PlayCoachApplyDao {
	return &PlayCoachApplyDao{
		group:    "default",
		table:    "play_coach_apply",
		columns:  playCoachApplyColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayCoachApplyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayCoachApplyDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayCoachApplyDao) Columns() PlayCoachApplyColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayCoachApplyDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayCoachApplyDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayCoachApplyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
