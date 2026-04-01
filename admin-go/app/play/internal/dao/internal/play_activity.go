// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayActivityDao is the data access object for the table play_activity.
type PlayActivityDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  PlayActivityColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// PlayActivityColumns defines and stores column names for the table play_activity.
type PlayActivityColumns struct {
	Id             string // 活动ID（Snowflake）
	Title          string // 活动名称
	CoverImage     string // 活动封面图
	DescContent    string // 活动详情描述（富文本，支持图文混排）
	Type           string // 活动类型:1=充值活动,2=下单活动,3=注册活动,4=图文步骤活动,5=自定义活动
	ConditionType  string // 参与条件:0=无条件,1=需报名,2=充值满额,3=下单满额,4=完成步骤
	ConditionValue string // 条件值（分/次，如充值满5000分、下单满3次）
	IsAutoReward   string // 是否自动发奖:0=否（需审核）,1=是（用户完成即发）
	StartAt        string // 活动开始时间
	EndAt          string // 活动结束时间
	MaxNum         string // 参与人数上限（0表示不限）
	JoinNum        string // 已参与人数
	Sort           string // 排序（升序）
	Status         string // 状态:0=关闭,1=开启
	CreatedBy      string // 创建人ID
	DeptId         string // 所属部门ID
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	DeletedAt      string // 软删除时间
}

// playActivityColumns holds the columns for the table play_activity.
var playActivityColumns = PlayActivityColumns{
	Id:             "id",
	Title:          "title",
	CoverImage:     "cover_image",
	DescContent:    "desc_content",
	Type:           "type",
	ConditionType:  "condition_type",
	ConditionValue: "condition_value",
	IsAutoReward:   "is_auto_reward",
	StartAt:        "start_at",
	EndAt:          "end_at",
	MaxNum:         "max_num",
	JoinNum:        "join_num",
	Sort:           "sort",
	Status:         "status",
	CreatedBy:      "created_by",
	DeptId:         "dept_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewPlayActivityDao creates and returns a new DAO object for table data access.
func NewPlayActivityDao(handlers ...gdb.ModelHandler) *PlayActivityDao {
	return &PlayActivityDao{
		group:    "default",
		table:    "play_activity",
		columns:  playActivityColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayActivityDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayActivityDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayActivityDao) Columns() PlayActivityColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayActivityDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayActivityDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayActivityDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
