package leetcode

import "math"

// 方法1 使用了类似动态规划的思想
// 定义 nums[i] 当前元素，nums[i-1] 前序列之和
func maxSubArray(nums []int) int {
	max := nums[0] //初始化最大值为前边一个元素
	// 是判断当前连续子序列能否对后面的数字产生增益的条件，在算法中起到非常重要的作用。
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		// 超过最大值，则更新
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// 方法2
func maxSubArray(nums []int) int {
	count := 0
	res := math.MinInt32

	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > res {
			res = count
		}

		if count < 0 {
			count = 0
		}
	}
	return res
}
