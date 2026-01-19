package leetcode724

func pivotIndex(nums []int) int {
	// 两边之和相等， 前缀和
	sums := make([]int, len(nums))
	for i, _ := range nums {
		for j := 0; j < i; j++ {
			sums[i] += nums[j]
		}
	}

	mid := sums[len(nums)-1] / 2 // 中间的值

	// 然后判断前缀和找到index 使得前后相等的idx
	for i, v := range sums {
		if v == mid {
			return i - 1
		}
	}

	return -1
}
