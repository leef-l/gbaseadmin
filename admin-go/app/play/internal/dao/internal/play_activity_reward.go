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
	Id            string // 奖励ID（Snowflake）
	ActivityId    string // 活动ID
	RewardType    string // 奖励类型:1=余额,2=优惠券,3=经验值,4=会员等级天数
	RewardValue   string // 奖励数值（余额=分，优惠券=coupon_id，经验=值，等级天数=天）
	RewardLevelId string // 会员等级ID（type=4时使用）
	RewardName    string // 奖励名称（展示用，如"送50元余额"）
	Sort          string // 排序（升序）
	CreatedBy     string // 创建人ID
	DeptId        string // 所属部门ID
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	DeletedAt     string // 软删除时间
}

// playActivityRewardColumns holds the columns for the table play_activity_reward.
var playActivityRewardColumns = PlayActivityRewardColumns{
	Id:            "id",
	ActivityId:    "activity_id",
	RewardType:    "reward_type",
	RewardValue:   "reward_value",
	RewardLevelId: "reward_level_id",
	RewardName:    "reward_name",
	Sort:          "sort",
	CreatedBy:     "created_by",
	DeptId:        "dept_id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
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
