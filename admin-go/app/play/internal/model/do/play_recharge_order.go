// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayRechargeOrder is the golang structure of table play_recharge_order for DAO operations like Where/Data.
type PlayRechargeOrder struct {
	g.Meta         `orm:"table:play_recharge_order, do:true"`
	Id             any         // å……å€¼è®¢å•IDï¼ˆSnowflakeï¼‰
	OrderNo        any         // å……å€¼è®¢å•å·
	MemberId       any         // ä¼šå‘˜ID
	RechargePlanId any         // å……å€¼æ–¹æ¡ˆID
	Amount         any         // å……å€¼é‡‘é¢ï¼ˆåˆ†ï¼‰
	GiftAmount     any         // èµ é€é‡‘é¢ï¼ˆåˆ†ï¼‰
	PayType        any         // æ”¯ä»˜æ–¹å¼:1=å¾®ä¿¡æ”¯ä»˜,2=æ”¯ä»˜å®æ”¯ä»˜
	TradeNo        any         // ç¬¬ä¸‰æ–¹äº¤æ˜“å·
	PayStatus      any         // æ”¯ä»˜çŠ¶æ€:0=å¾…æ”¯ä»˜,1=æ”¯ä»˜æˆåŠŸ,2=æ”¯ä»˜å¤±è´¥
	PayAt          *gtime.Time // æ”¯ä»˜æ—¶é—´
	CreatedBy      any         // åˆ›å»ºäººID
	DeptId         any         // æ‰€å±žéƒ¨é—¨ID
	CreatedAt      *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt      *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt      *gtime.Time // è½¯åˆ é™¤æ—¶é—´
}
