// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayPayment is the golang structure for table play_payment.
type PlayPayment struct {
	Id              uint64      `orm:"id"               description:"æ”¯ä»˜è®°å½•IDï¼ˆSnowflakeï¼‰"`                                      // æ”¯ä»˜è®°å½•IDï¼ˆSnowflakeï¼‰
	OrderId         uint64      `orm:"order_id"         description:"è®¢å•ID"`                                                           // è®¢å•ID
	MemberId        uint64      `orm:"member_id"        description:"ä¼šå‘˜ID"`                                                           // ä¼šå‘˜ID
	PaymentNo       string      `orm:"payment_no"       description:"æ”¯ä»˜æµæ°´å·"`                                                    // æ”¯ä»˜æµæ°´å·
	TradeNo         string      `orm:"trade_no"         description:"ç¬¬ä¸‰æ–¹äº¤æ˜“å·"`                                                 // ç¬¬ä¸‰æ–¹äº¤æ˜“å·
	PayType         int         `orm:"pay_type"         description:"æ”¯ä»˜æ–¹å¼:1=å¾®ä¿¡æ”¯ä»˜,2=æ”¯ä»˜å®æ”¯ä»˜,3=ä½™é¢æ”¯ä»˜"`       // æ”¯ä»˜æ–¹å¼:1=å¾®ä¿¡æ”¯ä»˜,2=æ”¯ä»˜å®æ”¯ä»˜,3=ä½™é¢æ”¯ä»˜
	PayAmount       int64       `orm:"pay_amount"       description:"æ”¯ä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰"`                                              // æ”¯ä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰
	PayStatus       int         `orm:"pay_status"       description:"æ”¯ä»˜çŠ¶æ€:0=å¾…æ”¯ä»˜,1=æ”¯ä»˜æˆåŠŸ,2=æ”¯ä»˜å¤±è´¥,3=å·²é€€æ¬¾"` // æ”¯ä»˜çŠ¶æ€:0=å¾…æ”¯ä»˜,1=æ”¯ä»˜æˆåŠŸ,2=æ”¯ä»˜å¤±è´¥,3=å·²é€€æ¬¾
	PayAt           *gtime.Time `orm:"pay_at"           description:"æ”¯ä»˜æˆåŠŸæ—¶é—´"`                                                 // æ”¯ä»˜æˆåŠŸæ—¶é—´
	RefundAt        *gtime.Time `orm:"refund_at"        description:"é€€æ¬¾æ—¶é—´"`                                                       // é€€æ¬¾æ—¶é—´
	RefundAmount    int64       `orm:"refund_amount"    description:"é€€æ¬¾é‡‘é¢ï¼ˆåˆ†ï¼‰"`                                              // é€€æ¬¾é‡‘é¢ï¼ˆåˆ†ï¼‰
	CallbackContent string      `orm:"callback_content" description:"å›žè°ƒæŠ¥æ–‡"`                                                       // å›žè°ƒæŠ¥æ–‡
	CreatedBy       uint64      `orm:"created_by"       description:"åˆ›å»ºäººID"`                                                        // åˆ›å»ºäººID
	DeptId          uint64      `orm:"dept_id"          description:"æ‰€å±žéƒ¨é—¨ID"`                                                     // æ‰€å±žéƒ¨é—¨ID
	CreatedAt       *gtime.Time `orm:"created_at"       description:"åˆ›å»ºæ—¶é—´"`                                                       // åˆ›å»ºæ—¶é—´
	UpdatedAt       *gtime.Time `orm:"updated_at"       description:"æ›´æ–°æ—¶é—´"`                                                       // æ›´æ–°æ—¶é—´
	DeletedAt       *gtime.Time `orm:"deleted_at"       description:"è½¯åˆ é™¤æ—¶é—´"`                                                    // è½¯åˆ é™¤æ—¶é—´
}
