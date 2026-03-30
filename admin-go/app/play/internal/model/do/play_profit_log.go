// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayProfitLog is the golang structure of table play_profit_log for DAO operations like Where/Data.
type PlayProfitLog struct {
	g.Meta         `orm:"table:play_profit_log, do:true"`
	Id             any         // æµæ°´IDï¼ˆSnowflakeï¼‰
	OrderId        any         // è®¢å•ID
	OrderNo        any         // è®¢å•ç¼–å·
	PayAmount      any         // å®žä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰
	CoachId        any         // é™ªçŽ©å¸ˆID
	ShopId         any         // åº—é“ºID
	PlatformRate   any         // å¹³å°æŠ½æˆæ¯”ä¾‹
	PlatformAmount any         // å¹³å°æŠ½æˆé‡‘é¢ï¼ˆåˆ†ï¼‰
	ShopRate       any         // åº—é“ºæŠ½æˆæ¯”ä¾‹
	ShopAmount     any         // åº—é“ºæŠ½æˆé‡‘é¢ï¼ˆåˆ†ï¼‰
	CoachAmount    any         // é™ªçŽ©å¸ˆæ”¶å…¥ï¼ˆåˆ†ï¼‰
	SettleStatus   any         // ç»“ç®—çŠ¶æ€:0=å¾…ç»“ç®—,1=å·²ç»“ç®—
	SettleAt       *gtime.Time // ç»“ç®—æ—¶é—´
	CreatedBy      any         // åˆ›å»ºäººID
	DeptId         any         // æ‰€å±žéƒ¨é—¨ID
	CreatedAt      *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt      *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt      *gtime.Time // è½¯åˆ é™¤æ—¶é—´
}
