# Phase 2：权限框架开发计划

## 总体进度

| 阶段 | 任务 | 状态 | 负责 Agent | 备注 |
|---|---|---|---|---|
| **Phase 0** | DevOps 环境搭建 | ✅ 完成 | devops | docker-compose、deploy/、Makefile |
| **Phase 1-A** | codegen-parser | ✅ 完成 | codegen-parser | parser/meta.go 等 4 文件 |
| **Phase 1-A** | codegen-backend 模板 | ✅ 完成 | codegen-backend | 7 个 .tpl 模板 |
| **Phase 1-A** | codegen-frontend 模板 | ✅ 完成 | codegen-frontend | 4 个 .tpl 模板 |
| **Phase 1-B** | codegen-engine 引擎组装 | ✅ 完成 | codegen-engine | CLI 入口 + 配置系统 |
| **Phase 1-C** | codegen-tester 验证 | ✅ 完成 | codegen-tester | 4 表 44 文件生成成功 |
| **Phase 2-Init** | 项目初始化 + DAO 生成 | ✅ 完成 | orchestrator | gf gen dao、路由注册 |
| **Phase 2-D1** | JWT 认证 + 中间件 + 工具包 | ✅ 完成 | backend-specialist | jwt、response、auth、middleware |
| **Phase 2-D2** | 树形业务 + 授权 + 数据权限 | ✅ 完成 | backend-specialist | 授权菜单/部门、关联表、动态路由 |
| **Phase 2-D3** | 特殊 UI 组件 | ✅ 完成 | frontend-specialist | 授权弹窗、部门树选择、角色多选 |
| **Phase 2-E** | 动态路由 + 权限指令 | ⏳ 待开始 | frontend-specialist | accessMode backend、v-access |
| **Phase 2-F** | 全量代码审查 | ⏳ 待开始 | reviewer | 问题清单输出 |

---

## Phase 2 详细任务分解

### Phase 2-Init：项目初始化（orchestrator 直接执行）

- [ ] 确认 GoFrame MonoRepo 结构正确
- [ ] 确认 hack/config.yaml 数据库配置
- [ ] 执行 `gf gen dao` 生成 DAO/DO/Entity 层
- [ ] 注册路由到 cmd.go（整合 dept/role/menu/users 路由）
- [ ] 确认 go mod tidy 编译通过

### Phase 2-D1：JWT 认证 + 中间件 + 工具包

**负责：backend-specialist**

产出文件：
- `utility/snowflake/snowflake.go` — Snowflake ID 生成器 + JsonInt64
- `utility/response/response.go` — 统一响应封装 {code, message, data}
- `app/system/internal/middleware/auth.go` — JWT 鉴权中间件
- `app/system/internal/logic/auth/auth.go` — 登录/Token 生成/刷新逻辑
- `app/system/api/system/v1/auth.go` — 登录接口 API 定义
- `app/system/internal/controller/auth/auth.go` — 登录控制器

### Phase 2-D2：树形业务 + 授权 + 数据权限

**负责：backend-specialist**

产出文件：
- `app/system/internal/logic/dept/dept.go` — 补充树形递归、数据权限验证
- `app/system/internal/logic/role/role.go` — 补充授权菜单(grant_menu)、授权部门(grant_dept)
- `app/system/internal/logic/menu/menu.go` — 补充动态路由菜单生成
- `app/system/internal/logic/users/users.go` — 补充数据权限过滤(data_scope)
- `app/system/internal/middleware/datascope.go` — 数据权限中间件

### Phase 2-D3：特殊 UI 组件

**负责：frontend-specialist**

产出文件：
- `views/system/role/modules/grant-menu.vue` — 角色授权菜单弹窗（树形 CheckBox）
- `views/system/role/modules/grant-dept.vue` — 数据权限弹窗（范围选择 + 部门树）
- `views/system/users/modules/form.vue` — 增强：部门树选择 + 角色多选
- `api/system/common/index.ts` — 公共 consts 接口

### Phase 2-E：动态路由 + 权限指令

**负责：frontend-specialist**（等待 D 完成后执行）

产出文件：
- Vben Admin accessMode 配置为 backend
- 菜单接口对接（/api/system/auth/info）
- 路由守卫配置
- 按钮级权限指令 v-access
- API 请求拦截器（JWT Token、401 刷新）

### Phase 2-F：全量代码审查

**负责：reviewer**（等待 E 完成后执行）

检查项：
1. GoFrame 分层是否严格
2. 所有 ID 字段是否使用 JsonInt64
3. 软删除是否正确
4. JWT 中间件覆盖范围
5. data_scope 过滤是否生效
6. 前端 BigInt 处理
7. 树形 children 空值处理
8. 错误处理统一性

---

## 执行日志

| 时间 | 操作 | 结果 |
|---|---|---|
| 2026-03-29 | 创建 Phase 2 计划文档 | ✅ |
