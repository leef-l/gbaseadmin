# GBaseAdmin Agent Team 工作流程设计

## 概述

本文档定义了使用 Claude Code Agent Team 模式构建 GBaseAdmin 权限框架的完整团队协作流程。
分两大阶段：**Phase 1 代码生成器**（最高优先级）→ **Phase 2 权限框架**。

---

## 团队角色定义

| Agent 名称 | 职责 | 工具权限 | 文件位置 |
|---|---|---|---|
| `orchestrator` | 总指挥，分配任务，汇总结果 | 全部 | 主会话 |
| `devops` | docker-compose 环境搭建 | Read/Write/Edit/Bash | `.claude/agents/` |
| `codegen-parser` | 解析数据库表结构 → 元数据 | Read/Write/Bash | `.claude/agents/` |
| `codegen-backend` | 生成后端代码（GoFrame v2） | Read/Write/Edit/Bash | `.claude/agents/` |
| `codegen-frontend` | 生成前端代码（Vben Admin） | Read/Write/Edit/Bash | `.claude/agents/` |
| `codegen-engine` | 组装生成器 CLI 入口 | Read/Write/Edit/Bash | `.claude/agents/` |
| `codegen-tester` | 验证生成器输出是否正确 | Read/Bash | `.claude/agents/` |
| `backend-specialist` | 编写 GoFrame 业务逻辑 | Read/Write/Edit/Bash | `.claude/agents/` |
| `frontend-specialist` | 编写 Vben Admin 页面 | Read/Write/Edit/Bash | `.claude/agents/` |
| `reviewer` | 代码审查，发现问题 | Read/Grep/Glob | `.claude/agents/` |

---

## Phase 0：环境搭建（最先执行）

### 执行方式

由 `devops` Agent 独立完成，与 Phase 1 并行启动（不阻塞代码生成器开发）。

**触发指令：**
```
使用 devops Agent：
读取架构设计文档，在 /www/wwwroot/project/gbaseadmin/admin-go/ 创建完整的 docker-compose 环境配置，
包括 MySQL 8.0（含自动建表 init.sql）、Redis 7、后端 Go 服务、前端 Vben Admin、adminer 管理界面。
同时创建 .env.example、Makefile 和 deploy/ 目录下的所有配置文件。
```

**产物：**
- `docker-compose.yml`（开发环境）
- `docker-compose.prod.yml`（生产环境）
- `.env.example`
- `Makefile`
- `deploy/mysql/init.sql`（含所有建表 SQL + 初始数据）
- `deploy/Dockerfile.system`
- `deploy/nginx/nginx.conf`

**启动命令：**
```bash
cp .env.example .env
# 修改 .env 中的密码
docker-compose up -d
```

---

## Phase 1：代码生成器

### 整体流程图

```
┌─────────────────────────────────────────────────────────────┐
│                        orchestrator                          │
│  分析架构文档 → 拆解任务 → 分发 → 汇总结果 → 验证           │
└──────────────────────────┬──────────────────────────────────┘
                           │
         ┌─────────────────┼─────────────────┐
         ▼                 ▼                 ▼
   [并行阶段 A]      [并行阶段 A]      [并行阶段 A]
  codegen-parser   codegen-backend  codegen-frontend
  解析DB表→元数据   后端模板设计      前端模板设计
         │                 │                 │
         └─────────────────┼─────────────────┘
                           ▼
                   [串行阶段 B]
                  codegen-engine
               组装 CLI + 配置系统
                           │
                           ▼
                   [串行阶段 C]
                  codegen-tester
               用 dept 表跑通全流程
```

### 详细步骤

#### 阶段 A（并行，3个 Agent 同时执行）

