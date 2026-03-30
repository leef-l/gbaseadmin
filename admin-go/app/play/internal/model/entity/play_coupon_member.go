// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayCouponMember is the golang structure for table play_coupon_member.
type PlayCouponMember struct {
	Id        uint64      `orm:"id"         description:"è®°å½•IDï¼ˆSnowflakeï¼‰"`                          // è®°å½•IDï¼ˆSnowflakeï¼‰
	CouponId  uint64      `orm:"coupon_id"  description:"ä¼˜æƒ åˆ¸æ¨¡æ¿ID"`                                // ä¼˜æƒ åˆ¸æ¨¡æ¿ID
	MemberId  uint64      `orm:"member_id"  description:"ä¼šå‘˜ID"`                                         // ä¼šå‘˜ID
	OrderId   uint64      `orm:"order_id"   description:"ä½¿ç”¨çš„è®¢å•ID"`                                // ä½¿ç”¨çš„è®¢å•ID
	UseStatus int         `orm:"use_status" description:"ä½¿ç”¨çŠ¶æ€:0=æœªä½¿ç”¨,1=å·²ä½¿ç”¨,2=å·²è¿‡æœŸ"` // ä½¿ç”¨çŠ¶æ€:0=æœªä½¿ç”¨,1=å·²ä½¿ç”¨,2=å·²è¿‡æœŸ
	ClaimAt   *gtime.Time `orm:"claim_at"   description:"é¢†å–æ—¶é—´"`                                     // é¢†å–æ—¶é—´
	UseAt     *gtime.Time `orm:"use_at"     description:"ä½¿ç”¨æ—¶é—´"`                                     // ä½¿ç”¨æ—¶é—´
	ExpireAt  *gtime.Time `orm:"expire_at"  description:"è¿‡æœŸæ—¶é—´"`                                     // è¿‡æœŸæ—¶é—´
	CreatedBy uint64      `orm:"created_by" description:"åˆ›å»ºäººID"`                                      // åˆ›å»ºäººID
	DeptId    uint64      `orm:"dept_id"    description:"æ‰€å±žéƒ¨é—¨ID"`                                   // æ‰€å±žéƒ¨é—¨ID
	CreatedAt *gtime.Time `orm:"created_at" description:"åˆ›å»ºæ—¶é—´"`                                     // åˆ›å»ºæ—¶é—´
	UpdatedAt *gtime.Time `orm:"updated_at" description:"æ›´æ–°æ—¶é—´"`                                     // æ›´æ–°æ—¶é—´
	DeletedAt *gtime.Time `orm:"deleted_at" description:"è½¯åˆ é™¤æ—¶é—´"`                                  // è½¯åˆ é™¤æ—¶é—´
}
