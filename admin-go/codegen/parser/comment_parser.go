package parser

import "strings"

// extractParentheses 从 label 中提取中文括号（）或英文括号()内的内容
// "排序（升序）" → shortLabel="排序", tooltip="升序"
// "部门名称"     → shortLabel="部门名称", tooltip=""
func extractParentheses(label string) (shortLabel string, tooltip string) {
	// 优先匹配中文括号
	if idx := strings.Index(label, "（"); idx >= 0 {
		if end := strings.Index(label, "）"); end > idx {
			shortLabel = strings.TrimSpace(label[:idx])
			tooltip = strings.TrimSpace(label[idx+len("（") : end])
			return
		}
	}
	// 再匹配英文括号
	if idx := strings.Index(label, "("); idx >= 0 {
		if end := strings.Index(label, ")"); end > idx {
			shortLabel = strings.TrimSpace(label[:idx])
			tooltip = strings.TrimSpace(label[idx+1 : end])
			return
		}
	}
	return label, ""
}

// ParseComment 解析字段备注
// 输入：状态:0=关闭,1=开启
// 输出：label="状态", shortLabel="状态", tooltipText="", enums=[{0,关闭},{1,开启}]
// 输入：排序（升序）
// 输出：label="排序（升序）", shortLabel="排序", tooltipText="升序", enums=[]
// 输入：部门名称
// 输出：label="部门名称", shortLabel="部门名称", tooltipText="", enums=[]
func ParseComment(comment string) (label string, shortLabel string, tooltipText string, enums []EnumValue) {
	comment = strings.TrimSpace(comment)
	if comment == "" {
		return "", "", "", nil
	}

	// 查找冒号分隔符（支持中文冒号和英文冒号）
	sepIdx := -1
	for i, ch := range comment {
		if ch == ':' || ch == '：' {
			sepIdx = i
			break
		}
	}

	// 没有冒号，整个备注就是 label
	if sepIdx < 0 {
		label = comment
		shortLabel, tooltipText = extractParentheses(label)
		return label, shortLabel, tooltipText, nil
	}

	label = strings.TrimSpace(comment[:sepIdx])
	enumPart := strings.TrimSpace(comment[sepIdx+1:])
	if len(label) == 0 {
		label = comment
		shortLabel, tooltipText = extractParentheses(label)
		return label, shortLabel, tooltipText, nil
	}

	shortLabel, tooltipText = extractParentheses(label)

	// 解析枚举部分：0=关闭,1=开启
	if enumPart == "" {
		return label, shortLabel, tooltipText, nil
	}

	pairs := strings.Split(enumPart, ",")
	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}
		eqIdx := strings.Index(pair, "=")
		if eqIdx < 0 {
			continue
		}
		val := strings.TrimSpace(pair[:eqIdx])
		lab := strings.TrimSpace(pair[eqIdx+1:])
		if val != "" && lab != "" {
			enums = append(enums, EnumValue{Value: val, Label: lab})
		}
	}

	return label, shortLabel, tooltipText, enums
}
