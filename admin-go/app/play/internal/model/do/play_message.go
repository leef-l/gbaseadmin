// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayMessage is the golang structure of table play_message for DAO operations like Where/Data.
type PlayMessage struct {
	g.Meta    `orm:"table:play_message, do:true"`
	Id        any         // 消息ID
	MemberId  any         // 接收者会员ID
	Title     any         // 消息标题
	Content   any         // 消息内容
	MsgType   any         // 消息类型 1=系统通知 2=订单消息 3=活动消息
	BizId     any         // 关联业务ID（订单ID/活动ID等）
	IsRead    any         // 是否已读 0=未读 1=已读
	Status    any         // 状态 1=正常 0=禁用
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
