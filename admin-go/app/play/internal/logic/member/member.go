package member

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"golang.org/x/crypto/bcrypt"

	"gbaseadmin/app/play/internal/dao"
	"gbaseadmin/app/play/internal/model"
	"gbaseadmin/app/play/internal/service"
	"gbaseadmin/utility/snowflake"
)

func init() {
	service.RegisterMember(New())
}

func New() *sMember {
	return &sMember{}
}

type sMember struct{}

// Create 创建会员表
func (s *sMember) Create(ctx context.Context, in *model.MemberCreateInput) error {
	id := snowflake.Generate()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = dao.PlayMember.Ctx(ctx).Data(g.Map{
		dao.PlayMember.Columns().Id:        id,
		dao.PlayMember.Columns().Phone: in.Phone,
		dao.PlayMember.Columns().Password: string(hashedPassword),
		dao.PlayMember.Columns().Nickname: in.Nickname,
		dao.PlayMember.Columns().Avatar: in.Avatar,
		dao.PlayMember.Columns().Gender: in.Gender,
		dao.PlayMember.Columns().MemberLevelId: in.MemberLevelID,
		dao.PlayMember.Columns().Exp: in.Exp,
		dao.PlayMember.Columns().Balance: in.Balance,
		dao.PlayMember.Columns().IsCoach: in.IsCoach,
		dao.PlayMember.Columns().Status: in.Status,
		dao.PlayMember.Columns().LastLoginAt: in.LastLoginAt,
		dao.PlayMember.Columns().CreatedAt: gtime.Now(),
		dao.PlayMember.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新会员表
func (s *sMember) Update(ctx context.Context, in *model.MemberUpdateInput) error {
	data := g.Map{
		dao.PlayMember.Columns().Phone: in.Phone,
		dao.PlayMember.Columns().Nickname: in.Nickname,
		dao.PlayMember.Columns().Avatar: in.Avatar,
		dao.PlayMember.Columns().Gender: in.Gender,
		dao.PlayMember.Columns().MemberLevelId: in.MemberLevelID,
		dao.PlayMember.Columns().Exp: in.Exp,
		dao.PlayMember.Columns().Balance: in.Balance,
		dao.PlayMember.Columns().IsCoach: in.IsCoach,
		dao.PlayMember.Columns().Status: in.Status,
		dao.PlayMember.Columns().LastLoginAt: in.LastLoginAt,
		dao.PlayMember.Columns().UpdatedAt: gtime.Now(),
	}
	if in.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		data[dao.PlayMember.Columns().Password] = string(hashed)
	}
	_, err := dao.PlayMember.Ctx(ctx).Where(dao.PlayMember.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除会员表
func (s *sMember) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayMember.Ctx(ctx).Where(dao.PlayMember.Columns().Id, id).Data(g.Map{
		dao.PlayMember.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取会员表详情
func (s *sMember) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.MemberDetailOutput, err error) {
	out = &model.MemberDetailOutput{}
	err = dao.PlayMember.Ctx(ctx).Where(dao.PlayMember.Columns().Id, id).Where(dao.PlayMember.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	// 查询会员等级ID关联显示
	if out.MemberLevelID != 0 {
		val, _ := g.DB().Ctx(ctx).Model("play_member_level").Where("id", out.MemberLevelID).Where("deleted_at", nil).Value("title")
		out.MemberLevelTitle = val.String()
	}
	return
}

// List 获取会员表列表
func (s *sMember) List(ctx context.Context, in *model.MemberListInput) (list []*model.MemberListOutput, total int, err error) {
	m := dao.PlayMember.Ctx(ctx).Where(dao.PlayMember.Columns().DeletedAt, nil)
	if in.Gender > 0 {
		m = m.Where(dao.PlayMember.Columns().Gender, in.Gender)
	}
	if in.IsCoach > 0 {
		m = m.Where(dao.PlayMember.Columns().IsCoach, in.IsCoach)
	}
	if in.Status > 0 {
		m = m.Where(dao.PlayMember.Columns().Status, in.Status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayMember.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	// 填充关联显示字段
	for _, item := range list {
		if item.MemberLevelID != 0 {
			val, _ := g.DB().Ctx(ctx).Model("play_member_level").Where("id", item.MemberLevelID).Where("deleted_at", nil).Value("title")
			item.MemberLevelTitle = val.String()
		}
	}
	return
}

