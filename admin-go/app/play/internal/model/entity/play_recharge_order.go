// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayRechargeOrder is the golang structure for table play_recharge_order.
type PlayRechargeOrder struct {
	Id             uint64      `orm:"id"               description:"å……å€¼è®¢å•IDï¼ˆSnowflakeï¼‰"`                          // å……å€¼è®¢å•IDï¼ˆSnowflakeï¼‰
	OrderNo        string      `orm:"order_no"         description:"å……å€¼è®¢å•å·"`                                        // å……å€¼è®¢å•å·
	MemberId       uint64      `orm:"member_id"        description:"ä¼šå‘˜ID"`                                               // ä¼šå‘˜ID
	RechargePlanId uint64      `orm:"recharge_plan_id" description:"å……å€¼æ–¹æ¡ˆID"`                                         // å……å€¼æ–¹æ¡ˆID
	Amount         int64       `orm:"amount"           description:"å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰"`                                  // å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰
	GiftAmount     int64       `orm:"gift_amount"      description:"èµ é€é‡‘é¢ï¼ˆåˆ†ï¼‰"`                                  // èµ é€é‡‘é¢ï¼ˆåˆ†ï¼‰
	PayType        int         `orm:"pay_type"         description:"æ”¯ä»˜æ–¹å¼:1=å¾®ä¿¡æ”¯ä»˜,2=æ”¯ä»˜å®æ”¯ä»˜"`          // æ”¯ä»˜æ–¹å¼:1=å¾®ä¿¡æ”¯ä»˜,2=æ”¯ä»˜å®æ”¯ä»˜
	TradeNo        string      `orm:"trade_no"         description:"ç¬¬ä¸‰æ–¹äº¤æ˜“å·"`                                     // ç¬¬ä¸‰æ–¹äº¤æ˜“å·
	PayStatus      int         `orm:"pay_status"       description:"æ”¯ä»˜çŠ¶æ€:0=å¾…æ”¯ä»˜,1=æ”¯ä»˜æˆåŠŸ,2=æ”¯ä»˜å¤±è´¥"` // æ”¯ä»˜çŠ¶æ€:0=å¾…æ”¯ä»˜,1=æ”¯ä»˜æˆåŠŸ,2=æ”¯ä»˜å¤±è´¥
	PayAt          *gtime.Time `orm:"pay_at"           description:"æ”¯ä»˜æ—¶é—´"`                                           // æ”¯ä»˜æ—¶é—´
	CreatedBy      uint64      `orm:"created_by"       description:"åˆ›å»ºäººID"`                                            // åˆ›å»ºäººID
	DeptId         uint64      `orm:"dept_id"          description:"æ‰€å±žéƒ¨é—¨ID"`                                         // æ‰€å±žéƒ¨é—¨ID
	CreatedAt      *gtime.Time `orm:"created_at"       description:"åˆ›å»ºæ—¶é—´"`                                           // åˆ›å»ºæ—¶é—´
	UpdatedAt      *gtime.Time `orm:"updated_at"       description:"æ›´æ–°æ—¶é—´"`                                           // æ›´æ–°æ—¶é—´
	DeletedAt      *gtime.Time `orm:"deleted_at"       description:"è½¯åˆ é™¤æ—¶é—´"`                                        // è½¯åˆ é™¤æ—¶é—´
}
