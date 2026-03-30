package oauth

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
	service.RegisterOauth(New())
}

func New() *sOauth {
	return &sOauth{}
}

type sOauth struct{}

// Create 创建ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨
func (s *sOauth) Create(ctx context.Context, in *model.OauthCreateInput) error {
	id := snowflake.Generate()
	_, err := dao.PlayOauth.Ctx(ctx).Data(g.Map{
		dao.PlayOauth.Columns().Id:        id,
		dao.PlayOauth.Columns().MemberId: in.MemberID,
		dao.PlayOauth.Columns().Provider: in.Provider,
		dao.PlayOauth.Columns().OpenId: in.OpenID,
		dao.PlayOauth.Columns().UnionId: in.UnionID,
		dao.PlayOauth.Columns().Nickname: in.Nickname,
		dao.PlayOauth.Columns().Avatar: in.Avatar,
		dao.PlayOauth.Columns().AccessToken: in.AccessToken,
		dao.PlayOauth.Columns().RefreshToken: in.RefreshToken,
		dao.PlayOauth.Columns().ExpireAt: in.ExpireAt,
		dao.PlayOauth.Columns().CreatedAt: gtime.Now(),
		dao.PlayOauth.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨
func (s *sOauth) Update(ctx context.Context, in *model.OauthUpdateInput) error {
	data := g.Map{
		dao.PlayOauth.Columns().MemberId: in.MemberID,
		dao.PlayOauth.Columns().Provider: in.Provider,
		dao.PlayOauth.Columns().OpenId: in.OpenID,
		dao.PlayOauth.Columns().UnionId: in.UnionID,
		dao.PlayOauth.Columns().Nickname: in.Nickname,
		dao.PlayOauth.Columns().Avatar: in.Avatar,
		dao.PlayOauth.Columns().AccessToken: in.AccessToken,
		dao.PlayOauth.Columns().RefreshToken: in.RefreshToken,
		dao.PlayOauth.Columns().ExpireAt: in.ExpireAt,
		dao.PlayOauth.Columns().UpdatedAt: gtime.Now(),
	}
	_, err := dao.PlayOauth.Ctx(ctx).Where(dao.PlayOauth.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨
func (s *sOauth) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.PlayOauth.Ctx(ctx).Where(dao.PlayOauth.Columns().Id, id).Data(g.Map{
		dao.PlayOauth.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨详情
func (s *sOauth) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.OauthDetailOutput, err error) {
	out = &model.OauthDetailOutput{}
	err = dao.PlayOauth.Ctx(ctx).Where(dao.PlayOauth.Columns().Id, id).Where(dao.PlayOauth.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	return
}

// List 获取ç¬¬ä¸‰æ–¹ç™»å½•ç»‘å®šè¡¨列表
func (s *sOauth) List(ctx context.Context, in *model.OauthListInput) (list []*model.OauthListOutput, total int, err error) {
	m := dao.PlayOauth.Ctx(ctx).Where(dao.PlayOauth.Columns().DeletedAt, nil)
	if in.Provider > 0 {
		m = m.Where(dao.PlayOauth.Columns().Provider, in.Provider)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.PlayOauth.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	return
}

