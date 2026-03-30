// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayActivityReward is the golang structure of table play_activity_reward for DAO operations like Where/Data.
type PlayActivityReward struct {
	g.Meta      `orm:"table:play_activity_reward, do:true"`
	Id          any         // å¥–åŠ±IDï¼ˆSnowflakeï¼‰
	ActivityId  any         // æ´»åŠ¨ID
	RewardType  any         // å¥–åŠ±ç±»åž‹:1=ä½™é¢,2=ä¼˜æƒ åˆ¸,3=ç»éªŒå€¼,4=ä¼šå‘˜ç­‰çº§å¤©æ•°
	RewardValue any         // å¥–åŠ±æ•°å€¼
	RewardName  any         // å¥–åŠ±åç§°
	Sort        any         // æŽ’åº
	CreatedBy   any         // åˆ›å»ºäººID
	DeptId      any         // æ‰€å±žéƒ¨é—¨ID
	CreatedAt   *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt   *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt   *gtime.Time // è½¯åˆ é™¤æ—¶é—´
}