**A1 - codegen-parser（DB 解析器）**
```
输入：数据库连接配置
输出：codegen/parser/ 完整代码
产物：
  - FieldMeta 结构（字段名/类型/备注/是否可空等）
  - TableMeta 结构（表名/注释/字段列表/是否有parent_id等）
  - 字段备注解析函数（提取 label 和枚举值）
  - 组件类型识别函数（按字段名规则映射前端组件）
  - 数据库连接和查询函数
```

**A2 - codegen-backend（后端模板）**
```
输入：架构设计文档 + GoFrame v2 规范
输出：codegen/templates/backend/ 所有模板
产物：
  - api.tpl       （Req/Res 结构体定义）
  - controller.tpl（Controller 层）
  - logic.tpl     （Logic 层业务实现）
  - service.tpl   （Service 接口定义）
  - model.tpl     （DTO 模型）
  - consts.tpl    （枚举常量）
  - router.tpl    （路由注册）
```

**A3 - codegen-frontend（前端模板）**
```
输入：架构设计文档 + Vben Admin web-antd 规范
输出：codegen/templates/frontend/ 所有模板
产物：
  - api.tpl   （TypeScript API 调用）
  - types.tpl （TypeScript 类型定义）
  - list.tpl  （列表页 list.vue，含树形/普通切换）
  - form.tpl  （表单弹窗 form.vue，含组件映射）
```

#### 阶段 B（串行，等待 A 完成）

**B1 - codegen-engine（引擎组装）**
```
输入：A1/A2/A3 的产物
输出：
  - codegen/generator/backend/  （调用解析器+后端模板生成代码）
  - codegen/generator/frontend/ （调用解析器+前端模板生成代码）
  - codegen/main.go             （CLI 入口，支持 --table/--only/--force 参数）
  - codegen/codegen.yaml        （默认配置，可覆盖字段规则）
```

#### 阶段 C（串行，等待 B 完成）

**C1 - codegen-tester（验证生成器）**
```
输入：已完成的 codegen 工具 + dept 表
执行：
  1. go run codegen/main.go --table=dept
  2. 检查生成的后端文件结构是否符合 GoFrame v2 规范
  3. 检查生成的前端文件结构是否符合 Vben Admin 规范
  4. 检查 Snowflake ID 是否正确使用 JsonInt64
  5. 检查含 parent_id 是否生成了 /tree 和 /list 两个接口
  6. 检查枚举备注是否正确生成了 consts
输出：验证报告，列出问题（如有）
```

---

## Phase 2：权限框架

### 前置条件

- 代码生成器已完成并验证通过
- `gf init` 项目初始化已完成
- 数据库表已创建

### 整体流程图

```
┌─────────────────────────────────────────────────────────────┐
│                        orchestrator                          │
│  调用 codegen 批量生成 → 分配手写任务 → 汇总联调            │
└──────────────────────────┬──────────────────────────────────┘
                           │
              [串行] 批量代码生成
              codegen 生成 dept/role/menu/users
                           │
         ┌─────────────────┼─────────────────┐
         ▼                 ▼                 ▼
   [并行阶段 D]      [并行阶段 D]      [并行阶段 D]
  backend-specialist  backend-specialist  frontend-specialist
  JWT认证+中间件       树形+授权业务逻辑    特殊UI组件
         │                 │                 │
         └─────────────────┼─────────────────┘
                           ▼
                   [串行阶段 E]
                  frontend-specialist
                  动态路由+权限指令
                           │
                           ▼
                   [串行阶段 F]
                     reviewer
                   全量代码审查
```

### 详细步骤

#### 初始化（orchestrator 直接执行）

```bash
# 1. 初始化 GoFrame MonoRepo
cd admin-go && gf init gbaseadmin -m

# 2. 初始化权限服务
gf init app/system -a

# 3. 配置数据库连接（hack/config.yaml）
# 4. 执行建表 SQL
# 5. gf gen dao（生成 DAO 层）
# 6. 批量运行代码生成器
go run codegen/main.go --table=dept,role,menu,users
```

#### 阶段 D（并行，3个 Agent 同时执行）

