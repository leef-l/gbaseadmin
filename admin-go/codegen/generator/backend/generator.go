package backend

import (
	"gbaseadmin/codegen/generator/util"
	"gbaseadmin/codegen/parser"
)

var mappings = []util.TemplateMapping{
	{"api.tpl", "api/{app}/v1/{module}.go"},
	{"controller.tpl", "internal/controller/{module}/{module}.go"},
	{"logic.tpl", "internal/logic/{module}/{module}.go"},
	{"service.tpl", "internal/service/{module}.go"},
	{"model.tpl", "internal/model/{module}.go"},
	{"consts.tpl", "internal/consts/{module}.go"},
}

// Config 后端生成器配置
type Config struct {
	TemplateDir string // 模板目录路径
	OutputDir   string // 输出根目录，如 ./app/system/
	Force       bool   // 是否强制覆盖
}

// Generator 后端代码生成器
type Generator struct {
	config Config
}

// New 创建后端代码生成器实例
func New(cfg Config) *Generator {
	return &Generator{config: cfg}
}

// Generate 为一张表生成所有后端代码
func (g *Generator) Generate(meta *parser.TableMeta) ([]string, error) {
	return util.GenerateFiles(mappings, g.config.TemplateDir, g.config.OutputDir, meta.AppName, meta.ModuleName, g.config.Force, meta)
}
