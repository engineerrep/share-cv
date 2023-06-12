package filter

import (
	"github.com/zeromicro/go-zero/core/stringx"
)

// CheckSensitiveWord 检查敏感字符
func CheckSensitiveWord(val string) (string, []string, bool) {
	filter := stringx.NewTrie([]string{
		"AV",
	}, stringx.WithMask('*')) // 默认替换为*
	safe, keywords, found := filter.Filter(val)
	return safe, keywords, found
}
