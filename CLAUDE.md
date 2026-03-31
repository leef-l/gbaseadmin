# GBaseAdmin 项目指令

## 技术栈

### 后端
- **语言**: Go 1.23
- **框架**: GoFrame v2.10（MonoRepo 多应用架构）
- **数据库**: MySQL 8.0 + Redis
- **认证**: JWT（自定义实现，`utility/jwt`）
- **ID策略**: Snowflake 雪花ID（`utility/snowflake`，类型 `snowflake.JsonInt64`）
- **部署**: Docker Compose

### 管理端前端
- **框架**: Vue 3 + Vben Admin v5.7（`vue-vben-admin/apps/web-antd`）
- **UI库**: Ant Design Vue
- **表格**: VxeTable（通过 `useVbenVxeGrid`）
- **表单**: VbenForm（支持自定义组件 `ImageUpload`/`FileUpload`）
- **构建**: Vite + pnpm

### WAP端（C端）
- **框架**: Taro 4.x + React
- **UI库**: NutUI React Taro
- **目标**: 微信小程序 + H5

### 代码生成器
- **位置**: `admin-go/codegen/`
- **语言**: Go（独立工具，非 GoFrame 应用）
- **模板**: Go `text/template`（`codegen/templates/`）

---

## 项目目录结构

```
gbaseadmin/
├── admin-go/                     # 后端（Go MonoRepo）
│   ├── app/                      # 应用目录（每个应用独立 GoFrame 服务）
│   │   ├── system/               # 系统管理（用户、角色、部门、菜单）
│   │   ├── upload/               # 上传管理（文件上传、目录、配置）
│   │   ├── play/                 # 陪玩业务
│   │   ├── svc-template/         # 新应用模板
│   │   └── job-template/         # 定时任务模板
│   ├── codegen/                  # 代码生成器
│   │   ├── main.go               # CLI 入口
│   │   ├── codegen.yaml           # 数据库连接和输出配置
│   │   ├── parser/               # 表结构解析器
│   │   ├── generator/            # 生成器（backend/frontend/menu）
│   │   ├── templates/            # 模板文件
│   │   │   ├── backend/          # 后端模板（api/controller/logic/model/service/...）
│   │   │   └── frontend/         # 前端模板（api/types/list/form）
│   │   └── sql/                  # 数据库初始化SQL
│   ├── utility/                  # 公共工具
│   │   ├── jwt/                  # JWT 工具
│   │   ├── snowflake/            # 雪花ID生成器
│   │   └── response/             # 统一响应
│   └── deploy/                   # 部署配置
│
├── vue-vben-admin/               # 管理端前端
│   └── apps/web-antd/src/
│       ├── api/                  # API 调用层（按应用/模块分目录）
│       ├── views/                # 页面（按应用/模块分目录）
│       ├── components/           # 自定义组件
│       │   ├── file-manager/     # 文件管理器（弹窗选择文件/图片）
│       │   ├── upload/           # 上传组件（ImageUpload/FileUpload）
│       │   ├── tinymce/          # 富文本编辑器
│       │   └── json-editor/      # JSON编辑器
│       ├── router/routes/modules/ # 路由模块
│       └── adapter/              # VbenForm 组件适配（注册自定义组件）
│
├── wap-ui/                       # WAP端（Taro + React）
└── wap-ui-prototype/             # WAP端原型页面
```

### 后端应用内部结构（GoFrame 标准分层）

```
app/{应用名}/
├── api/{应用名}/v1/     # API 请求/响应结构体（路由注册靠 g.Meta tag）
├── internal/
│   ├── cmd/cmd.go       # 路由注册和服务启动
│   ├── middleware/       # 中间件（JWT鉴权等）
│   ├── controller/      # 控制器（薄层，直接调用 service）
│   ├── logic/           # 业务逻辑实现
│   ├── service/         # 接口定义（IoC）
│   ├── model/           # DTO（Input/Output 模型）
│   ├── consts/          # 枚举常量
│   ├── dao/             # ORM DAO（gf gen dao 自动生成）
│   └── model/do/entity/ # DO/Entity（gf gen dao 自动生成）
└── main.go              # 入口（注册所有 logic init）
```

---

## 代码生成器使用

