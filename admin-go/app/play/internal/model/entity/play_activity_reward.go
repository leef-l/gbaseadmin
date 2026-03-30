// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayActivityReward is the golang structure for table play_activity_reward.
type PlayActivityReward struct {
	Id          uint64      `orm:"id"           description:"å¥–åŠ±IDï¼ˆSnowflakeï¼‰"`                                            // å¥–åŠ±IDï¼ˆSnowflakeï¼‰
	ActivityId  uint64      `orm:"activity_id"  description:"æ´»åŠ¨ID"`                                                           // æ´»åŠ¨ID
	RewardType  int         `orm:"reward_type"  description:"å¥–åŠ±ç±»åž‹:1=ä½™é¢,2=ä¼˜æƒ åˆ¸,3=ç»éªŒå€¼,4=ä¼šå‘˜ç­‰çº§å¤©æ•°"` // å¥–åŠ±ç±»åž‹:1=ä½™é¢,2=ä¼˜æƒ åˆ¸,3=ç»éªŒå€¼,4=ä¼šå‘˜ç­‰çº§å¤©æ•°
	RewardValue int64       `orm:"reward_value" description:"å¥–åŠ±æ•°å€¼"`                                                       // å¥–åŠ±æ•°å€¼
	RewardName  string      `orm:"reward_name"  description:"å¥–åŠ±åç§°"`                                                       // å¥–åŠ±åç§°
	Sort        int         `orm:"sort"         description:"æŽ’åº"`                                                             // æŽ’åº
	CreatedBy   uint64      `orm:"created_by"   description:"åˆ›å»ºäººID"`                                                        // åˆ›å»ºäººID
	DeptId      uint64      `orm:"dept_id"      description:"æ‰€å±žéƒ¨é—¨ID"`                                                     // æ‰€å±žéƒ¨é—¨ID
	CreatedAt   *gtime.Time `orm:"created_at"   description:"åˆ›å»ºæ—¶é—´"`                                                       // åˆ›å»ºæ—¶é—´
	UpdatedAt   *gtime.Time `orm:"updated_at"   description:"æ›´æ–°æ—¶é—´"`                                                       // æ›´æ–°æ—¶é—´
	DeletedAt   *gtime.Time `orm:"deleted_at"   description:"è½¯åˆ é™¤æ—¶é—´"`                                                    // è½¯åˆ é™¤æ—¶é—´
}
