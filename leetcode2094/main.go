package leetcode2094

import "sort"

func findEvenNumbers(digits []int) []int {
	resMap := make(map[int]bool) // 使用map 进行去重
	n := len(digits)
	for i := 0; i < n; i++ {
		if digits[i] == 0 { // 首位数组不能为0
			continue
		}
		for j := 0; j < n; j++ {
			if j == i { // index 不要相同
				continue
			}
			for k := 0; k < n; k++ {
				if k == i || k == j { // index 不要相同
					continue
				}
				if digits[k]%2 == 0 {
					num := digits[i]*100 + digits[j]*10 + digits[k]
					resMap[num] = true // map
				}
			}
		}
	}
	res := []int{}
	for num := range resMap {
		res = append(res, num) // map转ints
	}
	sort.Ints(res)
	return res
}
