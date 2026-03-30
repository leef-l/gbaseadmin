package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MemberInfoReq 获取个人信息
type MemberInfoReq struct {
	g.Meta `path:"/member/info" method:"get" tags:"C端会员" summary:"获取个人信息"`
}

type MemberInfoRes struct {
	g.Meta      `mime:"application/json"`
	MemberID    string `json:"memberId" dc:"会员ID"`
	Phone       string `json:"phone" dc:"手机号（脱敏）"`
	Nickname    string `json:"nickname" dc:"昵称"`
	Avatar      string `json:"avatar" dc:"头像URL"`
	Gender      int    `json:"gender" dc:"性别:0=未知,1=男,2=女"`
	Balance     int64  `json:"balance" dc:"余额(分)"`
	LevelTitle  string `json:"levelTitle" dc:"等级名称"`
	LevelIcon   string `json:"levelIcon" dc:"等级图标"`
	Discount    int    `json:"discount" dc:"会员折扣"`
	Exp         int    `json:"exp" dc:"当前经验值"`
	IsCoach     int    `json:"isCoach" dc:"是否陪玩师:0=否,1=是"`
	CoachID     string `json:"coachId" dc:"陪玩师ID"`
	CurrentRole string `json:"currentRole" dc:"当前身份:member/coach"`
	WxBound     bool   `json:"wxBound" dc:"是否绑定微信"`
	AlipayBound bool   `json:"alipayBound" dc:"是否绑定支付宝"`
}

// MemberUpdateReq 编辑资料
type MemberUpdateReq struct {
	g.Meta   `path:"/member/update" method:"put" tags:"C端会员" summary:"编辑资料"`
	Nickname string `json:"nickname" v:"max-length:20#昵称最多20个字符" dc:"昵称"`
	Avatar   string `json:"avatar" dc:"头像URL"`
	Gender   *int   `json:"gender" v:"in:0,1,2#性别值不合法" dc:"性别:0=未知,1=男,2=女"`
}

type MemberUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// MemberSwitchRoleReq 切换身份
type MemberSwitchRoleReq struct {
	g.Meta `path:"/member/switch_role" method:"post" tags:"C端会员" summary:"切换身份"`
	Role   string `json:"role" v:"required|in:member,coach#身份不能为空|身份值必须为member或coach" dc:"目标身份:member/coach"`
}

type MemberSwitchRoleRes struct {
	g.Meta `mime:"application/json"`
	Token  string `json:"token" dc:"新访问令牌"`
}

// MemberBalanceLogReq 余额流水列表
type MemberBalanceLogReq struct {
	g.Meta   `path:"/member/balance_log" method:"get" tags:"C端会员" summary:"余额流水列表"`
	Type     string `json:"type" dc:"类型筛选:income=收入,expense=支出，空=全部"`
	Page     int    `json:"page" v:"min:1" dc:"页码" d:"1"`
	PageSize int    `json:"pageSize" v:"max:50" dc:"每页条数" d:"20"`
}

type MemberBalanceLogRes struct {
	g.Meta `mime:"application/json"`
	Total  int                    `json:"total" dc:"总数"`
	List   []MemberBalanceLogItem `json:"list" dc:"流水列表"`
}

type MemberBalanceLogItem struct {
	ID        string `json:"id" dc:"流水ID"`
	Type      string `json:"type" dc:"类型:recharge=充值,pay=支付,refund=退款,income=收入,withdraw=提现"`
	Amount    int64  `json:"amount" dc:"金额(分)，正数=收入，负数=支出"`
	Balance   int64  `json:"balance" dc:"变动后余额(分)"`
	Title     string `json:"title" dc:"标题"`
	Remark    string `json:"remark" dc:"备注"`
	CreatedAt string `json:"createdAt" dc:"创建时间"`
}
