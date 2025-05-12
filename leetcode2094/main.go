package leetcode2094

import "sort"

func findEvenNumbers(digits []int) []int {
	//暴力
	res := []int{}
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			for k := j + 1; k < len(digits); k++ {
				if digits[i] != 0 && digits[k]%2 == 0 {
					tmpNums := digits[i]*100 + digits[k]*10 + digits[k]
					res = append(res, tmpNums)
				}
			}
		}
	}
	sort.Ints(res)
	return res
}
