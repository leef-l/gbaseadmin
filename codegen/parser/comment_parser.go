package parser

import "strings"

// ParseComment 解析字段备注
// 输入：状态:0=关闭,1=开启
// 输出：label="状态", enums=[{0,关闭},{1,开启}]
// 输入：部门名称
// 输出：label="部门名称", enums=[]
func ParseComment(comment string) (label string, enums []EnumValue) {
	comment = strings.TrimSpace(comment)
	if comment == "" {
		return "", nil
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
		return comment, nil
	}

	label = strings.TrimSpace(comment[:sepIdx])
	enumPart := strings.TrimSpace(comment[sepIdx+1:])
	// 冒号后面的字节偏移需要考虑中文冒号占3字节
	if len(label) == 0 {
		label = comment
		return label, nil
	}

	// 解析枚举部分：0=关闭,1=开启
	if enumPart == "" {
		return label, nil
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

	return label, enums
}
