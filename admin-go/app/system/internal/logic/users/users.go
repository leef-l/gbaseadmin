package users

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"golang.org/x/crypto/bcrypt"

	"gbaseadmin/app/system/internal/dao"
	"gbaseadmin/app/system/internal/model"
	"gbaseadmin/app/system/internal/service"
	"gbaseadmin/utility/snowflake"
)

func init() {
	service.RegisterUsers(New())
}

func New() *sUsers {
	return &sUsers{}
}

type sUsers struct{}

// Create 创建用户表
func (s *sUsers) Create(ctx context.Context, in *model.UsersCreateInput) error {
	id := snowflake.Generate()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = dao.Users.Ctx(ctx).Data(g.Map{
		dao.Users.Columns().Id:        id,
		dao.Users.Columns().Username: in.Username,
		dao.Users.Columns().Password: string(hashedPassword),
		dao.Users.Columns().Nickname: in.Nickname,
		dao.Users.Columns().Email: in.Email,
		dao.Users.Columns().Avatar: in.Avatar,
		dao.Users.Columns().Status: in.Status,
		dao.Users.Columns().CreatedAt: gtime.Now(),
		dao.Users.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 更新用户表
func (s *sUsers) Update(ctx context.Context, in *model.UsersUpdateInput) error {
	data := g.Map{
		dao.Users.Columns().Username: in.Username,
		dao.Users.Columns().Nickname: in.Nickname,
		dao.Users.Columns().Email: in.Email,
		dao.Users.Columns().Avatar: in.Avatar,
		dao.Users.Columns().Status: in.Status,
		dao.Users.Columns().UpdatedAt: gtime.Now(),
	}
	if in.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		data[dao.Users.Columns().Password] = string(hashed)
	}
	_, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, in.ID).Data(data).Update()
	return err
}

// Delete 软删除用户表
func (s *sUsers) Delete(ctx context.Context, id snowflake.JsonInt64) error {
	_, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, id).Data(g.Map{
		dao.Users.Columns().DeletedAt: gtime.Now(),
	}).Update()
	return err
}

// Detail 获取用户表详情
func (s *sUsers) Detail(ctx context.Context, id snowflake.JsonInt64) (out *model.UsersDetailOutput, err error) {
	out = &model.UsersDetailOutput{}
	err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, id).Where(dao.Users.Columns().DeletedAt, nil).Scan(out)
	if err != nil {
		return nil, err
	}
	return
}

// List 获取用户表列表
func (s *sUsers) List(ctx context.Context, in *model.UsersListInput) (list []*model.UsersListOutput, total int, err error) {
	m := dao.Users.Ctx(ctx).Where(dao.Users.Columns().DeletedAt, nil)
	if in.Status > 0 {
		m = m.Where(dao.Users.Columns().Status, in.Status)
	}
	total, err = m.Count()
	if err != nil {
		return
	}
	err = m.Page(in.PageNum, in.PageSize).OrderAsc(dao.Users.Columns().Id).Scan(&list)
	if err != nil {
		return
	}
	return
}

