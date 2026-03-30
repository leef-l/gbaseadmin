// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayRechargePlan is the golang structure for table play_recharge_plan.
type PlayRechargePlan struct {
	Id         uint64      `orm:"id"          description:"æ–¹æ¡ˆIDï¼ˆSnowflakeï¼‰"`  // æ–¹æ¡ˆIDï¼ˆSnowflakeï¼‰
	Title      string      `orm:"title"       description:"æ–¹æ¡ˆåç§°"`             // æ–¹æ¡ˆåç§°
	Amount     int64       `orm:"amount"      description:"å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰"`    // å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰
	GiftAmount int64       `orm:"gift_amount" description:"èµ é€é‡‘é¢ï¼ˆåˆ†ï¼‰"`    // èµ é€é‡‘é¢ï¼ˆåˆ†ï¼‰
	CoverImage string      `orm:"cover_image" description:"æ–¹æ¡ˆå°é¢å›¾"`          // æ–¹æ¡ˆå°é¢å›¾
	Sort       int         `orm:"sort"        description:"æŽ’åºï¼ˆå‡åºï¼‰"`       // æŽ’åºï¼ˆå‡åºï¼‰
	Status     int         `orm:"status"      description:"çŠ¶æ€:0=å…³é—­,1=å¼€å¯"` // çŠ¶æ€:0=å…³é—­,1=å¼€å¯
	CreatedBy  uint64      `orm:"created_by"  description:"åˆ›å»ºäººID"`              // åˆ›å»ºäººID
	DeptId     uint64      `orm:"dept_id"     description:"æ‰€å±žéƒ¨é—¨ID"`           // æ‰€å±žéƒ¨é—¨ID
	CreatedAt  *gtime.Time `orm:"created_at"  description:"åˆ›å»ºæ—¶é—´"`             // åˆ›å»ºæ—¶é—´
	UpdatedAt  *gtime.Time `orm:"updated_at"  description:"æ›´æ–°æ—¶é—´"`             // æ›´æ–°æ—¶é—´
	DeletedAt  *gtime.Time `orm:"deleted_at"  description:"è½¯åˆ é™¤æ—¶é—´"`          // è½¯åˆ é™¤æ—¶é—´
}
