// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayReview is the golang structure of table play_review for DAO operations like Where/Data.
type PlayReview struct {
	g.Meta        `orm:"table:play_review, do:true"`
	Id            any         // è¯„ä»·IDï¼ˆSnowflakeï¼‰
	OrderId       any         // è®¢å•ID
	MemberId      any         // è¯„ä»·ä¼šå‘˜ID
	CoachId       any         // è¢«è¯„é™ªçŽ©å¸ˆID
	Score         any         // è¯„åˆ†ï¼ˆä¹˜100ï¼‰
	ReviewContent any         // è¯„ä»·å†…å®¹
	ReviewImage   any         // è¯„ä»·å›¾ç‰‡ï¼ˆå¤šå¼ é€—å·åˆ†éš”ï¼‰
	ReplyContent  any         // é™ªçŽ©å¸ˆå›žå¤å†…å®¹
	ReplyAt       *gtime.Time // å›žå¤æ—¶é—´
	IsAnonymous   any         // æ˜¯å¦åŒ¿å:0=å¦,1=æ˜¯
	Status        any         // çŠ¶æ€:0=éšè—,1=æ˜¾ç¤º
	CreatedBy     any         // åˆ›å»ºäººID
	DeptId        any         // æ‰€å±žéƒ¨é—¨ID
	CreatedAt     *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt     *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt     *gtime.Time // è½¯åˆ é™¤æ—¶é—´
}
