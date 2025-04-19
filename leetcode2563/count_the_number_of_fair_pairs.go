package leetcode2563

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	var count int64
	n := len(nums)
	for i := 0; i < n; i++ {
		lowVal := lower - nums[i]
		highVal := upper - nums[i]
		start := i + 1
		end := n - 1
		if start > end {
			continue
		}
		subStart := start
		subLength := end - subStart + 1

		// 查找第一个大于等于lowVal的位置
		leftInSub := sort.Search(subLength, func(k int) bool {
			return nums[subStart+k] >= lowVal
		})
		left := subStart + leftInSub

		// 查找第一个大于highVal的位置，然后减一得到最后一个<=highVal的位置
		rightInSub := sort.Search(subLength, func(k int) bool {
			return nums[subStart+k] > highVal
		}) - 1
		right := subStart + rightInSub

		if left <= right && right >= subStart {
			count += int64(right - left + 1)
		}
	}
	return count
}
