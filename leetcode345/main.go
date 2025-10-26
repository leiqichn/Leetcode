package leetcode345

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {

	// 记录index,
	res := []int{}
	bytes := []byte(s)

	for i, c := range s {
		if strings.Contains("aeiouAEIOU", string(c)) {
			res = append(res, i)
		}

	}
	fmt.Println(res)
	fmt.Println(bytes)
	// 然后反转index list 对应的值
	reverse(bytes, res)
	return string(bytes)
}

func reverse(s []byte, l []int) {
	for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
		s[l[i]], s[l[j]] = s[l[j]], s[l[i]]
	}
	fmt.Println(s)
}
