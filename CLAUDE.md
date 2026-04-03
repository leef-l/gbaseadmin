# GBaseAdmin AI 协作说明

这份文件不是产品介绍，而是给后续 AI / 自动化代理的高密度工作说明。目标只有一个：**让任何新来的 AI 在最短时间内明白这套系统应该从哪里读、哪里改、哪里不要乱动。**

---

## 1. 一句话认识项目

GBaseAdmin 是一套围绕“陪玩平台”构建的全栈系统：

- 后端：`admin-go/`，GoFrame v2 MonoRepo
- 管理端：`vue-vben-admin/apps/web-antd/src/`
- C 端：`wap-ui/`
- CRUD 生成器：`admin-go/codegen/`

**这不是单体仓库里的单应用，而是“后端多服务 + 管理后台 + WAP 前端 + 生成器工具”的组合系统。**

---

## 2. 先读哪里

如果你刚接手任务，优先阅读顺序如下：

1. `README.MD`
2. `admin-go/app/play/internal/cmd/cmd.go`
3. `admin-go/app/system/internal/cmd/cmd.go`
4. `admin-go/app/upload/internal/cmd/cmd.go`
5. `admin-go/codegen/`
6. `vue-vben-admin/apps/web-antd/src/`

原因：

- `README.MD` 负责讲清楚系统地图
- `cmd.go` 系列直接告诉你每个服务暴露了哪些路由
- `codegen/` 决定了这套系统很多 CRUD 的“正确改法”
- 前端业务代码只在 `apps/web-antd/src/`，不是整个 Vben Monorepo 到处都能动

---

## 3. 真实源码边界

### 主要工作区

```text
admin-go/                    # 后端源码
vue-vben-admin/apps/web-antd/src/  # 管理端业务源码
wap-ui/                      # C 端源码
admin-go/codegen/            # 代码生成器
文档专区/                     # 设计/接口/部署文档
```

### 非核心或谨慎处理区

```text
vue-vben-admin/packages/
vue-vben-admin/internal/
vue-vben-admin/playground/
wap-ui-prototype/
admin-go/app/*/dao/
admin-go/app/*/internal/model/do/
admin-go/app/*/internal/model/entity/
admin-go/codegen/sql/init_backup_*.sql
```

说明：

- `vue-vben-admin/` 是上游 Vben Admin Monorepo，**真正频繁修改的是 `apps/web-antd/src/`**
- `dao/do/entity` 是生成结果，不要把它们当业务逻辑层改
- `wap-ui-prototype/` 是静态原型，不是正式业务前端

---

## 4. 后端结构速记

### 三个正式服务

| 服务 | 端口 | 前缀 | 职责 |
| --- | --- | --- | --- |
| `system` | `8000` | `/api/system` | 登录、用户、角色、部门、菜单、权限 |
| `play` | `8001` | `/api/play`、`/api/playapi` | 陪玩核心业务 |
| `upload` | `8002` | `/api/upload` | 文件/目录/上传配置 |

### 标准分层

```text
app/{app}/
├── api/{app}/v1/
├── internal/cmd/
├── internal/controller/
├── internal/service/
├── internal/logic/
├── internal/model/
├── internal/consts/
├── internal/middleware/
├── internal/dao/           # 生成
├── internal/model/do/      # 生成
├── internal/model/entity/  # 生成
└── main.go
```

### 路由判断方法

- 是否是管理端接口：看 `cmd.go` 里是否在 `/api/system` 或 `/api/play`、`/api/upload`
- 是否是 C 端接口：看 `play/internal/cmd/cmd.go` 里的 `/api/playapi`
- 是否要求登录：看挂载的中间件是 `Auth`、`MemberAuth`、`MemberAuthOptional` 还是 `CoachOnly`

---

## 5. play 服务是核心中的核心

绝大多数复杂业务都在 `admin-go/app/play/`。

### 这里有两套 API

- `/api/play`：后台管理 API
- `/api/playapi`：C 端 API

### `/api/playapi` 又分三层

- 公开接口：可不登录
- 会员接口：必须会员身份
- 陪玩师接口：必须陪玩师身份

只要任务涉及下面这些关键词，优先去 `play`：

- 会员
- 陪玩师
- 商品
- 店铺
- 订单
- 支付
- 充值
- 活动
- 优惠券
- 消息
- 提现
- 评价

---

## 6. 前端结构速记

### 管理端只看这里

```text
vue-vben-admin/apps/web-antd/src/
├── api/
├── views/
├── components/
├── adapter/
├── router/routes/modules/
├── store/
└── locales/
```

### 管理端改动的经验法则

- 改接口：优先去 `src/api/{system|play|upload}/`
- 改页面：优先去 `src/views/{system|play|upload}/`
- 改自定义上传、文件选择、富文本、JSON 编辑：去 `src/components/`
- 改表单组件注册：去 `src/adapter/component/index.ts`
- 改路由：去 `src/router/routes/modules/`

### 管理端通用约定

- 列表页通常用 `useVbenVxeGrid`
- 表单通常用 VbenForm schema
- 图片上传字段优先 `ImageUpload`
- 文件上传字段优先 `FileUpload`
- 文件选择优先 `FileManagerModal`

