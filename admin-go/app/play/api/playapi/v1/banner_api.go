package v1

import "github.com/gogf/gf/v2/frame/g"

// ========== C端 Banner 公开接口 ==========

type BannerListReq struct {
	g.Meta `path:"/banner/list" method:"get" tags:"C端Banner" summary:"首页Banner列表"`
}

type BannerListItem struct {
	BannerID  string `json:"bannerId"  dc:"Banner ID"`
	Title     string `json:"title"     dc:"标题"`
	Image     string `json:"image"     dc:"图片URL"`
	LinkType  int    `json:"linkType"  dc:"跳转类型:1内页 2外链 3活动页 4商品页 5陪玩师页 6唤醒App"`
	LinkValue string `json:"linkValue" dc:"跳转值"`
}

type BannerListRes struct {
	g.Meta `mime:"application/json"`
	List   []BannerListItem `json:"list" dc:"Banner列表"`
}
