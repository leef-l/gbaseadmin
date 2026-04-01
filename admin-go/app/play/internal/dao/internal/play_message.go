// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlayMessageDao is the data access object for the table play_message.
type PlayMessageDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PlayMessageColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PlayMessageColumns defines and stores column names for the table play_message.
type PlayMessageColumns struct {
	Id        string // 消息ID
	MemberId  string // 接收者会员ID
	Title     string // 消息标题
	Content   string // 消息内容
	MsgType   string // 消息类型 1=系统通知 2=订单消息 3=活动消息
	BizId     string // 关联业务ID（订单ID/活动ID等）
	IsRead    string // 是否已读 0=未读 1=已读
	Status    string // 状态 1=正常 0=禁用
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// playMessageColumns holds the columns for the table play_message.
var playMessageColumns = PlayMessageColumns{
	Id:        "id",
	MemberId:  "member_id",
	Title:     "title",
	Content:   "content",
	MsgType:   "msg_type",
	BizId:     "biz_id",
	IsRead:    "is_read",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewPlayMessageDao creates and returns a new DAO object for table data access.
func NewPlayMessageDao(handlers ...gdb.ModelHandler) *PlayMessageDao {
	return &PlayMessageDao{
		group:    "default",
		table:    "play_message",
		columns:  playMessageColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlayMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlayMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlayMessageDao) Columns() PlayMessageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlayMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlayMessageDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlayMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
