# GBaseAdmin

GBaseAdmin 是一套围绕陪玩业务构建的全栈管理系统，包含：

- `admin-go/`：GoFrame v2 MonoRepo 后端
- `vue-vben-admin/`：Vue 3 + Vben Admin 管理端
- `wap-ui/`：Taro + React 的 C 端 WAP/小程序前端
- `admin-go/codegen/`：从数据库表生成后端 CRUD、管理端页面和菜单的代码生成器

这份 README 负责回答三个问题：

1. 这套系统由哪些真实模块组成
2. 日常开发应该从哪里进入
3. 哪些目录是源码，哪些只是文档、原型、部署脚本或产物

## 1. 仓库总览

### 真实源码目录

```text
gbaseadmin/
├── admin-go/          # 后端 MonoRepo：system / play / upload + utility + codegen
├── vue-vben-admin/    # 管理端前端，实际业务代码在 apps/web-antd/src/
├── wap-ui/            # C 端 WAP / 微信小程序前端
└── 文档专区/           # 设计、接口、部署文档
```

### 辅助目录

```text
├── wap-ui-prototype/  # 早期静态原型，仅供视觉/交互参考
├── deploy*.ps1/.bat   # Windows 部署脚本
└── CLAUDE.md          # 面向 AI 代理的高密度项目说明
```

### 已确认不是主源码入口的内容

- 顶层 `play/`、`system/` 已清理：它们是 stray 构建产物，不是源码目录。
- 顶层 `upload/` 仍保留 `resource/`，更像运行期上传残留，不属于主源码。
- `dist/`、`deploy-logs/`、历史 `gbaseadmin_*.tar.gz` 属于部署产物/日志，可再生。

## 2. 技术栈

### 后端

- 语言：Go（`admin-go/go.mod` 当前声明 `go 1.25.0`）
- 框架：GoFrame v2.10
- 数据库：MySQL 8.0
- 缓存：Redis 7
- 认证：自定义 JWT（`admin-go/utility/jwt`）
- ID：雪花 ID（`admin-go/utility/snowflake`，统一使用 `snowflake.JsonInt64`）

### 管理端

- 框架：Vue 3
- 后台框架：Vben Admin v5.7
- UI：Ant Design Vue
- 表格：VxeTable（统一通过 `useVbenVxeGrid`）
- 表单：VbenForm
- 包管理：pnpm workspace + Turbo

### WAP / 小程序端

- 框架：Taro 4.x + React 18
- UI：NutUI React Taro
- 状态管理：Zustand
- 构建目标：H5 / 微信小程序 / 头条小程序

## 3. 后端架构

后端位于 `admin-go/`，是标准 GoFrame MonoRepo。

### 应用划分

| 应用 | 端口 | 路由前缀 | 作用 |
| --- | --- | --- | --- |
| `system` | `8000` | `/api/system` | 登录、用户、角色、部门、菜单、权限 |
| `play` | `8001` | `/api/play` / `/api/playapi` | 陪玩业务核心：会员、商品、订单、支付、活动、评价、提现等 |
| `upload` | `8002` | `/api/upload` | 文件、目录、上传配置、上传入口 |
| `svc-template` | - | - | 新 GoFrame 服务模板 |
| `job-template` | - | - | 定时任务模板 |

### 单个应用的标准结构

```text
admin-go/app/{app}/
├── api/{app}/v1/          # API 请求/响应结构体，靠 g.Meta 注册路由
├── internal/
│   ├── cmd/               # 路由分组和服务启动
│   ├── controller/        # 控制器，保持薄层
│   ├── service/           # 接口定义
│   ├── logic/             # 业务实现
│   ├── model/             # DTO
│   ├── consts/            # 枚举常量
│   ├── middleware/        # 鉴权等中间件
│   ├── dao/               # gf gen dao 生成
│   └── model/do|entity/   # gf gen dao 生成
├── manifest/              # 运行配置、部署资源
└── main.go                # 应用入口
```

### 三个关键公共包

- `admin-go/utility/jwt`：管理端 JWT + 会员端 JWT 双体系
- `admin-go/utility/snowflake`：雪花 ID 与 `JsonInt64`
- `admin-go/utility/response`：统一响应 `{ code, message, data }`

### play 服务的特殊性

`play` 同时服务管理端和 C 端：

- `/api/play`：后台管理 API，统一走 `middleware.Auth`
- `/api/playapi`：C 端 API
  - 公开接口：`MemberAuthOptional`
  - 会员接口：`MemberAuth`
  - 陪玩师接口：`CoachOnly`

如果涉及会员端、订单、支付、陪玩师工作台，优先从 `admin-go/app/play/internal/cmd/cmd.go` 和 `internal/controller/playapi/` 入手。

## 4. 管理端前端架构

前端仓库位于 `vue-vben-admin/`，但**真正需要频繁修改的业务代码主要集中在**：

