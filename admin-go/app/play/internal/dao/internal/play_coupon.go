// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayCouponDao is the data access object for the table play_coupon.
type PlayCouponDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PlayCouponColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PlayCouponColumns defines and stores column names for the table play_coupon.
type PlayCouponColumns struct {
	Id           string // ä¼˜æƒ åˆ¸IDï¼ˆSnowflakeï¼‰
	Title        string // ä¼˜æƒ åˆ¸åç§°
	Type         string // ä¼˜æƒ åˆ¸ç±»åž‹:1=æ»¡å‡åˆ¸,2=æŠ˜æ‰£åˆ¸,3=æ— é—¨æ§›åˆ¸
	IsNewMember  string // æ˜¯å¦æ–°äººä¸“äº«:0=å¦,1=æ˜¯
	FaceValue    string // é¢å€¼ï¼ˆåˆ†ï¼‰
	MinAmount    string // æœ€ä½Žæ¶ˆè´¹é‡‘é¢ï¼ˆåˆ†ï¼‰
	TotalNum     string // å‘æ”¾æ€»é‡ï¼ˆ0ä¸é™ï¼‰
	UsedNum      string // å·²ä½¿ç”¨æ•°é‡
	ClaimNum     string // å·²é¢†å–æ•°é‡
	PerLimit     string // æ¯äººé™é¢†å¼ æ•°
	ValidStartAt string // æœ‰æ•ˆæœŸå¼€å§‹æ—¶é—´
	ValidEndAt   string // æœ‰æ•ˆæœŸç»“æŸæ—¶é—´
	Sort         string // æŽ’åº
	Status       string // çŠ¶æ€:0=å…³é—­,1=å¼€å¯
	CreatedBy    string // åˆ›å»ºäººID
	DeptId       string // æ‰€å±žéƒ¨é—¨ID
	CreatedAt    string // åˆ›å»ºæ—¶é—´
	UpdatedAt    string // æ›´æ–°æ—¶é—´
	DeletedAt    string // è½¯åˆ é™¤æ—¶é—´
}

// playCouponColumns holds the columns for the table play_coupon.
var playCouponColumns = PlayCouponColumns{
	Id:           "id",
	Title:        "title",
	Type:         "type",
	IsNewMember:  "is_new_member",
	FaceValue:    "face_value",
	MinAmount:    "min_amount",
	TotalNum:     "total_num",
	UsedNum:      "used_num",
	ClaimNum:     "claim_num",
	PerLimit:     "per_limit",
	ValidStartAt: "valid_start_at",
	ValidEndAt:   "valid_end_at",
	Sort:         "sort",
	Status:       "status",
	CreatedBy:    "created_by",
	DeptId:       "dept_id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// NewPlayCouponDao creates and returns a new DAO object for table data access.
func NewPlayCouponDao(handlers ...gdb.ModelHandler) *PlayCouponDao {
	return &PlayCouponDao{
		group:    "default",
		table:    "play_coupon",
		columns:  playCouponColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayCouponDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayCouponDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayCouponDao) Columns() PlayCouponColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayCouponDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayCouponDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayCouponDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
