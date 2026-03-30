# GBaseAdmin 代码生成器

## 表名规范

所有数据库表必须使用 `{应用名}_{模块名}` 格式命名：

| 表名 | 应用 | 模块 |
|------|------|------|
| `system_dept` | system | dept |
| `system_role` | system | role |
| `system_menu` | system | menu |
| `system_users` | system | users |
| `system_user_role` | system | user_role |
| `system_role_menu` | system | role_menu |
| `system_role_dept` | system | role_dept |
| `system_user_dept` | system | user_dept |
| `demo_demo` | demo | demo |

代码生成器根据第一个 `_` 拆分表名，前半部分为应用名，后半部分为模块名。

## 使用方法

```bash
cd admin-go/codegen

# 生成单表
go run . --table system_dept

# 生成多表
go run . --table system_dept,system_role

# 只生成后端
go run . --table system_dept --only backend

# 只生成前端
go run . --table system_dept --only frontend

# 强制覆盖已有文件
go run . --table system_dept --force

# 预览（不写入文件）
go run . --table system_dept --dry-run
```

## 自动创建应用

当表名前缀对应的应用目录不存在时，代码生成器会自动执行 `gf init app/{appName} -a` 创建应用骨架。

例如：`--table demo_demo` 会自动创建 `app/demo/` 应用目录。

## 生成文件列表

### 后端（输出到 `app/{appName}/`）

| 模板 | 输出路径 | 说明 |
|------|---------|------|
| `api.tpl` | `api/{app}/v1/{module}.go` | API 请求/响应结构体 |
| `controller.tpl` | `internal/controller/{module}/{module}.go` | 控制器 |
| `logic.tpl` | `internal/logic/{module}/{module}.go` | 业务逻辑 |
| `service.tpl` | `internal/service/{module}.go` | 服务接口 |
| `model.tpl` | `internal/model/{module}.go` | DTO 模型 |
| `consts.tpl` | `internal/consts/{module}.go` | 枚举常量 |

### 前端（输出到 `vue-vben-admin/apps/web-antd/src/`）

| 模板 | 输出路径 | 说明 |
|------|---------|------|
| `types.tpl` | `api/{app}/{module}/types.ts` | TypeScript 类型定义 |
| `api.tpl` | `api/{app}/{module}/index.ts` | API 请求函数 |
| `list.tpl` | `views/{app}/{module}/index.vue` | 列表页面 |
| `form.tpl` | `views/{app}/{module}/modules/form.vue` | 表单弹窗 |

## 配置文件

`codegen.yaml`：

```yaml
database:
  host: 127.0.0.1
  port: 9306
  user: root
  password: "gbaseadmin123"
  dbname: gbaseadmin

backend:
  output: ../app/          # 后端输出根目录

frontend:
  output: ../../vue-vben-admin/apps/web-antd/src/

skip_fields:               # 跳过生成的公共字段
  - created_at
  - updated_at
  - deleted_at
  - created_by
  - dept_id
```

## 数据库表设计规范

- 主键统一使用 `id BIGINT UNSIGNED`（Snowflake ID）
- 软删除使用 `deleted_at DATETIME`
- 公共字段：`created_at`、`updated_at`、`deleted_at`、`created_by`、`dept_id`
- 树形结构使用 `parent_id BIGINT UNSIGNED`
- 状态字段使用 `status TINYINT(1)`，注释格式：`状态:0=关闭,1=开启`
- 枚举字段注释格式：`字段说明:值1=标签1,值2=标签2`
- 外键字段命名：`{关联模块}_id`（如 `dept_id`、`role_id`）
- 多选外键字段命名：`{关联模块}_ids`（如 `role_ids`）
