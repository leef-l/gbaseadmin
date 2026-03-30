// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayBalanceLog is the golang structure for table play_balance_log.
type PlayBalanceLog struct {
	Id            uint64      `orm:"id"             description:"æµæ°´IDï¼ˆSnowflakeï¼‰"`                                         // æµæ°´IDï¼ˆSnowflakeï¼‰
	MemberId      uint64      `orm:"member_id"      description:"ä¼šå‘˜ID"`                                                        // ä¼šå‘˜ID
	BizType       int         `orm:"biz_type"       description:"ä¸šåŠ¡ç±»åž‹:1=å……å€¼,2=æ¶ˆè´¹,3=é€€æ¬¾,4=æ´»åŠ¨èµ é€,5=æçŽ°"` // ä¸šåŠ¡ç±»åž‹:1=å……å€¼,2=æ¶ˆè´¹,3=é€€æ¬¾,4=æ´»åŠ¨èµ é€,5=æçŽ°
	BizId         uint64      `orm:"biz_id"         description:"å…³è”ä¸šåŠ¡ID"`                                                  // å…³è”ä¸šåŠ¡ID
	ChangeAmount  int64       `orm:"change_amount"  description:"å˜åŠ¨é‡‘é¢ï¼ˆåˆ†ï¼‰"`                                           // å˜åŠ¨é‡‘é¢ï¼ˆåˆ†ï¼‰
	BeforeBalance int64       `orm:"before_balance" description:"å˜åŠ¨å‰ä½™é¢ï¼ˆåˆ†ï¼‰"`                                        // å˜åŠ¨å‰ä½™é¢ï¼ˆåˆ†ï¼‰
	AfterBalance  int64       `orm:"after_balance"  description:"å˜åŠ¨åŽä½™é¢ï¼ˆåˆ†ï¼‰"`                                        // å˜åŠ¨åŽä½™é¢ï¼ˆåˆ†ï¼‰
	Remark        string      `orm:"remark"         description:"å¤‡æ³¨è¯´æ˜Ž"`                                                    // å¤‡æ³¨è¯´æ˜Ž
	CreatedBy     uint64      `orm:"created_by"     description:"åˆ›å»ºäººID"`                                                     // åˆ›å»ºäººID
	DeptId        uint64      `orm:"dept_id"        description:"æ‰€å±žéƒ¨é—¨ID"`                                                  // æ‰€å±žéƒ¨é—¨ID
	CreatedAt     *gtime.Time `orm:"created_at"     description:"åˆ›å»ºæ—¶é—´"`                                                    // åˆ›å»ºæ—¶é—´
	UpdatedAt     *gtime.Time `orm:"updated_at"     description:"æ›´æ–°æ—¶é—´"`                                                    // æ›´æ–°æ—¶é—´
	DeletedAt     *gtime.Time `orm:"deleted_at"     description:"è½¯åˆ é™¤æ—¶é—´"`                                                 // è½¯åˆ é™¤æ—¶é—´
}
