// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayProfitLog is the golang structure for table play_profit_log.
type PlayProfitLog struct {
	Id             uint64      `orm:"id"              description:"æµæ°´IDï¼ˆSnowflakeï¼‰"`              // æµæ°´IDï¼ˆSnowflakeï¼‰
	OrderId        uint64      `orm:"order_id"        description:"è®¢å•ID"`                             // è®¢å•ID
	OrderNo        string      `orm:"order_no"        description:"è®¢å•ç¼–å·"`                         // è®¢å•ç¼–å·
	PayAmount      int64       `orm:"pay_amount"      description:"å®žä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰"`                // å®žä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰
	CoachId        uint64      `orm:"coach_id"        description:"é™ªçŽ©å¸ˆID"`                          // é™ªçŽ©å¸ˆID
	ShopId         uint64      `orm:"shop_id"         description:"åº—é“ºID"`                             // åº—é“ºID
	PlatformRate   int         `orm:"platform_rate"   description:"å¹³å°æŠ½æˆæ¯”ä¾‹"`                   // å¹³å°æŠ½æˆæ¯”ä¾‹
	PlatformAmount int64       `orm:"platform_amount" description:"å¹³å°æŠ½æˆé‡‘é¢ï¼ˆåˆ†ï¼‰"`          // å¹³å°æŠ½æˆé‡‘é¢ï¼ˆåˆ†ï¼‰
	ShopRate       int         `orm:"shop_rate"       description:"åº—é“ºæŠ½æˆæ¯”ä¾‹"`                   // åº—é“ºæŠ½æˆæ¯”ä¾‹
	ShopAmount     int64       `orm:"shop_amount"     description:"åº—é“ºæŠ½æˆé‡‘é¢ï¼ˆåˆ†ï¼‰"`          // åº—é“ºæŠ½æˆé‡‘é¢ï¼ˆåˆ†ï¼‰
	CoachAmount    int64       `orm:"coach_amount"    description:"é™ªçŽ©å¸ˆæ”¶å…¥ï¼ˆåˆ†ï¼‰"`             // é™ªçŽ©å¸ˆæ”¶å…¥ï¼ˆåˆ†ï¼‰
	SettleStatus   int         `orm:"settle_status"   description:"ç»“ç®—çŠ¶æ€:0=å¾…ç»“ç®—,1=å·²ç»“ç®—"` // ç»“ç®—çŠ¶æ€:0=å¾…ç»“ç®—,1=å·²ç»“ç®—
	SettleAt       *gtime.Time `orm:"settle_at"       description:"ç»“ç®—æ—¶é—´"`                         // ç»“ç®—æ—¶é—´
	CreatedBy      uint64      `orm:"created_by"      description:"åˆ›å»ºäººID"`                          // åˆ›å»ºäººID
	DeptId         uint64      `orm:"dept_id"         description:"æ‰€å±žéƒ¨é—¨ID"`                       // æ‰€å±žéƒ¨é—¨ID
	CreatedAt      *gtime.Time `orm:"created_at"      description:"åˆ›å»ºæ—¶é—´"`                         // åˆ›å»ºæ—¶é—´
	UpdatedAt      *gtime.Time `orm:"updated_at"      description:"æ›´æ–°æ—¶é—´"`                         // æ›´æ–°æ—¶é—´
	DeletedAt      *gtime.Time `orm:"deleted_at"      description:"è½¯åˆ é™¤æ—¶é—´"`                      // è½¯åˆ é™¤æ—¶é—´
}
