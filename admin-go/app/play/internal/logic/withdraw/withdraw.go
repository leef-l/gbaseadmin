package withdraw

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
	"gbaseadmin/utility/snowflake"
)

func init() {
	service.RegisterWithdraw(New())
}

func New() *sWithdraw {
	return &sWithdraw{}
}

type sWithdraw struct{}

// Create 创建陪玩师提现记录
func (s *sWithdraw) Create(ctx context.Context, in *model.WithdrawCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayWithdraw.Ctx(ctx).Data(g.Map{
		dao.PlayWithdraw.Columns().Id:        id,
		dao.PlayWithdraw.Columns().CoachId: in.CoachID,
		dao.PlayWithdraw.Columns().MemberId: in.MemberID,
		dao.PlayWithdraw.Columns().Amount: in.Amount,
		dao.PlayWithdraw.Columns().Status: in.Status,
		dao.PlayWithdraw.Columns().Reason: in.Reason,
		dao.PlayWithdraw.Columns().AuditedAt: in.AuditedAt,
		dao.PlayWithdraw.Columns().CreatedAt: gtime.Now(),
		dao.PlayWithdraw.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新陪玩师提现记录
func (s *sWithdraw) Update(ctx context.Context, in *model.WithdrawUpdateInput) error {
	data := g.Map{
		dao.PlayWithdraw.Columns().CoachId: in.CoachID,
		dao.PlayWithdraw.Columns().MemberId: in.MemberID,
		dao.PlayWithdraw.Columns().Amount: in.Amount,
		dao.PlayWithdraw.Columns().Status: in.Status,
		dao.PlayWithdraw.Columns().Reason: in.Reason,
		dao.PlayWithdraw.Columns().AuditedAt: in.AuditedAt,
		dao.PlayWithdraw.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayWithdraw.Ctx(ctx).Where(dao.PlayWithdraw.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除陪玩师提现记录
func (s *sWithdraw) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayWithdraw.Ctx(ctx).Where(dao.PlayWithdraw.Columns().Id, id).Data(g.Map{
		dao.PlayWithdraw.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// BatchDelete 批量软删除陪玩师提现记录
func (s *sWithdraw) BatchDelete(ctx context.Context, ids []snowflake.JsonInt64) error {
	_, err := dao.PlayWithdraw.Ctx(ctx).WhereIn(dao.PlayWithdraw.Columns().Id, ids).Data(g.Map{
		dao.PlayWithdraw.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取陪玩师提现记录详情
func (s *sWithdraw) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.WithdrawDetailOutput, err error) {
	out = &model.WithdrawDetailOutput{}
	err = dao.PlayWithdraw.Ctx(ctx).Where(dao.PlayWithdraw.Columns().Id, id).Where(dao.PlayWithdraw.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	// 查询陪玩师ID关联显示
	if out.CoachID != 0 {
		val, err := g.DB().Ctx(ctx).Model("play_coach").Where("id", out.CoachID).Where("deleted_at", nil).Value("real_name")
		if err == nil {
			out.CoachRealName = val.String()
		}
	}
	// 查询会员ID关联显示
	if out.MemberID != 0 {
		val, err := g.DB().Ctx(ctx).Model("play_member").Where("id", out.MemberID).Where("deleted_at", nil).Value("nickname")
		if err == nil {
			out.MemberNickname = val.String()
		}
	}
	return
}

// applyListFilter 应用列表通用过滤条件
func (s *sWithdraw) applyListFilter(ctx context.Context, in *model.WithdrawListInput) *gdb.Model {
	m := dao.PlayWithdraw.Ctx(ctx).Where(dao.PlayWithdraw.Columns().DeletedAt, nil)
	if in.StartTime != "" {
		m = m.WhereGTE(dao.PlayWithdraw.Columns().CreatedAt, in.StartTime)
	}
	if in.EndTime != "" {
		m = m.WhereLTE(dao.PlayWithdraw.Columns().CreatedAt, in.EndTime)
	}
	return m
}

// fillRefFields 批量填充关联显示字段（避免 N+1 查询）
func (s *sWithdraw) fillRefFields(ctx context.Context, list []*model.WithdrawListOutput) {
	{
		idSet := make(map[int64]struct{})
		for _, item := range list {
			if item.CoachID != 0 {
				idSet[int64(item.CoachID)] = struct{}{}
			}
		}
		if len(idSet) > 0 {
			ids := make([]int64, 0, len(idSet))
			for id := range idSet {
				ids = append(ids, id)
			}
			rows, err := g.DB().Ctx(ctx).Model("play_coach").
				Fields("id", "real_name").
				Where("deleted_at", nil).
				WhereIn("id", ids).
				All()
			if err == nil {
				refMap := make(map[int64]string, len(rows))
				for _, row := range rows {
					refMap[row["id"].Int64()] = row["real_name"].String()
				}
				for _, item := range list {
					if val, ok := refMap[int64(item.CoachID)]; ok {
						item.CoachRealName = val
					}
				}
			}
		}
	}
	{
		idSet := make(map[int64]struct{})
		for _, item := range list {
			if item.MemberID != 0 {
				idSet[int64(item.MemberID)] = struct{}{}
			}
		}
		if len(idSet) > 0 {
			ids := make([]int64, 0, len(idSet))
			for id := range idSet {
				ids = append(ids, id)
			}
			rows, err := g.DB().Ctx(ctx).Model("play_member").
				Fields("id", "nickname").
				Where("deleted_at", nil).
				WhereIn("id", ids).
				All()
			if err == nil {
				refMap := make(map[int64]string, len(rows))
				for _, row := range rows {
					refMap[row["id"].Int64()] = row["nickname"].String()
				}
				for _, item := range list {
					if val, ok := refMap[int64(item.MemberID)]; ok {
						item.MemberNickname = val
					}
				}
			}
		}
	}
}

// List 获取陪玩师提现记录列表
func (s *sWithdraw) List(ctx context.Context, in *model.WithdrawListInput) (list []*model.WithdrawListOutput, total int, err error) {
	m := s.applyListFilter(ctx, in)
	total, err = m.Count()
	if err != nil {
		return
	}
	// 动态排序
	if in.OrderBy != "" {
		if in.OrderDir == "desc" {
			m = m.OrderDesc(in.OrderBy)
		} else {
			m = m.OrderAsc(in.OrderBy)
		}
	} else {
		m = m.OrderAsc(dao.PlayWithdraw.Columns().Id)
	}
	err = m.Page(in.PageNum, in.PageSize).Scan(&list)
	if err != nil {
		return
	}
	s.fillRefFields(ctx, list)
	return
}
// Export 导出陪玩师提现记录（不分页）
func (s *sWithdraw) Export(ctx context.Context, in *model.WithdrawListInput) (list []*model.WithdrawListOutput, err error) {
	m := s.applyListFilter(ctx, in)
	err = m.OrderAsc(dao.PlayWithdraw.Columns().Id).Limit(10000).Scan(&list)
	if err != nil {
		return
	}
	s.fillRefFields(ctx, list)
	return
}


