// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayMessage is the golang structure for table play_message.
type PlayMessage struct {
	Id        int64       `orm:"id"         description:"消息ID"`                      // 消息ID
	MemberId  int64       `orm:"member_id"  description:"接收者会员ID"`                   // 接收者会员ID
	Title     string      `orm:"title"      description:"消息标题"`                      // 消息标题
	Content   string      `orm:"content"    description:"消息内容"`                      // 消息内容
	MsgType   int         `orm:"msg_type"   description:"消息类型 1=系统通知 2=订单消息 3=活动消息"` // 消息类型 1=系统通知 2=订单消息 3=活动消息
	BizId     string      `orm:"biz_id"     description:"关联业务ID（订单ID/活动ID等）"`        // 关联业务ID（订单ID/活动ID等）
	IsRead    int         `orm:"is_read"    description:"是否已读 0=未读 1=已读"`            // 是否已读 0=未读 1=已读
	Status    int         `orm:"status"     description:"状态 1=正常 0=禁用"`              // 状态 1=正常 0=禁用
	CreatedAt *gtime.Time `orm:"created_at" description:"创建时间"`                      // 创建时间
	UpdatedAt *gtime.Time `orm:"updated_at" description:"更新时间"`                      // 更新时间
	DeletedAt *gtime.Time `orm:"deleted_at" description:"删除时间"`                      // 删除时间
}
