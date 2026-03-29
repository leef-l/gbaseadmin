package users

import (
	"context"

	"github.com/gogf/gf/v2/crypto/gsha256"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

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
	hashedPassword := gsha256.Encrypt(in.Password)
	_, err := dao.Users.Ctx(ctx).Data(g.Map{
		dao.Users.Columns().Id:        id,
		dao.Users.Columns().Username: in.Username,
		dao.Users.Columns().Password: hashedPassword,
		dao.Users.Columns().Nickname: in.Nickname,
		dao.Users.Columns().Email: in.Email,
		dao.Users.Columns().Avatar: in.Avatar,
		dao.Users.Columns().Status: in.Status,
		dao.Users.Columns().DeptId: in.DeptID,
		dao.Users.Columns().CreatedAt: gtime.Now(),
		dao.Users.Columns().UpdatedAt: gtime.Now(),
	}).Insert()
	if err != nil {
		return err
	}
	// 写入用户角色关联
	if len(in.RoleIDs) > 0 {
		data := make([]g.Map, 0, len(in.RoleIDs))
		for _, roleID := range in.RoleIDs {
			data = append(data, g.Map{
				dao.UserRole.Columns().UserId: id,
				dao.UserRole.Columns().RoleId: roleID,
			})
		}
		_, err = dao.UserRole.Ctx(ctx).Data(data).Insert()
	}
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
		dao.Users.Columns().DeptId: in.DeptID,
		dao.Users.Columns().UpdatedAt: gtime.Now(),
	}
	if in.Password != "" {
		data[dao.Users.Columns().Password] = gsha256.Encrypt(in.Password)
	}
	_, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, in.ID).Data(data).Update()
	if err != nil {
		return err
	}
	// 更新用户角色关联（先删后插）
	_, err = dao.UserRole.Ctx(ctx).Where(dao.UserRole.Columns().UserId, in.ID).Delete()
	if err != nil {
		return err
	}
	if len(in.RoleIDs) > 0 {
		roleData := make([]g.Map, 0, len(in.RoleIDs))
		for _, roleID := range in.RoleIDs {
			roleData = append(roleData, g.Map{
				dao.UserRole.Columns().UserId: in.ID,
				dao.UserRole.Columns().RoleId: roleID,
			})
		}
		_, err = dao.UserRole.Ctx(ctx).Data(roleData).Insert()
	}
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
	// 查询部门名称
	if out.DeptID != 0 {
		val, _ := g.DB().Ctx(ctx).Model("dept").Where("id", out.DeptID).Where("deleted_at", nil).Value("title")
		out.DeptTitle = val.String()
	}
	// 查询用户角色ID列表
	var roles []struct {
		RoleId int64 `json:"roleId"`
	}
	_ = dao.UserRole.Ctx(ctx).Where(dao.UserRole.Columns().UserId, id).Scan(&roles)
	out.RoleIDs = make([]snowflake.JsonInt64, 0, len(roles))
	for _, r := range roles {
		out.RoleIDs = append(out.RoleIDs, snowflake.JsonInt64(r.RoleId))
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
	// 填充部门名称
	for _, item := range list {
		if item.DeptID != 0 {
			val, _ := g.DB().Ctx(ctx).Model("dept").Where("id", item.DeptID).Where("deleted_at", nil).Value("title")
			item.DeptTitle = val.String()
		}
	}
	return
}

