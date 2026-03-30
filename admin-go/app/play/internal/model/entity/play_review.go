// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayReview is the golang structure for table play_review.
type PlayReview struct {
	Id            uint64      `orm:"id"             description:"è¯„ä»·IDï¼ˆSnowflakeï¼‰"`              // è¯„ä»·IDï¼ˆSnowflakeï¼‰
	OrderId       uint64      `orm:"order_id"       description:"è®¢å•ID"`                             // è®¢å•ID
	MemberId      uint64      `orm:"member_id"      description:"è¯„ä»·ä¼šå‘˜ID"`                       // è¯„ä»·ä¼šå‘˜ID
	CoachId       uint64      `orm:"coach_id"       description:"è¢«è¯„é™ªçŽ©å¸ˆID"`                    // è¢«è¯„é™ªçŽ©å¸ˆID
	Score         int         `orm:"score"          description:"è¯„åˆ†ï¼ˆä¹˜100ï¼‰"`                   // è¯„åˆ†ï¼ˆä¹˜100ï¼‰
	ReviewContent string      `orm:"review_content" description:"è¯„ä»·å†…å®¹"`                         // è¯„ä»·å†…å®¹
	ReviewImage   string      `orm:"review_image"   description:"è¯„ä»·å›¾ç‰‡ï¼ˆå¤šå¼ é€—å·åˆ†éš”ï¼‰"` // è¯„ä»·å›¾ç‰‡ï¼ˆå¤šå¼ é€—å·åˆ†éš”ï¼‰
	ReplyContent  string      `orm:"reply_content"  description:"é™ªçŽ©å¸ˆå›žå¤å†…å®¹"`                // é™ªçŽ©å¸ˆå›žå¤å†…å®¹
	ReplyAt       *gtime.Time `orm:"reply_at"       description:"å›žå¤æ—¶é—´"`                         // å›žå¤æ—¶é—´
	IsAnonymous   int         `orm:"is_anonymous"   description:"æ˜¯å¦åŒ¿å:0=å¦,1=æ˜¯"`             // æ˜¯å¦åŒ¿å:0=å¦,1=æ˜¯
	Status        int         `orm:"status"         description:"çŠ¶æ€:0=éšè—,1=æ˜¾ç¤º"`             // çŠ¶æ€:0=éšè—,1=æ˜¾ç¤º
	CreatedBy     uint64      `orm:"created_by"     description:"åˆ›å»ºäººID"`                          // åˆ›å»ºäººID
	DeptId        uint64      `orm:"dept_id"        description:"æ‰€å±žéƒ¨é—¨ID"`                       // æ‰€å±žéƒ¨é—¨ID
	CreatedAt     *gtime.Time `orm:"created_at"     description:"åˆ›å»ºæ—¶é—´"`                         // åˆ›å»ºæ—¶é—´
	UpdatedAt     *gtime.Time `orm:"updated_at"     description:"æ›´æ–°æ—¶é—´"`                         // æ›´æ–°æ—¶é—´
	DeletedAt     *gtime.Time `orm:"deleted_at"     description:"è½¯åˆ é™¤æ—¶é—´"`                      // è½¯åˆ é™¤æ—¶é—´
}
