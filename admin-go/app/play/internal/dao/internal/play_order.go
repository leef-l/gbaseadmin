// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayOrderDao is the data access object for the table play_order.
type PlayOrderDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PlayOrderColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PlayOrderColumns defines and stores column names for the table play_order.
type PlayOrderColumns struct {
	Id             string // è®¢å•IDï¼ˆSnowflakeï¼‰
	OrderNo        string // è®¢å•ç¼–å·
	MemberId       string // ä¸‹å•ä¼šå‘˜ID
	CoachId        string // é™ªçŽ©å¸ˆID
	ShopId         string // åº—é“ºID
	GoodsId        string // å•†å“ID
	GoodsTitle     string // å•†å“åç§°ï¼ˆå†—ä½™ï¼‰
	GoodsPrice     string // å•†å“å•ä»·ï¼ˆåˆ†ï¼‰
	Quantity       string // æ•°é‡
	TotalAmount    string // è®¢å•æ€»é¢ï¼ˆåˆ†ï¼‰
	DiscountAmount string // ä¼šå‘˜æŠ˜æ‰£é‡‘é¢ï¼ˆåˆ†ï¼‰
	CouponAmount   string // ä¼˜æƒ åˆ¸æŠµæ‰£é‡‘é¢ï¼ˆåˆ†ï¼‰
	PayAmount      string // å®žä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰
	CouponMemberId string // ä½¿ç”¨çš„ä¼˜æƒ åˆ¸é¢†å–è®°å½•ID
	PayType        string // æ”¯ä»˜æ–¹å¼:0=æœªæ”¯ä»˜,1=å¾®ä¿¡æ”¯ä»˜,2=æ”¯ä»˜å®æ”¯ä»˜,3=ä½™é¢æ”¯ä»˜
	OrderStatus    string // è®¢å•çŠ¶æ€:0=å¾…æ”¯ä»˜,1=å·²æ”¯ä»˜,2=è¿›è¡Œä¸­,3=å·²å®Œæˆ,4=å·²å–æ¶ˆ,5=é€€æ¬¾ä¸­,6=å·²é€€æ¬¾
	PayAt          string // æ”¯ä»˜æ—¶é—´
	StartAt        string // æœåŠ¡å¼€å§‹æ—¶é—´
	FinishAt       string // æœåŠ¡å®Œæˆæ—¶é—´
	CancelAt       string // å–æ¶ˆæ—¶é—´
	CancelReason   string // å–æ¶ˆåŽŸå›
	Remark         string // è®¢å•å¤‡æ³¨
	CreatedBy      string // åˆ›å»ºäººID
	DeptId         string // æ‰€å±žéƒ¨é—¨ID
	CreatedAt      string // åˆ›å»ºæ—¶é—´
	UpdatedAt      string // æ›´æ–°æ—¶é—´
	DeletedAt      string // è½¯åˆ é™¤æ—¶é—´
}

// playOrderColumns holds the columns for the table play_order.
var playOrderColumns = PlayOrderColumns{
	Id:             "id",
	OrderNo:        "order_no",
	MemberId:       "member_id",
	CoachId:        "coach_id",
	ShopId:         "shop_id",
	GoodsId:        "goods_id",
	GoodsTitle:     "goods_title",
	GoodsPrice:     "goods_price",
	Quantity:       "quantity",
	TotalAmount:    "total_amount",
	DiscountAmount: "discount_amount",
	CouponAmount:   "coupon_amount",
	PayAmount:      "pay_amount",
	CouponMemberId: "coupon_member_id",
	PayType:        "pay_type",
	OrderStatus:    "order_status",
	PayAt:          "pay_at",
	StartAt:        "start_at",
	FinishAt:       "finish_at",
	CancelAt:       "cancel_at",
	CancelReason:   "cancel_reason",
	Remark:         "remark",
	CreatedBy:      "created_by",
	DeptId:         "dept_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewPlayOrderDao creates and returns a new DAO object for table data access.
func NewPlayOrderDao(handlers ...gdb.ModelHandler) *PlayOrderDao {
	return &PlayOrderDao{
		group:    "default",
		table:    "play_order",
		columns:  playOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayOrderDao) Columns() PlayOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayOrderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
