package coach_apply

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
	service.RegisterCoachApply(New())
}

func New() *sCoachApply {
	return &sCoachApply{}
}

type sCoachApply struct{}

// Create 创建陪玩师申请表
func (s *sCoachApply) Create(ctx context.Context, in *model.CoachApplyCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayCoachApply.Ctx(ctx).Data(g.Map{
		dao.PlayCoachApply.Columns().Id:        id,
		dao.PlayCoachApply.Columns().MemberId: in.MemberID,
		dao.PlayCoachApply.Columns().RealName: in.RealName,
		dao.PlayCoachApply.Columns().IdCard: in.IDCard,
		dao.PlayCoachApply.Columns().IdCardFrontImage: in.IDCardFrontImage,
		dao.PlayCoachApply.Columns().IdCardBackImage: in.IDCardBackImage,
		dao.PlayCoachApply.Columns().SkillDesc: in.SkillDesc,
		dao.PlayCoachApply.Columns().AuditStatus: in.AuditStatus,
		dao.PlayCoachApply.Columns().AuditRemark: in.AuditRemark,
		dao.PlayCoachApply.Columns().AuditAt: in.AuditAt,
		dao.PlayCoachApply.Columns().CreatedAt: gtime.Now(),
		dao.PlayCoachApply.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新陪玩师申请表
func (s *sCoachApply) Update(ctx context.Context, in *model.CoachApplyUpdateInput) error {
	data := g.Map{
		dao.PlayCoachApply.Columns().MemberId: in.MemberID,
		dao.PlayCoachApply.Columns().RealName: in.RealName,
		dao.PlayCoachApply.Columns().IdCard: in.IDCard,
		dao.PlayCoachApply.Columns().IdCardFrontImage: in.IDCardFrontImage,
		dao.PlayCoachApply.Columns().IdCardBackImage: in.IDCardBackImage,
		dao.PlayCoachApply.Columns().SkillDesc: in.SkillDesc,
		dao.PlayCoachApply.Columns().AuditStatus: in.AuditStatus,
		dao.PlayCoachApply.Columns().AuditRemark: in.AuditRemark,
		dao.PlayCoachApply.Columns().AuditAt: in.AuditAt,
		dao.PlayCoachApply.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayCoachApply.Ctx(ctx).Where(dao.PlayCoachApply.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除陪玩师申请表
func (s *sCoachApply) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayCoachApply.Ctx(ctx).Where(dao.PlayCoachApply.Columns().Id, id).Data(g.Map{
		dao.PlayCoachApply.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取陪玩师申请表详情
func (s *sCoachApply) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.CoachApplyDetailOutput, err error) {
	out = &model.CoachApplyDetailOutput{}
	err = dao.PlayCoachApply.Ctx(ctx).Where(dao.PlayCoachApply.Columns().Id, id).Where(dao.PlayCoachApply.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	// 查询会员昵称
	if out.MemberID != 0 {
		val, _ := g.DB().Ctx(ctx).Model("play_member").Where("id", out.MemberID).Value("nickname")
		out.MemberNickname = val.String()
	}
	return
}

// List 获取陪玩师申请表列表
func (s *sCoachApply) List(ctx context.Context, in *model.CoachApplyListInput) (list []*model.CoachApplyListOutput, total int, err error) {
	m := dao.PlayCoachApply.Ctx(ctx).Where(dao.PlayCoachApply.Columns().DeletedAt, nil)
	if in.AuditStatus > 0 {
		m = m.Where(dao.PlayCoachApply.Columns().AuditStatus, in.AuditStatus)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayCoachApply.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	// 填充关联显示字段
	for _, item := range list {
		if item.MemberID != 0 {
			val, _ := g.DB().Ctx(ctx).Model("play_member").Where("id", item.MemberID).Value("nickname")
			item.MemberNickname = val.String()
		}
	}
	return
}



