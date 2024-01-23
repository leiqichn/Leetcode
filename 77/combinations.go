/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/1/23 下午10:57
 */

package _7

func combine(n int, k int) [][]int {
	res := [][]int{}
	path := []int{}

	// 结束条件
	var backtracking func(n, k, startIdx int)
	backtracking = func(n, k, startIdx int) {
		// 提前返回需要return
		if len(path) == k {
			pathTmp := make([]int, len(path))
			copy(pathTmp, path)
			res = append(res, pathTmp)
			return
		}

		for i := startIdx; i <= n; i++ {
			path = append(path, i)
			backtracking(n, k, i+1)
			path = path[:(len(path) - 1)]
		}
	}
	backtracking(n, k, 1)
	return res
}
