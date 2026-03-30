// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlayActivityStep is the golang structure for table play_activity_step.
type PlayActivityStep struct {
	Id          uint64      `orm:"id"           description:"æ­¥éª¤IDï¼ˆSnowflakeï¼‰"` // æ­¥éª¤IDï¼ˆSnowflakeï¼‰
	ActivityId  uint64      `orm:"activity_id"  description:"æ´»åŠ¨ID"`                // æ´»åŠ¨ID
	StepNum     int         `orm:"step_num"     description:"æ­¥éª¤åºå·"`            // æ­¥éª¤åºå·
	Title       string      `orm:"title"        description:"æ­¥éª¤æ ‡é¢˜"`            // æ­¥éª¤æ ‡é¢˜
	DescContent string      `orm:"desc_content" description:"æ­¥éª¤è¯´æ˜Ž"`            // æ­¥éª¤è¯´æ˜Ž
	StepImage   string      `orm:"step_image"   description:"æ­¥éª¤ç¤ºä¾‹å›¾ç‰‡"`      // æ­¥éª¤ç¤ºä¾‹å›¾ç‰‡
	Sort        int         `orm:"sort"         description:"æŽ’åº"`                  // æŽ’åº
	CreatedBy   uint64      `orm:"created_by"   description:"åˆ›å»ºäººID"`             // åˆ›å»ºäººID
	DeptId      uint64      `orm:"dept_id"      description:"æ‰€å±žéƒ¨é—¨ID"`          // æ‰€å±žéƒ¨é—¨ID
	CreatedAt   *gtime.Time `orm:"created_at"   description:"åˆ›å»ºæ—¶é—´"`            // åˆ›å»ºæ—¶é—´
	UpdatedAt   *gtime.Time `orm:"updated_at"   description:"æ›´æ–°æ—¶é—´"`            // æ›´æ–°æ—¶é—´
	DeletedAt   *gtime.Time `orm:"deleted_at"   description:"è½¯åˆ é™¤æ—¶é—´"`         // è½¯åˆ é™¤æ—¶é—´
}
