package message

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
	service.RegisterMessage(New())
}

func New() *sMessage {
	return &sMessage{}
}

type sMessage struct{}

// Create 创建会员消息
func (s *sMessage) Create(ctx context.Context, in *model.MessageCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayMessage.Ctx(ctx).Data(g.Map{
		dao.PlayMessage.Columns().Id:        id,
		dao.PlayMessage.Columns().MemberId: in.MemberID,
		dao.PlayMessage.Columns().Title: in.Title,
		dao.PlayMessage.Columns().Content: in.Content,
		dao.PlayMessage.Columns().MsgType: in.MsgType,
		dao.PlayMessage.Columns().BizId: in.BizID,
		dao.PlayMessage.Columns().IsRead: in.IsRead,
		dao.PlayMessage.Columns().Status: in.Status,
		dao.PlayMessage.Columns().CreatedAt: gtime.Now(),
		dao.PlayMessage.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新会员消息
func (s *sMessage) Update(ctx context.Context, in *model.MessageUpdateInput) error {
	data := g.Map{
		dao.PlayMessage.Columns().MemberId: in.MemberID,
		dao.PlayMessage.Columns().Title: in.Title,
		dao.PlayMessage.Columns().Content: in.Content,
		dao.PlayMessage.Columns().MsgType: in.MsgType,
		dao.PlayMessage.Columns().BizId: in.BizID,
		dao.PlayMessage.Columns().IsRead: in.IsRead,
		dao.PlayMessage.Columns().Status: in.Status,
		dao.PlayMessage.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayMessage.Ctx(ctx).Where(dao.PlayMessage.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除会员消息
func (s *sMessage) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayMessage.Ctx(ctx).Where(dao.PlayMessage.Columns().Id, id).Data(g.Map{
		dao.PlayMessage.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// BatchDelete 批量软删除会员消息
func (s *sMessage) BatchDelete(ctx context.Context, ids []snowflake.JsonInt64) error {
	_, err := dao.PlayMessage.Ctx(ctx).WhereIn(dao.PlayMessage.Columns().Id, ids).Data(g.Map{
		dao.PlayMessage.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取会员消息详情
func (s *sMessage) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.MessageDetailOutput, err error) {
	out = &model.MessageDetailOutput{}
	err = dao.PlayMessage.Ctx(ctx).Where(dao.PlayMessage.Columns().Id, id).Where(dao.PlayMessage.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	// 查询接收者会员ID关联显示
	if out.MemberID != 0 {
		val, err := g.DB().Ctx(ctx).Model("play_member").Where("id", out.MemberID).Where("deleted_at", nil).Value("nickname")
		if err == nil {
			out.MemberNickname = val.String()
		}
	}
	return
}

// applyListFilter 应用列表通用过滤条件
func (s *sMessage) applyListFilter(ctx context.Context, in *model.MessageListInput) *gdb.Model {
	m := dao.PlayMessage.Ctx(ctx).Where(dao.PlayMessage.Columns().DeletedAt, nil)
	if in.Title != "" {
		m = m.WhereLike(dao.PlayMessage.Columns().Title, "%"+in.Title+"%")
	}
	if in.StartTime != "" {
		m = m.WhereGTE(dao.PlayMessage.Columns().CreatedAt, in.StartTime)
	}
	if in.EndTime != "" {
		m = m.WhereLTE(dao.PlayMessage.Columns().CreatedAt, in.EndTime)
	}
	return m
}

// fillRefFields 批量填充关联显示字段（避免 N+1 查询）
func (s *sMessage) fillRefFields(ctx context.Context, list []*model.MessageListOutput) {
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

// List 获取会员消息列表
func (s *sMessage) List(ctx context.Context, in *model.MessageListInput) (list []*model.MessageListOutput, total int, err error) {
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
		m = m.OrderAsc(dao.PlayMessage.Columns().Id)
	}
	err = m.Page(in.PageNum, in.PageSize).Scan(&list)
	if err != nil {
		return
	}
	s.fillRefFields(ctx, list)
	return
}
// Export 导出会员消息（不分页）
func (s *sMessage) Export(ctx context.Context, in *model.MessageListInput) (list []*model.MessageListOutput, err error) {
	m := s.applyListFilter(ctx, in)
	err = m.OrderAsc(dao.PlayMessage.Columns().Id).Limit(10000).Scan(&list)
	if err != nil {
		return
	}
	s.fillRefFields(ctx, list)
	return
}


