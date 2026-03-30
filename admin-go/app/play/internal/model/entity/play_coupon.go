// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayCoupon is the golang structure for table play_coupon.
type PlayCoupon struct {
	Id           uint64      `orm:"id"             description:"ä¼˜æƒ åˆ¸IDï¼ˆSnowflakeï¼‰"`                             // ä¼˜æƒ åˆ¸IDï¼ˆSnowflakeï¼‰
	Title        string      `orm:"title"          description:"ä¼˜æƒ åˆ¸åç§°"`                                        // ä¼˜æƒ åˆ¸åç§°
	Type         int         `orm:"type"           description:"ä¼˜æƒ åˆ¸ç±»åž‹:1=æ»¡å‡åˆ¸,2=æŠ˜æ‰£åˆ¸,3=æ— é—¨æ§›åˆ¸"` // ä¼˜æƒ åˆ¸ç±»åž‹:1=æ»¡å‡åˆ¸,2=æŠ˜æ‰£åˆ¸,3=æ— é—¨æ§›åˆ¸
	IsNewMember  int         `orm:"is_new_member"  description:"æ˜¯å¦æ–°äººä¸“äº«:0=å¦,1=æ˜¯"`                         // æ˜¯å¦æ–°äººä¸“äº«:0=å¦,1=æ˜¯
	FaceValue    int64       `orm:"face_value"     description:"é¢å€¼ï¼ˆåˆ†ï¼‰"`                                        // é¢å€¼ï¼ˆåˆ†ï¼‰
	MinAmount    int64       `orm:"min_amount"     description:"æœ€ä½Žæ¶ˆè´¹é‡‘é¢ï¼ˆåˆ†ï¼‰"`                            // æœ€ä½Žæ¶ˆè´¹é‡‘é¢ï¼ˆåˆ†ï¼‰
	TotalNum     int         `orm:"total_num"      description:"å‘æ”¾æ€»é‡ï¼ˆ0ä¸é™ï¼‰"`                              // å‘æ”¾æ€»é‡ï¼ˆ0ä¸é™ï¼‰
	UsedNum      int         `orm:"used_num"       description:"å·²ä½¿ç”¨æ•°é‡"`                                        // å·²ä½¿ç”¨æ•°é‡
	ClaimNum     int         `orm:"claim_num"      description:"å·²é¢†å–æ•°é‡"`                                        // å·²é¢†å–æ•°é‡
	PerLimit     int         `orm:"per_limit"      description:"æ¯äººé™é¢†å¼ æ•°"`                                     // æ¯äººé™é¢†å¼ æ•°
	ValidStartAt *gtime.Time `orm:"valid_start_at" description:"æœ‰æ•ˆæœŸå¼€å§‹æ—¶é—´"`                                  // æœ‰æ•ˆæœŸå¼€å§‹æ—¶é—´
	ValidEndAt   *gtime.Time `orm:"valid_end_at"   description:"æœ‰æ•ˆæœŸç»“æŸæ—¶é—´"`                                  // æœ‰æ•ˆæœŸç»“æŸæ—¶é—´
	Sort         int         `orm:"sort"           description:"æŽ’åº"`                                                 // æŽ’åº
	Status       int         `orm:"status"         description:"çŠ¶æ€:0=å…³é—­,1=å¼€å¯"`                               // çŠ¶æ€:0=å…³é—­,1=å¼€å¯
	CreatedBy    uint64      `orm:"created_by"     description:"åˆ›å»ºäººID"`                                            // åˆ›å»ºäººID
	DeptId       uint64      `orm:"dept_id"        description:"æ‰€å±žéƒ¨é—¨ID"`                                         // æ‰€å±žéƒ¨é—¨ID
	CreatedAt    *gtime.Time `orm:"created_at"     description:"åˆ›å»ºæ—¶é—´"`                                           // åˆ›å»ºæ—¶é—´
	UpdatedAt    *gtime.Time `orm:"updated_at"     description:"æ›´æ–°æ—¶é—´"`                                           // æ›´æ–°æ—¶é—´
	DeletedAt    *gtime.Time `orm:"deleted_at"     description:"è½¯åˆ é™¤æ—¶é—´"`                                        // è½¯åˆ é™¤æ—¶é—´
}