**D1 - backend-specialist（JWT + 中间件）**
```
任务：
  - utility/snowflake/snowflake.go（JsonInt64 实现）
  - utility/response/response.go（统一响应封装）
  - app/system/internal/middleware/auth.go（JWT 鉴权中间件）
  - app/system/internal/logic/auth/auth.go（登录/Token生成/刷新）
  - app/system/api/system/v1/auth.go（登录接口定义）
  - app/system/internal/controller/auth/auth.go（登录控制器）
```

**D2 - backend-specialist（树形 + 授权业务）**
```
任务（补充代码生成器生成的代码中缺少的业务逻辑）：
  - dept/logic：树形递归组装、数据权限验证
  - role/logic：树形递归组装、授权菜单(grant_menu)、授权部门(grant_dept)
  - menu/logic：树形递归组装、动态路由菜单生成
  - user/logic：数据权限过滤（data_scope 字段处理）
  - 数据权限中间件：根据 data_scope 自动注入 SQL WHERE 条件
```

**D3 - frontend-specialist（特殊 UI 组件）**
```
任务：
  - 角色授权弹窗（菜单树 CheckBox 多选）
  - 数据权限弹窗（范围选择 + 部门树 CheckBox 多选，data_scope=5时显示）
  - 用户表单部门选择（树形多选，仅显示有权限的部门）
  - 用户表单角色选择（树形多选，仅显示有权限的角色）
  - 公共 consts 接口调用（枚举常量统一获取）
```

#### 阶段 E（串行，等待 D 完成）

**E1 - frontend-specialist（动态路由 + 权限）**
```
任务：
  - 配置 Vben Admin 使用 backend 模式权限（accessMode: 'backend'）
  - 实现菜单接口对接（/auth/info 返回用户菜单树）
  - 路由守卫配置
  - 按钮级权限指令（v-access）
  - API 请求拦截器（携带 JWT Token，401 自动刷新）
```

#### 阶段 F（串行）

**F1 - reviewer（全量代码审查）**
```
检查项：
  1. GoFrame 分层是否严格（Controller 不含业务逻辑）
  2. 所有 ID 字段是否使用 JsonInt64
  3. 软删除（deleted_at）是否正确使用 gorm 软删除
  4. JWT 中间件是否覆盖所有需鉴权路由
  5. data_scope 过滤是否在所有列表接口生效
  6. 前端 API 调用是否处理了 BigInt string 转换
  7. 树形接口 children 字段是否正确（空时为 [] 不为 null）
  8. 错误处理是否统一使用 gerror 包
输出：问题清单（Markdown 格式）
```

---

## Orchestrator 完整指令脚本

### Phase 1 启动指令

```
我要构建 GBaseAdmin 代码生成器。请阅读 /www/wwwroot/project/gbaseadmin/架构设计文档.md
作为总指挥，按以下顺序协调工作：

【并行启动 - 3个Agent同时工作】
1. 使用 codegen-parser Agent：
   读取架构文档中的字段规则，在 /www/wwwroot/project/gbaseadmin/codegen/parser/
   实现完整的数据库解析器，包含 TableMeta/FieldMeta 结构体和字段→组件映射规则。

2. 使用 codegen-backend Agent：
   读取架构文档，在 /www/wwwroot/project/gbaseadmin/codegen/templates/backend/
   实现所有后端 Go template 模板文件（api/controller/logic/service/model/consts/router）。

3. 使用 codegen-frontend Agent：
   读取架构文档，在 /www/wwwroot/project/gbaseadmin/codegen/templates/frontend/
   实现所有前端 template 模板文件（api.ts/types.ts/list.vue/form.vue）。

【等待上面3个完成后】
4. 使用 codegen-engine Agent：
   组装 /www/wwwroot/project/gbaseadmin/codegen/ 的生成器引擎和 CLI 入口。

【等待步骤4完成后】
5. 使用 codegen-tester Agent：
   用 dept 表测试生成器全流程，报告验证结果。
```

