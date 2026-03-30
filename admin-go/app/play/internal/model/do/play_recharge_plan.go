// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayRechargePlan is the golang structure of table play_recharge_plan for DAO operations like Where/Data.
type PlayRechargePlan struct {
	g.Meta     `orm:"table:play_recharge_plan, do:true"`
	Id         any         // æ–¹æ¡ˆIDï¼ˆSnowflakeï¼‰
	Title      any         // æ–¹æ¡ˆåç§°
	Amount     any         // å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰
	GiftAmount any         // èµ é€é‡‘é¢ï¼ˆåˆ†ï¼‰
	CoverImage any         // æ–¹æ¡ˆå°é¢å›¾
	Sort       any         // æŽ’åºï¼ˆå‡åºï¼‰
	Status     any         // çŠ¶æ€:0=å…³é—­,1=å¼€å¯
	CreatedBy  any         // åˆ›å»ºäººID
	DeptId     any         // æ‰€å±žéƒ¨é—¨ID
	CreatedAt  *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt  *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt  *gtime.Time // è½¯åˆ é™¤æ—¶é—´
}
