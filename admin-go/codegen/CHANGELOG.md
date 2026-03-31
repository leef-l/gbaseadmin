# Codegen 更新日志

## v1.3.0 — 2026-03-31

### Tooltip 括号语法

字段注释中的中文括号 `（）` 或英文括号 `()` 自动提取为 Tooltip 提示，括号前文字作为精简标签。

**示例：** `排序（升序）` → 标签显示"排序"，鼠标悬停显示"升序"

**变更文件：**

- `parser/meta.go` — `FieldMeta` 新增 `ShortLabel`、`TooltipText`；`TableMeta` 新增 `HasTooltip`
- `parser/comment_parser.go` — 新增 `extractParentheses()` 函数，`ParseComment()` 返回值扩展为 4 个
- `parser/parser.go` — `buildFieldMeta()` 适配新返回值，自动检测 `HasTooltip`
- `templates/frontend/form.tpl` — 表单 label 条件渲染 Tooltip + QuestionCircleOutlined 图标
- `templates/frontend/list.tpl` — 列头使用 `ShortLabel`，有提示时渲染 `slots.header` Tooltip

**生成效果：**

```vue
<!-- 无括号：普通文字 -->
label: '部门名称'

<!-- 有括号：Tooltip 渲染 -->
label: () => h('span', {}, ['排序 ', h(Tooltip, { title: '升序' }, { default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }) })])
```

---

## v1.2.0 — 2026-03-28

### 菜单生成器

新增 `--menu` 和 `--only menu` 参数，支持将菜单数据直接写入 `system_menu` 表。

每个模块自动生成目录 + 页面 + 按钮权限（新增/编辑/删除）三级菜单。

---

## v1.1.0 — 初始版本

### 核心功能

- 数据库表结构自动解析，支持 `{应用名}_{模块名}` 表名规范
- 后端生成：API / Controller / Logic / Service / Model / Consts
- 前端生成：TypeScript 类型 / API 函数 / 列表页 / 表单弹窗
- 智能组件映射：按字段名、枚举数量、数据库类型自动选择前端组件
- 树形结构检测（`parent_id`）、密码加密、外键关联、Snowflake ID 处理
- 枚举常量自动生成（Go 常量 + 前端 options）
- `--dry-run` 预览模式、`--force` 强制覆盖、`--only backend/frontend` 部分生成