---

## 7. WAP 端速记

`wap-ui/` 是真实业务前端，技术栈为 Taro + React + NutUI React Taro。

不要把 `wap-ui-prototype/` 误判成正式工程。

任务如果涉及：

- 小程序
- H5
- 首页/分类/活动/消息/我的
- C 端下单
- 陪玩师工作台前台体验

就优先去 `wap-ui/src/`。

---

## 8. 这套系统最重要的“铁律”

### 铁律 1：生成器优先

如果问题出在生成的 CRUD 上：

- **先看 `admin-go/codegen/templates/`**
- **先看 `admin-go/codegen/parser/`**
- **先看 `admin-go/codegen/generator/`**

不要一边说“这是生成代码”，一边去每个模块手修同样的问题。

### 铁律 2：不要把生成层当业务层

这些目录通常不应手工维护业务逻辑：

- `admin-go/app/*/internal/dao/`
- `admin-go/app/*/internal/model/do/`
- `admin-go/app/*/internal/model/entity/`

### 铁律 3：ID 一律注意精度

后端很多 ID 使用 `snowflake.JsonInt64`，JSON 序列化为字符串。

处理前后端交互时：

- 不要假设 ID 是普通 number
- 不要随手把字符串 ID 改成数值型

### 铁律 4：响应格式统一

默认响应契约：

```json
{ "code": 0, "message": "ok", "data": {} }
```

管理端请求层默认按这个结构解析。

### 铁律 5：权限与菜单不要只改前端

管理端权限通常同时涉及：

- 后端接口
- 后端菜单/权限码
- 前端页面/按钮显隐
- `system_menu` 数据

如果是生成模块，菜单通常应由 codegen 统一生成或更新。

---

## 9. 高频入口文件

### 后端

- `admin-go/app/system/internal/cmd/cmd.go`
- `admin-go/app/play/internal/cmd/cmd.go`
- `admin-go/app/upload/internal/cmd/cmd.go`
- `admin-go/utility/jwt/jwt.go`
- `admin-go/utility/snowflake/snowflake.go`
- `admin-go/utility/response/response.go`

### 代码生成器

- `admin-go/codegen/main.go`
- `admin-go/codegen/codegen.yaml`
- `admin-go/codegen/parser/`
- `admin-go/codegen/generator/`
- `admin-go/codegen/templates/`
- `admin-go/codegen/CHANGELOG.md`

### 管理端

- `vue-vben-admin/apps/web-antd/src/api/request.ts`
- `vue-vben-admin/apps/web-antd/src/adapter/component/index.ts`
- `vue-vben-admin/apps/web-antd/src/router/routes/modules/`
- `vue-vben-admin/apps/web-antd/src/views/`

### WAP

- `wap-ui/package.json`
- `wap-ui/config/`
- `wap-ui/src/`

---

## 10. 常用命令

### 后端编译

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
pnpm run build:weapp
```

### 生成器

```bash
cd admin-go/codegen
go run . -table play_order -force -menu
go run . -table play_order -only backend -force
go run . -table play_order -only frontend -force
go run . -table play_order -only menu
```

---

## 11. 部署认知

根目录部署脚本：

- `deploy.ps1`：WSL + rsync
- `deploy-fast.ps1`：并行编译 + hash 缓存 + scp 压缩
- `deploy-scp.ps1`：纯 scp
- `deploy.bat`：菜单入口

`admin-go/docker-compose.yml` 是容器化开发/部署入口，包含：

- MySQL
- Redis
- system / play / upload
- frontend
- wap
- adminer
- gotools

脚本运行时会临时使用根目录 `dist/`，该目录是产物，不是源码。

---

## 12. 文档使用建议

### 先看系统总览

- `README.MD`

### 需要业务设计

- `文档专区/陪玩项目设计文档.md`
- `文档专区/陪玩核心业务流程设计.md`
- `文档专区/架构设计文档.md`

### 需要接口边界

- `文档专区/陪玩后端管理端接口设计.md`
- `文档专区/陪玩C端接口设计.md`

### 需要页面设计

- `文档专区/陪玩前端管理端页面设计.md`
- `文档专区/陪玩WAP端页面设计.md`

### 需要部署

- `文档专区/宝塔部署教程.md`
- `文档专区/Windows一键部署使用文档.md`

---

## 13. 给后续 AI 的执行建议

1. 先确认自己改的是哪一层：后端 / 管理端 / WAP / 生成器。
2. 如果是 CRUD 异常或批量重复问题，先怀疑生成器，不要先批量手改。
3. 如果是管理端页面，尽量只在 `apps/web-antd/src/` 内活动。
4. 如果是 C 端接口，先判断它走 `/api/playapi` 的哪种认证层级。
5. 如果改动涉及权限、菜单、按钮显隐，默认要同时检查前后端和菜单数据。
6. 如果看见 `dao/do/entity`，先停一下，确认是不是在错误层改代码。

这份文件的目标不是事无巨细，而是帮助你**第一时间走对方向**。
