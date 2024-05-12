/*
 * Copyright (c) 2023 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2023/10/16 下午10:56
 */

package leetcode260

func singleNumber(nums []int) []int {
	res := []int{}
	numsMap := make(map[int]int)
	for _, val := range nums {
		numsMap[val] += 1
	}
	for key, val := range numsMap {
		if val == 1 {
			res = append(res, key)
		}
	}
	return res
}

// hash
func singleNumber2(nums []int) int {
	numsMap := make(map[int]int)
	for _, val := range nums {
		numsMap[val] += 1
	}
	for key, val := range numsMap {
		if val == 1 {
			return key

		}
	}
	return 0
}
