package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// ReplacePlaceholders 将路径中的 {app} 和 {module} 占位符替换为实际名称
func ReplacePlaceholders(path, app, module string) string {
	result := strings.ReplaceAll(path, "{app}", app)
	result = strings.ReplaceAll(result, "{module}", module)
	return result
}

// TemplateMapping 模板文件名 → 输出相对路径
type TemplateMapping struct {
	TplFile    string
	OutputPath string
}

// GenerateFiles 通用文件生成逻辑
func GenerateFiles(mappings []TemplateMapping, tplDir, outDir, appName, moduleName string, force bool, data interface{}) ([]string, error) {
	var generated []string

	for _, m := range mappings {
		tplPath := filepath.Join(tplDir, m.TplFile)
		tpl, err := template.ParseFiles(tplPath)
		if err != nil {
			return generated, fmt.Errorf("解析模板 %s 失败: %w", m.TplFile, err)
		}

		relPath := ReplacePlaceholders(m.OutputPath, appName, moduleName)
		outPath := filepath.Join(outDir, relPath)

		if !force {
			if _, err := os.Stat(outPath); err == nil {
				fmt.Printf("  [跳过] %s（已存在，使用 --force 覆盖）\n", outPath)
				continue
			}
		}

		dir := filepath.Dir(outPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return generated, fmt.Errorf("创建目录 %s 失败: %w", dir, err)
		}

		file, err := os.Create(outPath)
		if err != nil {
			return generated, fmt.Errorf("创建文件 %s 失败: %w", outPath, err)
		}

		if err := tpl.Execute(file, data); err != nil {
			file.Close()
			return generated, fmt.Errorf("渲染模板 %s 失败: %w", m.TplFile, err)
		}
		file.Close()

		generated = append(generated, outPath)
		fmt.Printf("  [生成] %s\n", outPath)
	}

	return generated, nil
}
