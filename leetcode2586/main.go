package leetcode2586

import "strings"

func vowelStrings(words []string, left int, right int) int {
	// strings.Contains
	count := 0
	for i := left; i <= right; i++ {
		tmpStr := words[i]

		if strings.Contains("aeiou", string(tmpStr[0])) && strings.Contains("aeiou", string(tmpStr[len(tmpStr)-1])) {
			count++
		}
	}
	return count
}
