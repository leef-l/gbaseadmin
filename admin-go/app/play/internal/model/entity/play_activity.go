// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayActivity is the golang structure for table play_activity.
type PlayActivity struct {
	Id             uint64      `orm:"id"              description:"æ´»åŠ¨IDï¼ˆSnowflakeï¼‰"`                                                                          // æ´»åŠ¨IDï¼ˆSnowflakeï¼‰
	Title          string      `orm:"title"           description:"æ´»åŠ¨åç§°"`                                                                                     // æ´»åŠ¨åç§°
	CoverImage     string      `orm:"cover_image"     description:"æ´»åŠ¨å°é¢å›¾"`                                                                                  // æ´»åŠ¨å°é¢å›¾
	DescContent    string      `orm:"desc_content"    description:"æ´»åŠ¨è¯¦æƒ…æè¿°"`                                                                               // æ´»åŠ¨è¯¦æƒ…æè¿°
	Type           int         `orm:"type"            description:"æ´»åŠ¨ç±»åž‹:1=å……å€¼æ´»åŠ¨,2=ä¸‹å•æ´»åŠ¨,3=æ³¨å†Œæ´»åŠ¨,4=å›¾æ–‡æ­¥éª¤æ´»åŠ¨,5=è‡ªå®šä¹‰æ´»åŠ¨"` // æ´»åŠ¨ç±»åž‹:1=å……å€¼æ´»åŠ¨,2=ä¸‹å•æ´»åŠ¨,3=æ³¨å†Œæ´»åŠ¨,4=å›¾æ–‡æ­¥éª¤æ´»åŠ¨,5=è‡ªå®šä¹‰æ´»åŠ¨
	ConditionType  int         `orm:"condition_type"  description:"å‚ä¸Žæ¡ä»¶:0=æ— æ¡ä»¶,1=éœ€æŠ¥å,2=å……å€¼æ»¡é¢,3=ä¸‹å•æ»¡é¢,4=å®Œæˆæ­¥éª¤"`                // å‚ä¸Žæ¡ä»¶:0=æ— æ¡ä»¶,1=éœ€æŠ¥å,2=å……å€¼æ»¡é¢,3=ä¸‹å•æ»¡é¢,4=å®Œæˆæ­¥éª¤
	ConditionValue int64       `orm:"condition_value" description:"æ¡ä»¶å€¼"`                                                                                        // æ¡ä»¶å€¼
	IsAutoReward   int         `orm:"is_auto_reward"  description:"æ˜¯å¦è‡ªåŠ¨å‘å¥–:0=å¦,1=æ˜¯"`                                                                   // æ˜¯å¦è‡ªåŠ¨å‘å¥–:0=å¦,1=æ˜¯
	StartAt        *gtime.Time `orm:"start_at"        description:"æ´»åŠ¨å¼€å§‹æ—¶é—´"`                                                                               // æ´»åŠ¨å¼€å§‹æ—¶é—´
	EndAt          *gtime.Time `orm:"end_at"          description:"æ´»åŠ¨ç»“æŸæ—¶é—´"`                                                                               // æ´»åŠ¨ç»“æŸæ—¶é—´
	MaxNum         int         `orm:"max_num"         description:"å‚ä¸Žäººæ•°ä¸Šé™ï¼ˆ0ä¸é™ï¼‰"`                                                                  // å‚ä¸Žäººæ•°ä¸Šé™ï¼ˆ0ä¸é™ï¼‰
	JoinNum        int         `orm:"join_num"        description:"å·²å‚ä¸Žäººæ•°"`                                                                                  // å·²å‚ä¸Žäººæ•°
	Sort           int         `orm:"sort"            description:"æŽ’åº"`                                                                                           // æŽ’åº
	Status         int         `orm:"status"          description:"çŠ¶æ€:0=å…³é—­,1=å¼€å¯"`                                                                         // çŠ¶æ€:0=å…³é—­,1=å¼€å¯
	CreatedBy      uint64      `orm:"created_by"      description:"åˆ›å»ºäººID"`                                                                                      // åˆ›å»ºäººID
	DeptId         uint64      `orm:"dept_id"         description:"æ‰€å±žéƒ¨é—¨ID"`                                                                                   // æ‰€å±žéƒ¨é—¨ID
	CreatedAt      *gtime.Time `orm:"created_at"      description:"åˆ›å»ºæ—¶é—´"`                                                                                     // åˆ›å»ºæ—¶é—´
	UpdatedAt      *gtime.Time `orm:"updated_at"      description:"æ›´æ–°æ—¶é—´"`                                                                                     // æ›´æ–°æ—¶é—´
	DeletedAt      *gtime.Time `orm:"deleted_at"      description:"è½¯åˆ é™¤æ—¶é—´"`                                                                                  // è½¯åˆ é™¤æ—¶é—´
}
