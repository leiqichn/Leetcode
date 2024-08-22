/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description: 11. 盛最多水的容器 https://leetcode.cn/problems/container-with-most-water/description/
 * Date: 2024/8/22 下午10:52
 */

package _0210821

func maxArea(height []int) int {
	lens := len(height)
	left := 0
	right := lens - 1

	res := 0
	for left < right {
		area := (right - left) * min(height[right], height[left])
		res = max(area, res)

		// 谁短删除谁
		if height[right] < height[left] {
			right -= 1
		} else {
			left += 1
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
