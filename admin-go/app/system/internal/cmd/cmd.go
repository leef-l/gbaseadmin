package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gbaseadmin/app/system/internal/controller/dept"
	"gbaseadmin/app/system/internal/controller/hello"
	"gbaseadmin/app/system/internal/controller/menu"
	"gbaseadmin/app/system/internal/controller/role"
	"gbaseadmin/app/system/internal/controller/users"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
				)
				// 系统管理模块
				group.Group("/api/system", func(group *ghttp.RouterGroup) {
					group.Bind(
						dept.Dept,
						role.Role,
						menu.Menu,
						users.Users,
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