### 命令格式
```bash
cd admin-go/codegen
go run . -table <表名> [选项]
```

### 常用命令
```bash
# 生成完整代码（后端 + 前端 + 菜单）
go run . -table play_order -force -menu

# 只生成后端
go run . -table play_order -only backend -force

# 只生成前端
go run . -table play_order -only frontend -force

# 只生成菜单
go run . -table play_order -only menu

# 批量生成
go run . -table play_order,play_payment,play_review -force -menu

# 预览（不实际生成）
go run . -table play_order -dry-run
```

### 表名约定
- 表名格式: `{应用}_{模块}`，如 `play_order`、`upload_file`、`system_menu`
- 应用前缀决定生成到哪个 app 目录
- 生成器自动处理: API/Controller/Logic/Service/Model/Consts + 前端 API/Types/List/Form + 数据库菜单

### 生成后注意
- 生成的代码是**基础 CRUD 脚手架**
- **自定义业务逻辑**（如文件上传、树形结构、物理删除）需在生成后手动修改
- 重新生成会覆盖文件（`-force`），手动修改的代码需要备份或重新添加
- DAO/DO/Entity 由 `gf gen dao` 自动生成，不要手动修改
- **每次生成菜单后，必须同步数据库到 `admin-go/codegen/sql/init.sql` 并生成一个日期备份**：
  ```bash
  mysqldump -u root -pgbaseadmin123 -h 127.0.0.1 gbaseadmin --routines --triggers --single-transaction > admin-go/codegen/sql/init.sql
  cp admin-go/codegen/sql/init.sql admin-go/codegen/sql/init_backup_$(date +%Y%m%d%H%M).sql
  ```

---

## 开发流程（新增业务模块）

### 第一步：设计数据库表
- 在 MySQL 中创建表，遵循命名约定 `{应用}_{模块}`
- 公共字段: `id`(bigint PK), `created_at`, `updated_at`, `deleted_at`(软删除), `created_by`, `dept_id`

### 第二步：代码生成
```bash
cd admin-go/codegen
go run . -table {表名} -force -menu
```
- 自动生成后端全套代码 + 前端页面 + 菜单数据

### 第三步：自定义业务逻辑
- 修改 `logic/` 中的业务实现（生成的只是基础 CRUD）
- 修改 `api/` 中的请求结构体（添加特殊参数）
- 修改前端表单/列表（调整字段、添加自定义组件）

### 第四步：编译验证
```bash
cd admin-go/app/{应用}
go build ./...
```

### 第五步：提交代码
- 每完成一个任务立即 `git commit + push`
- 提交信息用中文，描述清楚改了什么

---

## 重要约定

### 后端约定
- 所有 ID 字段使用 `snowflake.JsonInt64` 类型（前端 JSON 序列化为字符串避免精度丢失）
- 软删除: `deleted_at` 字段，查询时加 `WHERE deleted_at IS NULL`
- 数据隔离: `dept_id` + `created_by` 字段
- 密码加密: SHA256（`gsha256.Encrypt`）

### 前端约定
- 表单图片字段用 `component: 'ImageUpload'`
- 表单文件字段用 `component: 'FileUpload'`
- 上传统一走文件管理器弹窗（`FileManagerModal`）
- API 前缀: `/{应用名}/{模块名}`（如 `/play/order`）

### 代码生成器铁律
- **生成的代码有问题，先修生成器模板再重新生成**，不手写修复
- 生成器模板在 `codegen/templates/` 下

---

## 常用命令

### 后端
```bash
# 编译某个应用
cd admin-go/app/system && go build ./...

# 编译所有应用
cd admin-go && go build ./app/system/... ./app/upload/... ./app/play/...

# 生成 DAO（通常由 codegen 自动调用）
cd admin-go/app/{应用} && gf gen dao
```

### 前端
```bash
cd vue-vben-admin
pnpm install
pnpm dev        # 开发服务
pnpm build      # 生产构建
```

### 数据库
```bash
mysql -u root -pgbaseadmin123 -h 127.0.0.1 -P 3306 gbaseadmin

# 导出数据库
mysqldump -u root -pgbaseadmin123 -h 127.0.0.1 gbaseadmin > admin-go/codegen/sql/init.sql
```
