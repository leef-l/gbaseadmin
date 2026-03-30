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
	Id            string // è¯„ä»·IDï¼ˆSnowflakeï¼‰
	OrderId       string // è®¢å•ID
	MemberId      string // è¯„ä»·ä¼šå‘˜ID
	CoachId       string // è¢«è¯„é™ªçŽ©å¸ˆID
	Score         string // è¯„åˆ†ï¼ˆä¹˜100ï¼‰
	ReviewContent string // è¯„ä»·å†…å®¹
	ReviewImage   string // è¯„ä»·å›¾ç‰‡ï¼ˆå¤šå¼ é€—å·åˆ†éš”ï¼‰
	ReplyContent  string // é™ªçŽ©å¸ˆå›žå¤å†…å®¹
	ReplyAt       string // å›žå¤æ—¶é—´
	IsAnonymous   string // æ˜¯å¦åŒ¿å:0=å¦,1=æ˜¯
	Status        string // çŠ¶æ€:0=éšè—,1=æ˜¾ç¤º
	CreatedBy     string // åˆ›å»ºäººID
	DeptId        string // æ‰€å±žéƒ¨é—¨ID
	CreatedAt     string // åˆ›å»ºæ—¶é—´
	UpdatedAt     string // æ›´æ–°æ—¶é—´
	DeletedAt     string // è½¯åˆ é™¤æ—¶é—´
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
