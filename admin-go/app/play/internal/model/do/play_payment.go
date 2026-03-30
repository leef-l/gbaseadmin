// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayPayment is the golang structure of table play_payment for DAO operations like Where/Data.
type PlayPayment struct {
	g.Meta          `orm:"table:play_payment, do:true"`
	Id              any         // æ”¯ä»˜è®°å½•IDï¼ˆSnowflakeï¼‰
	OrderId         any         // è®¢å•ID
	MemberId        any         // ä¼šå‘˜ID
	PaymentNo       any         // æ”¯ä»˜æµæ°´å·
	TradeNo         any         // ç¬¬ä¸‰æ–¹äº¤æ˜“å·
	PayType         any         // æ”¯ä»˜æ–¹å¼:1=å¾®ä¿¡æ”¯ä»˜,2=æ”¯ä»˜å®æ”¯ä»˜,3=ä½™é¢æ”¯ä»˜
	PayAmount       any         // æ”¯ä»˜é‡‘é¢ï¼ˆåˆ†ï¼‰
	PayStatus       any         // æ”¯ä»˜çŠ¶æ€:0=å¾…æ”¯ä»˜,1=æ”¯ä»˜æˆåŠŸ,2=æ”¯ä»˜å¤±è´¥,3=å·²é€€æ¬¾
	PayAt           *gtime.Time // æ”¯ä»˜æˆåŠŸæ—¶é—´
	RefundAt        *gtime.Time // é€€æ¬¾æ—¶é—´
	RefundAmount    any         // é€€æ¬¾é‡‘é¢ï¼ˆåˆ†ï¼‰
	CallbackContent any         // å›žè°ƒæŠ¥æ–‡
	CreatedBy       any         // åˆ›å»ºäººID
	DeptId          any         // æ‰€å±žéƒ¨é—¨ID
	CreatedAt       *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt       *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt       *gtime.Time // è½¯åˆ é™¤æ—¶é—´
}
