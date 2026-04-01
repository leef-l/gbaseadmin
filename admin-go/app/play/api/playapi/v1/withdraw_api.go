package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CoachWithdrawApiReq 申请提现
type CoachWithdrawApiReq struct {
	g.Meta `path:"/coach/withdraw" method:"post" tags:"C端陪玩师" summary:"申请提现"`
	Amount int64 `json:"amount" v:"required|min:100#提现金额不能为空|最低提现1元" dc:"提现金额(分)"`
}

type CoachWithdrawApiRes struct {
	WithdrawId string `json:"withdrawId" dc:"提现单号"`
}

// CoachWithdrawListApiReq 提现记录列表
type CoachWithdrawListApiReq struct {
	g.Meta   `path:"/coach/withdraw_list" method:"get" tags:"C端陪玩师" summary:"提现记录"`
	Page     int `json:"page" d:"1" dc:"页码"`
	PageSize int `json:"pageSize" d:"20" dc:"每页数量"`
}

type CoachWithdrawListApiRes struct {
	List  []CoachWithdrawItem `json:"list"`
	Total int                 `json:"total"`
}

type CoachWithdrawItem struct {
	Id        string `json:"id"`
	Amount    int64  `json:"amount"`
	Status    int    `json:"status" dc:"0=待审核 1=已打款 2=已拒绝"`
	Reason    string `json:"reason" dc:"拒绝原因"`
	CreatedAt string `json:"createdAt"`
}
