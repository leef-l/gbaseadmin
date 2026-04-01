package playapi

import (
	"context"
	"strconv"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	v1 "gbaseadmin/app/play/api/playapi/v1"
	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/service"
	"gbaseadmin/utility/snowflake"
)

type sPlayapiWithdraw struct{}

func init() {
	service.RegisterPlayapiWithdraw(&sPlayapiWithdraw{})
}

// Withdraw 申请提现
func (s *sPlayapiWithdraw) Withdraw(ctx context.Context, coachID int64, memberID int64, amount int64) (withdrawID string, err error) {
	cc := dao.PlayCoach.Columns()
	wc := dao.PlayWithdraw.Columns()

	err = dao.PlayWithdraw.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 1. 查询陪玩师可提现余额并加锁
		coachRow, e := dao.PlayCoach.Ctx(ctx).
			Where(cc.Id, coachID).
			Where(cc.Status, 1).
			LockUpdate().
			One()
		if e != nil {
			return e
		}
		if coachRow.IsEmpty() {
			return gerror.New("陪玩师信息不存在")
		}

		incomeBalance := coachRow[cc.IncomeBalance].Int64()
		if incomeBalance < amount {
			return gerror.New("可提现余额不足")
		}

		afterBalance := incomeBalance - amount

		// 2. 扣减可提现余额
		_, e = dao.PlayCoach.Ctx(ctx).
			Where(cc.Id, coachID).
			Data(g.Map{
				cc.IncomeBalance: afterBalance,
				cc.UpdatedAt:     gtime.Now(),
			}).Update()
		if e != nil {
			return e
		}

		// 3. 创建提现记录
		id := snowflake.Generate()
		_, e = dao.PlayWithdraw.Ctx(ctx).Data(g.Map{
			wc.Id:        id,
			wc.CoachId:   coachID,
			wc.MemberId:  memberID,
			wc.Amount:    amount,
			wc.Status:    0, // 待审核
			wc.Reason:    "",
			wc.CreatedAt: gtime.Now(),
			wc.UpdatedAt: gtime.Now(),
		}).Insert()
		if e != nil {
			return e
		}

		withdrawID = strconv.FormatInt(int64(id), 10)
		return nil
	})
	return
}

// WithdrawList 提现记录列表
func (s *sPlayapiWithdraw) WithdrawList(ctx context.Context, coachID int64, req *v1.CoachWithdrawListApiReq) (list []v1.CoachWithdrawItem, total int, err error) {
	wc := dao.PlayWithdraw.Columns()

	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 20
	}

	m := dao.PlayWithdraw.Ctx(ctx).
		Where(wc.CoachId, coachID).
		Where(wc.DeletedAt, nil).
		OrderDesc(wc.CreatedAt)

	total, err = m.Count()
	if err != nil {
		return
	}

	var records []struct {
		Id        int64  `json:"id"`
		Amount    int64  `json:"amount"`
		Status    int    `json:"status"`
		Reason    string `json:"reason"`
		CreatedAt string `json:"created_at"`
	}
	err = m.Page(page, pageSize).Scan(&records)
	if err != nil {
		return
	}

	list = make([]v1.CoachWithdrawItem, 0, len(records))
	for _, r := range records {
		list = append(list, v1.CoachWithdrawItem{
			Id:        strconv.FormatInt(r.Id, 10),
			Amount:    r.Amount,
			Status:    r.Status,
			Reason:    r.Reason,
			CreatedAt: r.CreatedAt,
		})
	}
	return
}
