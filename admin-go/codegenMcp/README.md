# codegenMcp

`codegenMcp` 是 `admin-go/codegen` 的独立 STDIO MCP Server 版本。

它复用了原有的 parser / generator / templates 能力，但把 CLI 入口替换成了 MCP 工具接口，方便 Claude Desktop、Cursor、VS Code 等 MCP 客户端直接调用。

## 目标

- 独立模块：`module gbaseadmin/codegenMcp`
- 通信方式：STDIO MCP
- 工具能力：
  - `inspect_table`
  - `preview_generate`
  - `generate_code`
- 明确禁用：
  - `gf init`
  - `gf gen dao`
  - `exec.Command` 外部命令链路

## 构建

```bash
cd admin-go/codegenMcp
go mod tidy
go build ./...
```

Windows 可执行文件示例：

```bash
go build -o codegenMcp.exe .
```

## 启动

默认读取当前目录下的 `codegen.yaml`，并自动读取上一级 `admin-go/.env`：

```bash
codegenMcp.exe
```

也可以显式指定配置文件：

```bash
codegenMcp.exe --config C:/project/gbaseadmin/admin-go/codegenMcp/codegen.yaml
```

或者使用环境变量：

```bash
set CODEGEN_CONFIG=C:/project/gbaseadmin/admin-go/codegenMcp/codegen.yaml
codegenMcp.exe
```

## MCP 客户端配置示例

### Claude Desktop

```json
{
  "mcpServers": {
    "codegenMcp": {
      "command": "C:/project/gbaseadmin/admin-go/codegenMcp/codegenMcp.exe",
      "args": [
        "--config",
        "C:/project/gbaseadmin/admin-go/codegenMcp/codegen.yaml"
      ]
    }
  }
}
```

### Cursor / VS Code

```json
{
  "servers": {
    "codegenMcp": {
      "type": "stdio",
      "command": "C:/project/gbaseadmin/admin-go/codegenMcp/codegenMcp.exe",
      "args": [
        "--config",
        "C:/project/gbaseadmin/admin-go/codegenMcp/codegen.yaml"
      ]
    }
  }
}
```

## MCP Tools

### 1. inspect_table

解析数据库表结构，返回字段摘要信息。

参数：

- `table: string` 必填，例：`system_dept`

返回内容包括：

- 表名、应用名、模块名、注释
- 字段数量
- 每个字段的 Go 类型、TS 类型、组件类型、是否必填、是否枚举

### 2. preview_generate

预览指定表会生成哪些文件，不实际写入磁盘。

参数：

- `table: string` 必填
- `only: string` 可选，`backend | frontend | menu`

返回：

- 文件路径列表
- 每个文件的字节数
- 生成错误列表（如果存在）

### 3. generate_code

真正执行代码生成并写入磁盘。

参数：

- `table: string` 必填，支持逗号分隔多个表
- `only: string` 可选，`backend | frontend | menu`
- `force: boolean` 可选，是否覆盖已有文件
- `with_menu: boolean` 可选，是否同时向数据库写入菜单

返回：

- 每张表的生成文件列表
- 菜单写入数量
- 每张表耗时
- 错误列表

## 配置文件

示例配置见 `codegen.yaml`，数据库连接建议统一来自 `admin-go/.env`。

关键字段：

- `database`: MySQL 连接配置
- `backend.output`: 后端输出根目录
- `frontend.output`: 前端输出根目录
- `skip_fields`: 跳过生成的公共字段
- `menu_apps`: 菜单应用配置
- `menu_modules`: 菜单模块配置
- `operation_log`: 是否生成操作日志相关代码
- `post_generation.run_gf_init`: 默认 `false`
- `post_generation.run_gf_gen_dao`: 默认 `false`

> 注意：即使配置里存在 `post_generation`，MCP 版本也不会执行 `gf init` 或 `gf gen dao`。
