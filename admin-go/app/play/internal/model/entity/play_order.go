// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayOrder is the golang structure for table play_order.
type PlayOrder struct {
	Id             uint64      `orm:"id"               description:"è®¢å•IDï¼ˆSnowflakeï¼‰"`                                                                          // è®¢å•IDï¼ˆSnowflakeï¼‰
	OrderNo        string      `orm:"order_no"         description:"è®¢å•ç¼–å·"`                                                                                     // è®¢å•ç¼–å·
	MemberId       uint64      `orm:"member_id"        description:"ä¸‹å•ä¼šå‘˜ID"`                                                                                   // ä¸‹å•ä¼šå‘˜ID
	CoachId        uint64      `orm:"coach_id"         description:"é™ªçŽ©å¸ˆID"`                                                                                      // é™ªçŽ©å¸ˆID
	ShopId         uint64      `orm:"shop_id"          description:"åº—é“ºID"`                                                                                         // åº—é“ºID
	GoodsId        uint64      `orm:"goods_id"         description:"å•†å“ID"`                                                                                         // å•†å“ID
	GoodsTitle     string      `orm:"goods_title"      description:"å•†å“åç§°ï¼ˆå†—ä½™ï¼‰"`                                                                         // å•†å“åç§°ï¼ˆå†—ä½™ï¼‰
	GoodsPrice     int64       `orm:"goods_price"      description:"å•†å“å•ä»·ï¼ˆåˆ†ï¼‰"`                                                                            // å•†å“å•ä»·ï¼ˆåˆ†ï¼‰
	Quantity       int         `orm:"quantity"         description:"æ•°é‡"`                                                                                           // æ•°é‡
	TotalAmount    int64       `orm:"total_amount"     description:"è®¢å•æ€»é¢ï¼ˆåˆ†ï¼‰"`                                                                            // è®¢å•æ€»é¢ï¼ˆåˆ†ï¼‰
	DiscountAmount int64       `orm:"discount_amount"  description:"ä¼šå‘˜æŠ˜æ‰£é‡‘é¢ï¼ˆåˆ†ï¼‰"`                                                                      // ä¼šå‘˜æŠ˜æ‰£é‡‘é¢ï¼ˆåˆ†ï¼‰
	CouponAmount   int64       `orm:"coupon_amount"    description:"ä¼˜æƒ åˆ¸æŠµæ‰£é‡‘é¢ï¼ˆåˆ†ï¼‰"`                                                                   // ä¼˜æƒ åˆ¸æŠµæ‰£é‡‘é¢ï¼ˆåˆ†ï¼‰
	PayAmount      int64       `orm:"pay_amount"       description:"å®žä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰"`                                                                            // å®žä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰
	CouponMemberId uint64      `orm:"coupon_member_id" description:"ä½¿ç”¨çš„ä¼˜æƒ åˆ¸é¢†å–è®°å½•ID"`                                                                 // ä½¿ç”¨çš„ä¼˜æƒ åˆ¸é¢†å–è®°å½•ID
	PayType        int         `orm:"pay_type"         description:"æ”¯ä»˜æ–¹å¼:0=æœªæ”¯ä»˜,1=å¾®ä¿¡æ”¯ä»˜,2=æ”¯ä»˜å®æ”¯ä»˜,3=ä½™é¢æ”¯ä»˜"`                         // æ”¯ä»˜æ–¹å¼:0=æœªæ”¯ä»˜,1=å¾®ä¿¡æ”¯ä»˜,2=æ”¯ä»˜å®æ”¯ä»˜,3=ä½™é¢æ”¯ä»˜
	OrderStatus    int         `orm:"order_status"     description:"è®¢å•çŠ¶æ€:0=å¾…æ”¯ä»˜,1=å·²æ”¯ä»˜,2=è¿›è¡Œä¸­,3=å·²å®Œæˆ,4=å·²å–æ¶ˆ,5=é€€æ¬¾ä¸­,6=å·²é€€æ¬¾"` // è®¢å•çŠ¶æ€:0=å¾…æ”¯ä»˜,1=å·²æ”¯ä»˜,2=è¿›è¡Œä¸­,3=å·²å®Œæˆ,4=å·²å–æ¶ˆ,5=é€€æ¬¾ä¸­,6=å·²é€€æ¬¾
	PayAt          *gtime.Time `orm:"pay_at"           description:"æ”¯ä»˜æ—¶é—´"`                                                                                     // æ”¯ä»˜æ—¶é—´
	StartAt        *gtime.Time `orm:"start_at"         description:"æœåŠ¡å¼€å§‹æ—¶é—´"`                                                                               // æœåŠ¡å¼€å§‹æ—¶é—´
	FinishAt       *gtime.Time `orm:"finish_at"        description:"æœåŠ¡å®Œæˆæ—¶é—´"`                                                                               // æœåŠ¡å®Œæˆæ—¶é—´
	CancelAt       *gtime.Time `orm:"cancel_at"        description:"å–æ¶ˆæ—¶é—´"`                                                                                     // å–æ¶ˆæ—¶é—´
	CancelReason   string      `orm:"cancel_reason"    description:"å–æ¶ˆåŽŸå› "`                                                                                     // å–æ¶ˆåŽŸå›
	Remark         string      `orm:"remark"           description:"è®¢å•å¤‡æ³¨"`                                                                                     // è®¢å•å¤‡æ³¨
	CreatedBy      uint64      `orm:"created_by"       description:"åˆ›å»ºäººID"`                                                                                      // åˆ›å»ºäººID
	DeptId         uint64      `orm:"dept_id"          description:"æ‰€å±žéƒ¨é—¨ID"`                                                                                   // æ‰€å±žéƒ¨é—¨ID
	CreatedAt      *gtime.Time `orm:"created_at"       description:"åˆ›å»ºæ—¶é—´"`                                                                                     // åˆ›å»ºæ—¶é—´
	UpdatedAt      *gtime.Time `orm:"updated_at"       description:"æ›´æ–°æ—¶é—´"`                                                                                     // æ›´æ–°æ—¶é—´
	DeletedAt      *gtime.Time `orm:"deleted_at"       description:"è½¯åˆ é™¤æ—¶é—´"`                                                                                  // è½¯åˆ é™¤æ—¶é—´
}
