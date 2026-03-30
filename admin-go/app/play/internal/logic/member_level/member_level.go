package member_level

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
	"gbaseadmin/utility/snowflake"
)

func init() {
	service.RegisterMemberLevel(New())
}

func New() *sMemberLevel {
	return &sMemberLevel{}
}

type sMemberLevel struct{}

// Create 创建会员等级表
func (s *sMemberLevel) Create(ctx context.Context, in *model.MemberLevelCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayMemberLevel.Ctx(ctx).Data(g.Map{
		dao.PlayMemberLevel.Columns().Id:        id,
		dao.PlayMemberLevel.Columns().Title: in.Title,
		dao.PlayMemberLevel.Columns().Level: in.Level,
		dao.PlayMemberLevel.Columns().Icon: in.Icon,
		dao.PlayMemberLevel.Columns().MinExp: in.MinExp,
		dao.PlayMemberLevel.Columns().Discount: in.Discount,
		dao.PlayMemberLevel.Columns().Sort: in.Sort,
		dao.PlayMemberLevel.Columns().Status: in.Status,
		dao.PlayMemberLevel.Columns().CreatedAt: gtime.Now(),
		dao.PlayMemberLevel.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新会员等级表
func (s *sMemberLevel) Update(ctx context.Context, in *model.MemberLevelUpdateInput) error {
	data := g.Map{
		dao.PlayMemberLevel.Columns().Title: in.Title,
		dao.PlayMemberLevel.Columns().Level: in.Level,
		dao.PlayMemberLevel.Columns().Icon: in.Icon,
		dao.PlayMemberLevel.Columns().MinExp: in.MinExp,
		dao.PlayMemberLevel.Columns().Discount: in.Discount,
		dao.PlayMemberLevel.Columns().Sort: in.Sort,
		dao.PlayMemberLevel.Columns().Status: in.Status,
		dao.PlayMemberLevel.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayMemberLevel.Ctx(ctx).Where(dao.PlayMemberLevel.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除会员等级表
func (s *sMemberLevel) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayMemberLevel.Ctx(ctx).Where(dao.PlayMemberLevel.Columns().Id, id).Data(g.Map{
		dao.PlayMemberLevel.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取会员等级表详情
func (s *sMemberLevel) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.MemberLevelDetailOutput, err error) {
	out = &model.MemberLevelDetailOutput{}
	err = dao.PlayMemberLevel.Ctx(ctx).Where(dao.PlayMemberLevel.Columns().Id, id).Where(dao.PlayMemberLevel.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	return
}

// List 获取会员等级表列表
func (s *sMemberLevel) List(ctx context.Context, in *model.MemberLevelListInput) (list []*model.MemberLevelListOutput, total int, err error) {
	m := dao.PlayMemberLevel.Ctx(ctx).Where(dao.PlayMemberLevel.Columns().DeletedAt, nil)
	if in.Level > 0 {
		m = m.Where(dao.PlayMemberLevel.Columns().Level, in.Level)
	}
	if in.Status > 0 {
		m = m.Where(dao.PlayMemberLevel.Columns().Status, in.Status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayMemberLevel.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	return
}

