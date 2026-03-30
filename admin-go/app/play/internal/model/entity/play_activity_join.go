// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayActivityJoin is the golang structure for table play_activity_join.
type PlayActivityJoin struct {
	Id          uint64      `orm:"id"           description:"è®°å½•IDï¼ˆSnowflakeï¼‰"`                                      // è®°å½•IDï¼ˆSnowflakeï¼‰
	ActivityId  uint64      `orm:"activity_id"  description:"æ´»åŠ¨ID"`                                                     // æ´»åŠ¨ID
	MemberId    uint64      `orm:"member_id"    description:"ä¼šå‘˜ID"`                                                     // ä¼šå‘˜ID
	JoinStatus  int         `orm:"join_status"  description:"å‚ä¸ŽçŠ¶æ€:0=å·²æŠ¥å,1=è¿›è¡Œä¸­,2=å·²å®Œæˆ,3=å·²é¢†å¥–"` // å‚ä¸ŽçŠ¶æ€:0=å·²æŠ¥å,1=è¿›è¡Œä¸­,2=å·²å®Œæˆ,3=å·²é¢†å¥–
	CurrentStep int         `orm:"current_step" description:"å½“å‰å®Œæˆåˆ°ç¬¬å‡ æ­¥"`                                     // å½“å‰å®Œæˆåˆ°ç¬¬å‡ æ­¥
	FinishAt    *gtime.Time `orm:"finish_at"    description:"å®Œæˆæ—¶é—´"`                                                 // å®Œæˆæ—¶é—´
	RewardAt    *gtime.Time `orm:"reward_at"    description:"é¢†å¥–æ—¶é—´"`                                                 // é¢†å¥–æ—¶é—´
	Remark      string      `orm:"remark"       description:"å¤‡æ³¨"`                                                       // å¤‡æ³¨
	CreatedBy   uint64      `orm:"created_by"   description:"åˆ›å»ºäººID"`                                                  // åˆ›å»ºäººID
	DeptId      uint64      `orm:"dept_id"      description:"æ‰€å±žéƒ¨é—¨ID"`                                               // æ‰€å±žéƒ¨é—¨ID
	CreatedAt   *gtime.Time `orm:"created_at"   description:"åˆ›å»ºæ—¶é—´"`                                                 // åˆ›å»ºæ—¶é—´
	UpdatedAt   *gtime.Time `orm:"updated_at"   description:"æ›´æ–°æ—¶é—´"`                                                 // æ›´æ–°æ—¶é—´
	DeletedAt   *gtime.Time `orm:"deleted_at"   description:"è½¯åˆ é™¤æ—¶é—´"`                                              // è½¯åˆ é™¤æ—¶é—´
}