### Phase 2 启动指令

```
代码生成器已就绪。现在构建权限框架本体。

【先执行初始化】
1. 执行 gf init 命令初始化项目
2. 执行建表 SQL
3. 运行 gf gen dao
4. 运行 codegen 批量生成 dept/role/menu/users

【并行启动 - 3个Agent同时工作】
5. 使用 backend-specialist Agent（任务A）：
   实现 JWT认证、utility工具包、登录接口

6. 使用 backend-specialist Agent（任务B）：
   补充树形业务逻辑、授权逻辑、数据权限中间件

7. 使用 frontend-specialist Agent（任务A）：
   实现特殊UI组件（授权弹窗、树形选择器等）

【等待上面完成后】
8. 使用 frontend-specialist Agent（任务B）：
   实现动态路由、权限指令、请求拦截器

【最后】
9. 使用 reviewer Agent：
   全量代码审查，输出问题清单
```

---

## 开启 Agent Team 实验模式

在项目根目录创建 `.env` 或在启动时设置：

```bash
export CLAUDE_CODE_EXPERIMENTAL_AGENT_TEAMS=1
claude
```

或在 `~/.claude/settings.json` 中添加：

```json
{
  "env": {
    "CLAUDE_CODE_EXPERIMENTAL_AGENT_TEAMS": "1"
  }
}
```

---

## 注意事项

1. **子 Agent 不能再生成子 Agent**，嵌套只能一层
2. **worktree 隔离**：并行 Agent 编写不同文件时不会冲突
3. **每次 Agent 调用都是全新上下文**，关键背景信息要在 prompt 中带入
4. **Agent 返回结果不会显示给用户**，orchestrator 需主动汇总后告知
5. **代码生成器 Agent** 每次调用都应传入架构文档路径，确保生成规则一致

---

## 执行进度记录

| 日期 | 阶段 | 任务 | 状态 | 备注 |
|---|---|---|---|---|
| 2026-03-29 | Phase 0 | DevOps 环境搭建 | ✅ 完成 | docker-compose、deploy/、.env.example、Makefile |
| 2026-03-29 | Phase 0 | Agent 定义文件 | ✅ 完成 | 9 个 Agent 全部创建于 .claude/agents/ |
| 2026-03-29 | Phase 1-A1 | codegen-parser | ✅ 完成 | parser/meta.go、parser.go、comment_parser.go、field_mapper.go |
| 2026-03-29 | Phase 1-A2 | codegen-backend 模板 | ✅ 完成 | 7 个 .tpl（api/controller/logic/service/model/consts/router） |
| 2026-03-29 | Phase 1-A3 | codegen-frontend 模板 | ✅ 完成 | 4 个 .tpl（types/api/list/form） |
| 2026-03-29 | 修复 | app/system 目录结构 | ✅ 完成 | api-template → system，修复 import 路径 |
| 2026-03-29 | Phase 1-B | codegen-engine 引擎组装 | ✅ 完成 | generator/backend、generator/frontend、main.go、config.go、codegen.yaml |
| 2026-03-29 | Phase 1-C | codegen-tester 验证 | ✅ 完成 | dept/role/menu/users 4 表 44 文件全部生成成功 |
| 2026-03-29 | 优化 | 代码生成器迭代 | ✅ 完成 | 软删除、查询条件透传、密码bcrypt、关联字段显示、空children、NameDao双命名 |
| 2026-03-29 | Phase 2-Init | 项目初始化 | ✅ 完成 | gf gen dao、路由注册、go mod tidy 编译通过 |
| 2026-03-29 | Phase 2-D1 | JWT 认证 + 中间件 + 工具包 | ✅ 完成 | jwt/response 工具包、auth 登录/信息/改密、鉴权中间件、路由分组 |
