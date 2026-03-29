package frontend

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"gbaseadmin/codegen/parser"
)

// templateMapping 模板文件名 → 输出相对路径模板
// {module} 会被替换为 meta.ModuleName
type templateMapping struct {
	TplFile    string // 模板文件名，如 types.tpl
	OutputPath string // 输出相对路径，含 {module} 占位符
}

var mappings = []templateMapping{
	{"types.tpl", "api/system/{module}/types.ts"},
	{"api.tpl", "api/system/{module}/index.ts"},
	{"list.tpl", "views/system/{module}/index.vue"},
	{"form.tpl", "views/system/{module}/modules/form.vue"},
}

// Config 前端生成器配置
type Config struct {
	TemplateDir string // 模板目录路径
	OutputDir   string // 输出根目录，如 ./vue-vben-admin/apps/web-antd/src/
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
// 返回生成的文件路径列表
func (g *Generator) Generate(meta *parser.TableMeta) ([]string, error) {
	var generated []string

	for _, m := range mappings {
		// 解析模板
		tplPath := filepath.Join(g.config.TemplateDir, m.TplFile)
		tpl, err := template.ParseFiles(tplPath)
		if err != nil {
			return generated, fmt.Errorf("解析模板 %s 失败: %w", m.TplFile, err)
		}

		// 构建输出路径，替换 {module} 占位符
		relPath := replaceModule(m.OutputPath, meta.ModuleName)
		outPath := filepath.Join(g.config.OutputDir, relPath)

		// 文件已存在且非强制覆盖，跳过
		if !g.config.Force {
			if _, err := os.Stat(outPath); err == nil {
				fmt.Printf("  [跳过] %s（已存在，使用 --force 覆盖）\n", outPath)
				continue
			}
		}

		// 创建目标目录
		dir := filepath.Dir(outPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return generated, fmt.Errorf("创建目录 %s 失败: %w", dir, err)
		}

		// 创建输出文件并渲染模板
		file, err := os.Create(outPath)
		if err != nil {
			return generated, fmt.Errorf("创建文件 %s 失败: %w", outPath, err)
		}

		if err := tpl.Execute(file, meta); err != nil {
			file.Close()
			return generated, fmt.Errorf("渲染模板 %s 失败: %w", m.TplFile, err)
		}
		file.Close()

		generated = append(generated, outPath)
		fmt.Printf("  [生成] %s\n", outPath)
	}

	return generated, nil
}

// replaceModule 将路径中的 {module} 替换为实际模块名
func replaceModule(path, module string) string {
	result := ""
	for i := 0; i < len(path); {
		if i+8 <= len(path) && path[i:i+8] == "{module}" {
			result += module
			i += 8
		} else {
			result += string(path[i])
			i++
		}
	}
	return result
}
