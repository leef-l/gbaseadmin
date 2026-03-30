// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayCouponMember is the golang structure of table play_coupon_member for DAO operations like Where/Data.
type PlayCouponMember struct {
	g.Meta    `orm:"table:play_coupon_member, do:true"`
	Id        any         // è®°å½•IDï¼ˆSnowflakeï¼‰
	CouponId  any         // ä¼˜æƒ åˆ¸æ¨¡æ¿ID
	MemberId  any         // ä¼šå‘˜ID
	OrderId   any         // ä½¿ç”¨çš„è®¢å•ID
	UseStatus any         // ä½¿ç”¨çŠ¶æ€:0=æœªä½¿ç”¨,1=å·²ä½¿ç”¨,2=å·²è¿‡æœŸ
	ClaimAt   *gtime.Time // é¢†å–æ—¶é—´
	UseAt     *gtime.Time // ä½¿ç”¨æ—¶é—´
	ExpireAt  *gtime.Time // è¿‡æœŸæ—¶é—´
	CreatedBy any         // åˆ›å»ºäººID
	DeptId    any         // æ‰€å±žéƒ¨é—¨ID
	CreatedAt *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt *gtime.Time // è½¯åˆ é™¤æ—¶é—´
}
