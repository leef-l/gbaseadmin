package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/crypto/bcrypt"

	"gbaseadmin/app/system/internal/dao"
	"gbaseadmin/app/system/internal/model"
	"gbaseadmin/app/system/internal/service"
	"gbaseadmin/utility/jwt"
	"gbaseadmin/utility/snowflake"
)

func init() {
	service.RegisterAuth(New())
}

func New() *sAuth {
	return &sAuth{}
}

type sAuth struct{}

// Login 用户登录
func (s *sAuth) Login(ctx context.Context, in *model.AuthLoginInput) (out *model.AuthLoginOutput, err error) {
	// 查询用户
	var user struct {
		Id       int64  `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
		DeptId   int64  `json:"deptId"`
		Status   int    `json:"status"`
	}
	err = dao.Users.Ctx(ctx).
		Where(dao.Users.Columns().Username, in.Username).
		Where(dao.Users.Columns().DeletedAt, nil).
		Scan(&user)
	if err != nil {
		return nil, gerror.New("用户名或密码错误")
	}
	if user.Id == 0 {
		return nil, gerror.New("用户名或密码错误")
	}

	// 校验状态
	if user.Status == 0 {
		return nil, gerror.New("账号已被禁用")
	}

	// 校验密码
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return nil, gerror.New("用户名或密码错误")
	}

	// 生成 Token
	token, err := jwt.GenerateToken(user.Id, user.Username, user.DeptId)
	if err != nil {
		return nil, gerror.New("生成Token失败")
	}

	out = &model.AuthLoginOutput{
		Token:    token,
		UserID:   snowflake.JsonInt64(user.Id),
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}
	return
}

// Info 获取当前用户信息
func (s *sAuth) Info(ctx context.Context, userID snowflake.JsonInt64) (out *model.AuthInfoOutput, err error) {
	var user struct {
		Id       int64  `json:"id"`
		Username string `json:"username"`
		Nickname string `json:"nickname"`
		Email    string `json:"email"`
		Avatar   string `json:"avatar"`
		DeptId   int64  `json:"deptId"`
		Status   int    `json:"status"`
	}
	err = dao.Users.Ctx(ctx).
		Where(dao.Users.Columns().Id, userID).
		Where(dao.Users.Columns().DeletedAt, nil).
		Scan(&user)
	if err != nil {
		return nil, err
	}

	out = &model.AuthInfoOutput{
		UserID:   snowflake.JsonInt64(user.Id),
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Avatar:   user.Avatar,
		DeptID:   snowflake.JsonInt64(user.DeptId),
		Status:   user.Status,
		Roles:    make([]string, 0),
		Perms:    make([]string, 0),
	}

	// 查询用户角色
	var userRoles []struct {
		RoleId int64 `json:"roleId"`
	}
	_ = dao.UserRole.Ctx(ctx).Where(dao.UserRole.Columns().UserId, userID).Scan(&userRoles)

	if len(userRoles) > 0 {
		roleIDs := make([]int64, 0, len(userRoles))
		for _, ur := range userRoles {
			roleIDs = append(roleIDs, ur.RoleId)
		}

		// 查询角色名称
		var roles []struct {
			Title string `json:"title"`
		}
		_ = g.DB().Ctx(ctx).Model("role").
			Where("id", roleIDs).
			Where("deleted_at", nil).
			Where("status", 1).
			Scan(&roles)
		for _, r := range roles {
			out.Roles = append(out.Roles, r.Title)
		}

		// 查询角色关联的菜单权限标识
		var menuIDs []struct {
			MenuId int64 `json:"menuId"`
		}
		_ = dao.RoleMenu.Ctx(ctx).WhereIn(dao.RoleMenu.Columns().RoleId, roleIDs).Scan(&menuIDs)

		if len(menuIDs) > 0 {
			mIDs := make([]int64, 0, len(menuIDs))
			for _, m := range menuIDs {
				mIDs = append(mIDs, m.MenuId)
			}
			var perms []struct {
				Permission string `json:"permission"`
			}
			_ = g.DB().Ctx(ctx).Model("menu").
				Where("id", mIDs).
				Where("deleted_at", nil).
				Where("status", 1).
				WhereNot("permission", "").
				Scan(&perms)
			seen := make(map[string]bool)
			for _, p := range perms {
				if p.Permission != "" && !seen[p.Permission] {
					out.Perms = append(out.Perms, p.Permission)
					seen[p.Permission] = true
				}
			}
		}
	}

	return
}

// ChangePassword 修改密码
func (s *sAuth) ChangePassword(ctx context.Context, in *model.AuthChangePasswordInput) error {
	// 查询当前密码
	password, err := dao.Users.Ctx(ctx).
		Where(dao.Users.Columns().Id, in.UserID).
		Value(dao.Users.Columns().Password)
	if err != nil {
		return err
	}

	// 校验旧密码
	if err = bcrypt.CompareHashAndPassword([]byte(password.String()), []byte(in.OldPassword)); err != nil {
		return gerror.New("旧密码错误")
	}

	// 加密新密码
	hashed, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 更新密码
	_, err = dao.Users.Ctx(ctx).
		Where(dao.Users.Columns().Id, in.UserID).
		Data(dao.Users.Columns().Password, string(hashed)).
		Update()
	return err
}

// Menus 获取当前用户的菜单树（动态路由）
func (s *sAuth) Menus(ctx context.Context, userID snowflake.JsonInt64) ([]*model.AuthMenuOutput, error) {
	// 查询用户角色
	var userRoles []struct {
		RoleId int64 `json:"roleId"`
	}
	err := dao.UserRole.Ctx(ctx).Where(dao.UserRole.Columns().UserId, userID).Scan(&userRoles)
	if err != nil {
		return nil, err
	}

	if len(userRoles) == 0 {
		return make([]*model.AuthMenuOutput, 0), nil
	}

	roleIDs := make([]int64, 0, len(userRoles))
	for _, ur := range userRoles {
		roleIDs = append(roleIDs, ur.RoleId)
	}

	// 查询角色关联的菜单ID（去重）
	var roleMenus []struct {
		MenuId int64 `json:"menuId"`
	}
	err = dao.RoleMenu.Ctx(ctx).WhereIn(dao.RoleMenu.Columns().RoleId, roleIDs).Scan(&roleMenus)
	if err != nil {
		return nil, err
	}

	if len(roleMenus) == 0 {
		return make([]*model.AuthMenuOutput, 0), nil
	}

	menuIDSet := make(map[int64]bool)
	menuIDs := make([]int64, 0)
	for _, rm := range roleMenus {
		if !menuIDSet[rm.MenuId] {
			menuIDSet[rm.MenuId] = true
			menuIDs = append(menuIDs, rm.MenuId)
		}
	}

	// 查询菜单详情
	var list []*model.AuthMenuOutput
	err = g.DB().Ctx(ctx).Model("menu").
		Where("id", menuIDs).
		Where("deleted_at", nil).
		Where("status", 1).
		OrderAsc("sort").
		Scan(&list)
	if err != nil {
		return nil, err
	}

	// 组装树
	nodeMap := make(map[int64]*model.AuthMenuOutput, len(list))
	for _, item := range list {
		item.Children = make([]*model.AuthMenuOutput, 0)
		nodeMap[int64(item.ID)] = item
	}

	tree := make([]*model.AuthMenuOutput, 0)
	for _, item := range list {
		if int64(item.ParentID) == 0 {
			tree = append(tree, item)
		} else if parent, ok := nodeMap[int64(item.ParentID)]; ok {
			parent.Children = append(parent.Children, item)
		} else {
			// 父节点不在权限范围内，作为顶级节点
			tree = append(tree, item)
		}
	}
	return tree, nil
}
