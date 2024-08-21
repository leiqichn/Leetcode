/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/8/21 下午10:47
 */

package leetcode1

func twoSum(nums []int, target int) []int {
	maps := make(map[int]int, len(nums))

	for idx, num := range nums {
		maps[num] = idx
	}

	for idx, num := range nums {
		if idx2, ok := maps[target-num]; ok && idx != idx2 { // 00 : 04 : 10 使用哈希表，需要注意的是，有可能使用了同一个idx
			return []int{idx, idx2}
		}
	}
	return []int{}
}
