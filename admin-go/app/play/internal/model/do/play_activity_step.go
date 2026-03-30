// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayActivityStep is the golang structure of table play_activity_step for DAO operations like Where/Data.
type PlayActivityStep struct {
	g.Meta      `orm:"table:play_activity_step, do:true"`
	Id          any         // æ­¥éª¤IDï¼ˆSnowflakeï¼‰
	ActivityId  any         // æ´»åŠ¨ID
	StepNum     any         // æ­¥éª¤åºå·
	Title       any         // æ­¥éª¤æ ‡é¢˜
	DescContent any         // æ­¥éª¤è¯´æ˜Ž
	StepImage   any         // æ­¥éª¤ç¤ºä¾‹å›¾ç‰‡
	Sort        any         // æŽ’åº
	CreatedBy   any         // åˆ›å»ºäººID
	DeptId      any         // æ‰€å±žéƒ¨é—¨ID
	CreatedAt   *gtime.Time // åˆ›å»ºæ—¶é—´
	UpdatedAt   *gtime.Time // æ›´æ–°æ—¶é—´
	DeletedAt   *gtime.Time // è½¯åˆ é™¤æ—¶é—´
}
