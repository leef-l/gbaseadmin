// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayBalanceLog is the golang structure of table play_balance_log for DAO operations like Where/Data.
type PlayBalanceLog struct {
	g.Meta        `orm:"table:play_balance_log, do:true"`
	Id            any         // æµæ°´IDï¼ˆSnowflakeï¼‰
	MemberId      any         // ä¼šå‘˜ID
	BizType       any         // ä¸šåŠ¡ç±»åž‹:1=å……å€¼,2=æ¶ˆè´¹,3=é€€æ¬¾,4=æ´»åŠ¨èµ é€,5=æçŽ°
	BizId         any         // å…³è”ä¸šåŠ¡ID
	ChangeAmount  any         // å˜åŠ¨é‡‘é¢ï¼ˆåˆ†ï¼‰
	BeforeBalance any         // å˜åŠ¨å‰ä½™é¢ï¼ˆåˆ†ï¼‰
	AfterBalance  any         // å˜åŠ¨åŽä½™é¢ï¼ˆåˆ†ï¼‰
	Remark        any         // å¤‡æ³¨è¯´æ˜Ž
	CreatedBy     any         // åˆ›å»ºäººID
	DeptId        any         // æ‰€å±žéƒ¨é—¨ID
	CreatedAt     *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt     *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt     *gtime.Time // è½¯åˆ é™¤æ—¶é—´
}
