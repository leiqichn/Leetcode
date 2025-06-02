package leetcode

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})
	for i := 0; i < len(nums); i++ { // 第二步
		if nums[i] < 0 && k > 0 {
			nums[i] *= -1
			k--
		}
	}
	if k%2 == 1 {
		nums[len(nums)-1] *= -1
	}
	return sumOfNums(nums)
}

func sumOfNums(nums []int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}
