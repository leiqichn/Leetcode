package leetcode2563

import (
	"slices"
	"sort"
)

func countFairPairs(nums []int, lower, upper int) (ans int64) {
	slices.Sort(nums)
	for j, x := range nums {
		// 注意要在 [0, j-1] 中二分，因为题目要求两个下标 i < j
		r := sort.SearchInts(nums[:j], upper-x+1)
		l := sort.SearchInts(nums[:j], lower-x)
		ans += int64(r - l)
	}
	return
}
