// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UploadConfig is the golang structure for table upload_config.
type UploadConfig struct {
	Id           uint64      `orm:"id"             description:"ID"`                                                  // ID
	Name         string      `orm:"name"           description:"é…ç½®åç§°"`                                        // é…ç½®åç§°
	Storage      int         `orm:"storage"        description:"å­˜å‚¨ç±»åž‹:1=æœ¬åœ°,2=é˜¿é‡Œäº‘OSS,3=è…¾è®¯äº‘COS"` // å­˜å‚¨ç±»åž‹:1=æœ¬åœ°,2=é˜¿é‡Œäº‘OSS,3=è…¾è®¯äº‘COS
	IsDefault    int         `orm:"is_default"     description:"æ˜¯å¦é»˜è®¤:0=å¦,1=æ˜¯"`                            // æ˜¯å¦é»˜è®¤:0=å¦,1=æ˜¯
	LocalPath    string      `orm:"local_path"     description:"æœ¬åœ°å­˜å‚¨è·¯å¾„"`                                  // æœ¬åœ°å­˜å‚¨è·¯å¾„
	OssEndpoint  string      `orm:"oss_endpoint"   description:"OSS Endpoint"`                                        // OSS Endpoint
	OssBucket    string      `orm:"oss_bucket"     description:"OSS Bucket"`                                          // OSS Bucket
	OssAccessKey string      `orm:"oss_access_key" description:"OSS AccessKey"`                                       // OSS AccessKey
	OssSecretKey string      `orm:"oss_secret_key" description:"OSS SecretKey"`                                       // OSS SecretKey
	CosRegion    string      `orm:"cos_region"     description:"COS Region"`                                          // COS Region
	CosBucket    string      `orm:"cos_bucket"     description:"COS Bucket"`                                          // COS Bucket
	CosSecretId  string      `orm:"cos_secret_id"  description:"COS SecretId"`                                        // COS SecretId
	CosSecretKey string      `orm:"cos_secret_key" description:"COS SecretKey"`                                       // COS SecretKey
	MaxSize      int         `orm:"max_size"       description:"æœ€å¤§æ–‡ä»¶å¤§å°(MB)"`                              // æœ€å¤§æ–‡ä»¶å¤§å°(MB)
	Status       int         `orm:"status"         description:"çŠ¶æ€:0=ç¦ç”¨,1=å¯ç”¨"`                            // çŠ¶æ€:0=ç¦ç”¨,1=å¯ç”¨
	CreatedAt    *gtime.Time `orm:"created_at"     description:"åˆ›å»ºæ—¶é—´"`                                        // åˆ›å»ºæ—¶é—´
	UpdatedAt    *gtime.Time `orm:"updated_at"     description:"æ›´æ–°æ—¶é—´"`                                        // æ›´æ–°æ—¶é—´
	DeletedAt    *gtime.Time `orm:"deleted_at"     description:"åˆ é™¤æ—¶é—´"`                                        // åˆ é™¤æ—¶é—´
	CreatedBy    uint64      `orm:"created_by"     description:"åˆ›å»ºäºº"`                                           // åˆ›å»ºäºº
	DeptId       uint64      `orm:"dept_id"        description:"éƒ¨é—¨ID"`                                            // éƒ¨é—¨ID
}
