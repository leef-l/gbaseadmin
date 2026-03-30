// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayActivity is the golang structure of table play_activity for DAO operations like Where/Data.
type PlayActivity struct {
	g.Meta         `orm:"table:play_activity, do:true"`
	Id             any         // æ´»åŠ¨IDï¼ˆSnowflakeï¼‰
	Title          any         // æ´»åŠ¨åç§°
	CoverImage     any         // æ´»åŠ¨å°é¢å›¾
	DescContent    any         // æ´»åŠ¨è¯¦æƒ…æè¿°
	Type           any         // æ´»åŠ¨ç±»åž‹:1=å……å€¼æ´»åŠ¨,2=ä¸‹å•æ´»åŠ¨,3=æ³¨å†Œæ´»åŠ¨,4=å›¾æ–‡æ­¥éª¤æ´»åŠ¨,5=è‡ªå®šä¹‰æ´»åŠ¨
	ConditionType  any         // å‚ä¸Žæ¡ä»¶:0=æ— æ¡ä»¶,1=éœ€æŠ¥å,2=å……å€¼æ»¡é¢,3=ä¸‹å•æ»¡é¢,4=å®Œæˆæ­¥éª¤
	ConditionValue any         // æ¡ä»¶å€¼
	IsAutoReward   any         // æ˜¯å¦è‡ªåŠ¨å‘å¥–:0=å¦,1=æ˜¯
	StartAt        *gtime.Time // æ´»åŠ¨å¼€å§‹æ—¶é—´
	EndAt          *gtime.Time // æ´»åŠ¨ç»“æŸæ—¶é—´
	MaxNum         any         // å‚ä¸Žäººæ•°ä¸Šé™ï¼ˆ0ä¸é™ï¼‰
	JoinNum        any         // å·²å‚ä¸Žäººæ•°
	Sort           any         // æŽ’åº
	Status         any         // çŠ¶æ€:0=å…³é—­,1=å¼€å¯
	CreatedBy      any         // åˆ›å»ºäººID
	DeptId         any         // æ‰€å±žéƒ¨é—¨ID
	CreatedAt      *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt      *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt      *gtime.Time // è½¯åˆ é™¤æ—¶é—´
}
