// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayReviewDao is the data access object for the table play_review.
type PlayReviewDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PlayReviewColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PlayReviewColumns defines and stores column names for the table play_review.
type PlayReviewColumns struct {
	Id            string // 评价ID（Snowflake）
	OrderId       string // 订单ID
	MemberId      string // 评价会员ID
	CoachId       string // 被评陪玩师ID
	Score         string // 评分（乘100，如 500=5.00分）
	ReviewContent string // 评价内容
	ReviewImage   string // 评价图片（多张逗号分隔）
	ReplyContent  string // 陪玩师回复内容
	ReplyAt       string // 回复时间
	IsAnonymous   string // 是否匿名:0=否,1=是
	Status        string // 状态:0=隐藏,1=显示
	CreatedBy     string // 创建人ID
	DeptId        string // 所属部门ID
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	DeletedAt     string // 软删除时间
}

// playReviewColumns holds the columns for the table play_review.
var playReviewColumns = PlayReviewColumns{
	Id:            "id",
	OrderId:       "order_id",
	MemberId:      "member_id",
	CoachId:       "coach_id",
	Score:         "score",
	ReviewContent: "review_content",
	ReviewImage:   "review_image",
	ReplyContent:  "reply_content",
	ReplyAt:       "reply_at",
	IsAnonymous:   "is_anonymous",
	Status:        "status",
	CreatedBy:     "created_by",
	DeptId:        "dept_id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewPlayReviewDao creates and returns a new DAO object for table data access.
func NewPlayReviewDao(handlers ...gdb.ModelHandler) *PlayReviewDao {
	return &PlayReviewDao{
		group:    "default",
		table:    "play_review",
		columns:  playReviewColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayReviewDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayReviewDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayReviewDao) Columns() PlayReviewColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayReviewDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayReviewDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayReviewDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
