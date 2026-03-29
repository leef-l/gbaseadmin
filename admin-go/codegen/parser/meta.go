package parser

// 前端组件类型常量
const (
	ComponentInput            = "Input"
	ComponentInputNumber      = "InputNumber"
	ComponentTextarea         = "Textarea"
	ComponentSwitch           = "Switch"
	ComponentRadio            = "Radio"
	ComponentSelect           = "Select"
	ComponentTreeSelectSingle = "TreeSelectSingle"
	ComponentTreeSelectMulti  = "TreeSelectMulti"
	ComponentSelectMulti      = "SelectMulti"
	ComponentImageUpload      = "ImageUpload"
	ComponentFileUpload       = "FileUpload"
	ComponentRichText         = "RichText"
	ComponentJsonEditor       = "JsonEditor"
	ComponentPassword         = "Password"
	ComponentInputUrl         = "InputUrl"
	ComponentDateTimePicker   = "DateTimePicker"
	ComponentIconPicker       = "IconPicker"
)

// EnumValue 枚举值
type EnumValue struct {
	Value string
	Label string
}

// FieldMeta 字段元数据
type FieldMeta struct {
	Name         string      // snake_case
	NameCamel    string      // CamelCase（Go 风格，ID/URL 全大写）
	NameDao      string      // CamelCase（GoFrame DAO 风格，Id/Url 首字母大写）
	NameLower    string      // camelCase（首字母小写）
	DBType       string      // varchar/int/bigint/tinyint/text 等
	GoType       string      // string/int/int64/JsonInt64 等
	TSType       string      // string/number/boolean
	Comment      string      // 原始备注
	Label        string      // 前端 Label
	EnumValues   []EnumValue // 枚举值列表
	Component    string      // 前端组件类型
	IsRequired   bool
	IsID         bool // 是否是主键 id
	IsParentID   bool // 是否是 parent_id
	IsForeignKey bool // 是否是 *_id（单选外键）
	IsMultiFK    bool // 是否是 *_ids（多选外键）
	IsTimeField  bool // 是否是 *_at 时间字段
	IsHidden     bool // 表单中隐藏（id/created_at/updated_at/deleted_at/created_by/dept_id）
	IsEnum       bool // 是否有枚举值
	IsPassword   bool // 是否是密码字段
	MaxLength    int
	DefaultValue string
	// 关联字段信息（仅 IsForeignKey 或 IsParentID 时有值）
	RefTable            string // 关联表名，如 article
	RefTableCamel       string // 关联表 CamelCase，如 Article
	RefTableLower       string // 关联表 camelCase，如 article
	RefDisplayField     string // 关联表显示字段 snake_case，如 title
	RefDisplayCamel     string // 关联显示字段 CamelCase，如 Title
	RefDisplayLower     string // 关联显示字段 camelCase，如 title
	RefFieldName        string // 结构体字段名 = RefTableCamel + RefDisplayCamel，如 ArticleTitle
	RefFieldJSON        string // json 名 = RefTableLower + RefDisplayCamel，如 articleTitle
}

// TableMeta 表元数据
type TableMeta struct {
	TableName   string
	ModelName   string // CamelCase，如 Dept
	ModuleName  string // 小写，如 dept
	PackageName string // 包名，如 dept
	Comment     string
	Fields      []FieldMeta
	HasParentID bool // 有 parent_id 字段
	HasStatus   bool // 有 status 字段
	HasSort     bool // 有 sort 字段
	HasPassword bool // 有 password 字段
}
