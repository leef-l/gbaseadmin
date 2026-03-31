package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/upload/internal/model"
	"gbaseadmin/utility/snowflake"
)

// 确保 gtime 被引用
var _ = gtime.New

// Config API

// ConfigCreateReq 创建ä¸Šä¼ é…ç½®请求
type ConfigCreateReq struct {
	g.Meta `path:"/config/create" method:"post" tags:"ä¸Šä¼ é…ç½®" summary:"创建ä¸Šä¼ é…ç½®"`
	Name string `json:"name" v:"required#é…ç½®åç§°不能为空" dc:"é…ç½®åç§°"`
	Storage int `json:"storage"  dc:"å­˜å‚¨ç±»åž‹"`
	IsDefault int `json:"isDefault"  dc:"æ˜¯å¦é»˜è®¤"`
	LocalPath string `json:"localPath"  dc:"æœ¬åœ°å­˜å‚¨è·¯å¾„"`
	OssEndpoint string `json:"ossEndpoint"  dc:"OSS Endpoint"`
	OssBucket string `json:"ossBucket"  dc:"OSS Bucket"`
	OssAccessKey string `json:"ossAccessKey"  dc:"OSS AccessKey"`
	OssSecretKey string `json:"ossSecretKey"  dc:"OSS SecretKey"`
	CosRegion string `json:"cosRegion"  dc:"COS Region"`
	CosBucket string `json:"cosBucket"  dc:"COS Bucket"`
	CosSecretID snowflake.JsonInt64 `json:"cosSecretID"  dc:"COS SecretId"`
	CosSecretKey string `json:"cosSecretKey"  dc:"COS SecretKey"`
	MaxSize int `json:"maxSize"  dc:"æœ€å¤§æ–‡ä»¶å¤§å°(MB)"`
	Status int `json:"status"  dc:"çŠ¶æ€"`
}

// ConfigCreateRes 创建ä¸Šä¼ é…ç½®响应
type ConfigCreateRes struct {
	g.Meta `mime:"application/json"`
}

// ConfigUpdateReq 更新ä¸Šä¼ é…ç½®请求
type ConfigUpdateReq struct {
	g.Meta `path:"/config/update" method:"put" tags:"ä¸Šä¼ é…ç½®" summary:"更新ä¸Šä¼ é…ç½®"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"ä¸Šä¼ é…ç½®ID"`
	Name string `json:"name" dc:"é…ç½®åç§°"`
	Storage int `json:"storage" dc:"å­˜å‚¨ç±»åž‹"`
	IsDefault int `json:"isDefault" dc:"æ˜¯å¦é»˜è®¤"`
	LocalPath string `json:"localPath" dc:"æœ¬åœ°å­˜å‚¨è·¯å¾„"`
	OssEndpoint string `json:"ossEndpoint" dc:"OSS Endpoint"`
	OssBucket string `json:"ossBucket" dc:"OSS Bucket"`
	OssAccessKey string `json:"ossAccessKey" dc:"OSS AccessKey"`
	OssSecretKey string `json:"ossSecretKey" dc:"OSS SecretKey"`
	CosRegion string `json:"cosRegion" dc:"COS Region"`
	CosBucket string `json:"cosBucket" dc:"COS Bucket"`
	CosSecretID snowflake.JsonInt64 `json:"cosSecretID" dc:"COS SecretId"`
	CosSecretKey string `json:"cosSecretKey" dc:"COS SecretKey"`
	MaxSize int `json:"maxSize" dc:"æœ€å¤§æ–‡ä»¶å¤§å°(MB)"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// ConfigUpdateRes 更新ä¸Šä¼ é…ç½®响应
type ConfigUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// ConfigDeleteReq 删除ä¸Šä¼ é…ç½®请求
type ConfigDeleteReq struct {
	g.Meta `path:"/config/delete" method:"delete" tags:"ä¸Šä¼ é…ç½®" summary:"删除ä¸Šä¼ é…ç½®"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"ä¸Šä¼ é…ç½®ID"`
}

// ConfigDeleteRes 删除ä¸Šä¼ é…ç½®响应
type ConfigDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ConfigDetailReq 获取ä¸Šä¼ é…ç½®详情请求
type ConfigDetailReq struct {
	g.Meta `path:"/config/detail" method:"get" tags:"ä¸Šä¼ é…ç½®" summary:"获取ä¸Šä¼ é…ç½®详情"`
	ID     snowflake.JsonInt64 `json:"id" v:"required#ID不能为空" dc:"ä¸Šä¼ é…ç½®ID"`
}

// ConfigDetailRes 获取ä¸Šä¼ é…ç½®详情响应
type ConfigDetailRes struct {
	g.Meta `mime:"application/json"`
	*model.ConfigDetailOutput
}

// ConfigListReq 获取ä¸Šä¼ é…ç½®列表请求
type ConfigListReq struct {
	g.Meta   `path:"/config/list" method:"get" tags:"ä¸Šä¼ é…ç½®" summary:"获取ä¸Šä¼ é…ç½®列表"`
	PageNum  int `json:"pageNum" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"10" dc:"每页数量"`
	Storage int `json:"storage" dc:"å­˜å‚¨ç±»åž‹"`
	IsDefault int `json:"isDefault" dc:"æ˜¯å¦é»˜è®¤"`
	Status int `json:"status" dc:"çŠ¶æ€"`
}

// ConfigListRes 获取ä¸Šä¼ é…ç½®列表响应
type ConfigListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.ConfigListOutput `json:"list" dc:"列表数据"`
	Total  int                               `json:"total" dc:"总数"`
}

