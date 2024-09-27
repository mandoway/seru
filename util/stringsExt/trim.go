package stringsExt

import "strings"

func TrimAfterLast(s, substr string) string {
	indexOfLast := strings.LastIndex(s, substr)
	if indexOfLast == -1 {
		return s
	}
	return s[:indexOfLast]
}
