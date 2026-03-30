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
	Id             string // æ´»åŠ¨IDï¼ˆSnowflakeï¼‰
	Title          string // æ´»åŠ¨åç§°
	CoverImage     string // æ´»åŠ¨å°é¢å›¾
	DescContent    string // æ´»åŠ¨è¯¦æƒ…æè¿°
	Type           string // æ´»åŠ¨ç±»åž‹:1=å……å€¼æ´»åŠ¨,2=ä¸‹å•æ´»åŠ¨,3=æ³¨å†Œæ´»åŠ¨,4=å›¾æ–‡æ­¥éª¤æ´»åŠ¨,5=è‡ªå®šä¹‰æ´»åŠ¨
	ConditionType  string // å‚ä¸Žæ¡ä»¶:0=æ— æ¡ä»¶,1=éœ€æŠ¥å,2=å……å€¼æ»¡é¢,3=ä¸‹å•æ»¡é¢,4=å®Œæˆæ­¥éª¤
	ConditionValue string // æ¡ä»¶å€¼
	IsAutoReward   string // æ˜¯å¦è‡ªåŠ¨å‘å¥–:0=å¦,1=æ˜¯
	StartAt        string // æ´»åŠ¨å¼€å§‹æ—¶é—´
	EndAt          string // æ´»åŠ¨ç»“æŸæ—¶é—´
	MaxNum         string // å‚ä¸Žäººæ•°ä¸Šé™ï¼ˆ0ä¸é™ï¼‰
	JoinNum        string // å·²å‚ä¸Žäººæ•°
	Sort           string // æŽ’åº
	Status         string // çŠ¶æ€:0=å…³é—­,1=å¼€å¯
	CreatedBy      string // åˆ›å»ºäººID
	DeptId         string // æ‰€å±žéƒ¨é—¨ID
	CreatedAt      string // åˆ›å»ºæ—¶é—´
	UpdatedAt      string // æ›´æ–°æ—¶é—´
	DeletedAt      string // è½¯åˆ é™¤æ—¶é—´
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
