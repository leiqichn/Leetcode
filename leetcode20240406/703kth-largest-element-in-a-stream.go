package leetcode20240406

import "sort"

type KthLargest struct {
	// 排序后的nums
	index      int
	sortedNums []int
}

func Constructor2(k int, nums []int) KthLargest {
	return KthLargest{k, nums}
}

func (this *KthLargest) Add(val int) int {
	this.sortedNums = append(this.sortedNums, val)
	sort.Ints(this.sortedNums)
	revert(this.sortedNums)
	return this.sortedNums[this.index-1]
}

func revert(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
