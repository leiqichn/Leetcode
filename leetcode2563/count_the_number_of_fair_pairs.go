package leetcode2563

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	// 两层遍历
	count := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// 超出时间显示 如何减枝
			if lower > nums[i]+nums[j] {
				continue
			}

			if upper < nums[i]+nums[j] {
				break
			}

			if lower <= nums[i]+nums[j] && upper >= nums[i]+nums[j] {
				count++
			}
		}
	}
	return int64(count)
}
