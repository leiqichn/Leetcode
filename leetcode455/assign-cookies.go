/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description: 贪心算法
 * Date: 2024/11/27 下午11:53
 */

package leetcode455

import "sort"

func findContentChildren(g []int, s []int) int {
	// 使用大的饼干gIndex 来满足孩子 sIndex
	// 都排序
	sort.Ints(g)
	sort.Ints(s)
	// 遍历饼干s  遍历孩子胃口g
	res := 0
	sIndex := 0
	for i := 0; i < len(s); i++ {
		if sIndex < len(g) && s[i] >= g[sIndex] {
			res++
			sIndex++
		}
	}
	return res
}
