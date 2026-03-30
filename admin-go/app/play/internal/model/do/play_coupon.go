// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayCoupon is the golang structure of table play_coupon for DAO operations like Where/Data.
type PlayCoupon struct {
	g.Meta       `orm:"table:play_coupon, do:true"`
	Id           any         // ä¼˜æƒ åˆ¸IDï¼ˆSnowflakeï¼‰
	Title        any         // ä¼˜æƒ åˆ¸åç§°
	Type         any         // ä¼˜æƒ åˆ¸ç±»åž‹:1=æ»¡å‡åˆ¸,2=æŠ˜æ‰£åˆ¸,3=æ— é—¨æ§›åˆ¸
	IsNewMember  any         // æ˜¯å¦æ–°äººä¸“äº«:0=å¦,1=æ˜¯
	FaceValue    any         // é¢å€¼ï¼ˆåˆ†ï¼‰
	MinAmount    any         // æœ€ä½Žæ¶ˆè´¹é‡‘é¢ï¼ˆåˆ†ï¼‰
	TotalNum     any         // å‘æ”¾æ€»é‡ï¼ˆ0ä¸é™ï¼‰
	UsedNum      any         // å·²ä½¿ç”¨æ•°é‡
	ClaimNum     any         // å·²é¢†å–æ•°é‡
	PerLimit     any         // æ¯äººé™é¢†å¼ æ•°
	ValidStartAt *gtime.Time // æœ‰æ•ˆæœŸå¼€å§‹æ—¶é—´
	ValidEndAt   *gtime.Time // æœ‰æ•ˆæœŸç»“æŸæ—¶é—´
	Sort         any         // æŽ’åº
	Status       any         // çŠ¶æ€:0=å…³é—­,1=å¼€å¯
	CreatedBy    any         // åˆ›å»ºäººID
	DeptId       any         // æ‰€å±žéƒ¨é—¨ID
	CreatedAt    *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt    *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt    *gtime.Time // è½¯åˆ é™¤æ—¶é—´
}
