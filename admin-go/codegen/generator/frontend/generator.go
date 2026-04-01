package frontend

import (
	"gbaseadmin/codegen/generator/util"
	"gbaseadmin/codegen/parser"
)

var mappings = []util.TemplateMapping{
	{"types.tpl", "api/{app}/{module}/types.ts"},
	{"api.tpl", "api/{app}/{module}/index.ts"},
	{"list.tpl", "views/{app}/{module}/index.vue"},
	{"form.tpl", "views/{app}/{module}/modules/form.vue"},
	{"detail-drawer.tpl", "views/{app}/{module}/modules/detail-drawer.vue"},
}

// Config 前端生成器配置
type Config struct {
	TemplateDir string // 模板目录路径
	OutputDir   string // 输出根目录
	Force       bool   // 是否强制覆盖
}

// Generator 前端代码生成器
type Generator struct {
	config Config
}

// New 创建前端代码生成器实例
func New(cfg Config) *Generator {
	return &Generator{config: cfg}
}

// Generate 为一张表生成所有前端代码
func (g *Generator) Generate(meta *parser.TableMeta) ([]string, error) {
	return util.GenerateFiles(mappings, g.config.TemplateDir, g.config.OutputDir, meta.AppName, meta.ModuleName, g.config.Force, meta)
}
