/*
 * Copyright (c) 2023 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2023/12/12 下午11:21
 */

package leetcode

func commonFactors(a int, b int) int {
	minNum := min(a, b)
	cnt := 0
	for i := 1; i <= minNum; i++ {
		if (a%i == 0) && (b%i == 0) {
			cnt++
		}
	}
	return cnt
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
