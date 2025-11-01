package leetcode1679

import "sort"

func maxOperations(nums []int, k int) int {
	// 先排序， 判断当前值值 大于 等于 等于的情况
	sort.Ints(nums)
	res := 0
	left := 0
	right := len(nums) - 1
	for left < right { // 双指针， 二分分法变形， left == right时是错误的，题目要求是2个数，但是在双指针中是可以等于的
		sum := nums[left] + nums[right]
		if sum > k {
			right--
		} else if sum < k {
			left++
		} else {
			res++
			left += 1
			right -= 1
		}
	}

	return res
}
