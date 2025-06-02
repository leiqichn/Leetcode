package leetcode

import "math"

func minSubArrayLen(target int, nums []int) int {
	// n 个正整数 数组 和target
	// 找出 子数组中大于等于target 的
	res := math.MaxInt64
	slow := 0
	fast := 0
	sum := 0
	for ; fast < len(nums); fast++ {
		sum += nums[fast]
		if sum >= target {
			subMinLens := fast - slow + 1
			if subMinLens < res {
				res = subMinLens
			}
			sum -= nums[slow]
			slow++
		}

	}
	if res == math.MaxInt64 {
		return 0
	} else {
		return res
	}
}
