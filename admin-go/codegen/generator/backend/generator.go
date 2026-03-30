package backend

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gbaseadmin/codegen/parser"
)

// templateMapping 模板文件名 → 输出相对路径模板
// {module} 会被替换为 meta.ModuleName
type templateMapping struct {
	TplFile    string // 模板文件名
	OutputPath string // 输出相对路径，含 {module} 占位符
}

var mappings = []templateMapping{
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
// 返回生成的文件路径列表
func (g *Generator) Generate(meta *parser.TableMeta) ([]string, error) {
	var generated []string

	for _, m := range mappings {
		// 解析模板文件
		tplPath := filepath.Join(g.config.TemplateDir, m.TplFile)
		tpl, err := template.ParseFiles(tplPath)
		if err != nil {
			return generated, fmt.Errorf("解析模板 %s 失败: %v", m.TplFile, err)
		}

		// 构建输出路径：将 {app} 和 {module} 替换为实际名称
		relPath := replacePlaceholders(m.OutputPath, meta.AppName, meta.ModuleName)
		outPath := filepath.Join(g.config.OutputDir, relPath)

		// 文件已存在且不强制覆盖，跳过
		if !g.config.Force {
			if _, err := os.Stat(outPath); err == nil {
				fmt.Printf("  跳过（已存在）: %s\n", outPath)
				continue
			}
		}

		// 创建目标目录
		dir := filepath.Dir(outPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return generated, fmt.Errorf("创建目录 %s 失败: %v", dir, err)
		}

		// 渲染模板并写入文件
		file, err := os.Create(outPath)
		if err != nil {
			return generated, fmt.Errorf("创建文件 %s 失败: %v", outPath, err)
		}

		if err := tpl.Execute(file, meta); err != nil {
			file.Close()
			return generated, fmt.Errorf("渲染模板 %s 失败: %v", m.TplFile, err)
		}
		file.Close()

		generated = append(generated, outPath)
	}

	return generated, nil
}

// replacePlaceholders 将路径中的 {app} 和 {module} 占位符替换为实际名称
func replacePlaceholders(path, app, module string) string {
	result := strings.ReplaceAll(path, "{app}", app)
	result = strings.ReplaceAll(result, "{module}", module)
	return result
}
