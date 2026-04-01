package util

import "strings"

// ReplacePlaceholders 将路径中的 {app} 和 {module} 占位符替换为实际名称
func ReplacePlaceholders(path, app, module string) string {
	result := strings.ReplaceAll(path, "{app}", app)
	result = strings.ReplaceAll(result, "{module}", module)
	return result
}
