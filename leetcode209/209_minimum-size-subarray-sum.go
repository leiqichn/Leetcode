package leetcode

import "math"

func minSubArrayLen2(target int, nums []int) int {
	// nums target
	// 连续子数组 长度最小
	sum := 0
	left := 0
	minLens := math.MaxInt
	for right, v := range nums {

		for sum >= target {
			minLens = min(minLens, right-left+1)
			sum -= nums[left]
			left--
		}
		sum += v
	}
	return minLens
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
