package leetcode28

import "strings"

func strStr(haystack string, needle string) int {
	list := strings.Split(haystack, needle)
	if len(list) > 1 {
		return len(list[0])
	}
	return 0
}