```text
vue-vben-admin/apps/web-antd/src/
```

### 核心目录

```text
src/
├── api/                   # 请求封装，按 system / play / upload 分目录
├── views/                 # 页面，按 system / play / upload 分目录
├── components/            # 自定义业务组件
├── adapter/               # VbenForm / VxeTable / 自定义组件注册
├── router/routes/modules/ # 路由模块
├── store/                 # Pinia 状态
└── locales/               # 国际化
```

### 重要约定

- 统一响应格式：`{ code: 0, message: 'ok', data: ... }`
- 登录流程：
  - `POST /system/auth/login`
  - `GET /system/auth/info`
  - `GET /system/auth/menus`
- 表格统一优先沿用 `useVbenVxeGrid`
- 表单统一优先沿用 VbenForm
- 上传字段优先使用已有组件：
  - `ImageUpload`
  - `FileUpload`
  - `FileManagerModal`

## 5. WAP 端架构

`wap-ui/` 是实际运行的 C 端项目，不是演示壳。

### 特征

- 页面数量较多，覆盖首页、分类、活动、消息、我的等完整流程
- 使用 Taro 4.x，目标同时覆盖 H5 与小程序
- 包管理和命令以项目自身 `package.json` 为准

### 常用命令

```bash
cd wap-ui
pnpm install
pnpm run dev:h5
pnpm run build:h5
pnpm run build:weapp
```

`wap-ui-prototype/` 只是静态原型页面，不参与正式构建和部署。

## 6. 代码生成器

代码生成器位于 `admin-go/codegen/`，是这套系统的高频核心工具。

### 它会生成什么

- 后端：API / Controller / Logic / Service / Model / Consts
- 前端：API / Types / List / Form / Detail Drawer
- 菜单：直接写入 `system_menu` 表

### 基本命令

```bash
cd admin-go/codegen

# 后端 + 前端 + 菜单
go run . -table play_order -force -menu

# 只生成后端
go run . -table play_order -only backend -force

# 只生成前端
go run . -table play_order -only frontend -force

# 只生成菜单
go run . -table play_order -only menu
```

### 表命名约定

- 表名格式：`{应用}_{模块}`
- 例如：
  - `system_menu`
  - `upload_file`
  - `play_order`

生成器会用表名前缀决定输出到哪个 app/前端模块目录。

### 最重要的规则

- 生成代码有问题，**优先修模板，再重新生成**
- 不要把生成器问题直接手改到每个业务文件里
- `dao/`、`model/do/`、`model/entity/` 属于生成内容，不要随意手工维护
- 生成菜单后，记得同步数据库初始化 SQL 到 `admin-go/codegen/sql/init.sql`

## 7. 本地开发入口

### 后端

```bash
cd admin-go
go build ./app/system/...
go build ./app/play/...
go build ./app/upload/...
```

### 管理端

```bash
cd vue-vben-admin
pnpm install
pnpm dev:antd
pnpm build:antd
```

### WAP

```bash
cd wap-ui
pnpm install
pnpm run dev:h5
pnpm run build:h5
```

## 8. Docker / 部署

### Docker Compose

`admin-go/docker-compose.yml` 会拉起：

- MySQL
- Redis
- system / play / upload 三个后端服务
- frontend 开发容器
- wap 开发容器
- adminer
- gotools 工具容器

国内环境还有：

- `admin-go/docker-compose.cn.yml`
- `admin-go/docker-compose.prod.yml`

### Windows 部署脚本

根目录提供三套部署脚本：

- `deploy.ps1`：WSL + rsync 版
- `deploy-fast.ps1`：并行编译 + hash 缓存 + scp 压缩版
- `deploy-scp.ps1`：纯 scp 版
- `deploy.bat`：交互式入口

这些脚本默认会构建并临时使用根目录 `dist/`，结束后会自动清理。

## 9. 文档索引

### 设计与接口文档

位于 `文档专区/`：

- `架构设计文档.md`
- `陪玩项目设计文档.md`
- `陪玩核心业务流程设计.md`
- `陪玩后端管理端接口设计.md`
- `陪玩前端管理端页面设计.md`
- `陪玩WAP端页面设计.md`
- `陪玩C端接口设计.md`
- `宝塔部署教程.md`
- `Windows一键部署使用文档.md`

## 10. 给新接手开发者的建议

如果你刚接手这套系统，建议按下面顺序读：

1. 本 README：先搞清楚真实模块边界
2. `CLAUDE.md`：看 AI/协作约定与高密度规则
3. `admin-go/app/play/internal/cmd/cmd.go`：理解最大业务域的路由结构
4. `admin-go/codegen/`：理解为什么很多 CRUD 不应该手写
5. `vue-vben-admin/apps/web-antd/src/`：只在业务代码目录内工作

如果你是 AI 代理，优先把 `CLAUDE.md` 当作操作说明，把本 README 当作系统地图。
