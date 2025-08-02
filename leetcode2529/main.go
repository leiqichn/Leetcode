package leetcode2529

import "sort"

func maximumCount(nums []int) int {
	// 返回反照非低价顺序排列的数组 nums ， 返回正整数数目和负整数数目中的最大值
	// pos neg pos neg max
	// 获得 >0 =》 >= target + 1 ;  <0   >=(target -1) + 1

	fuIndex := sort.SearchInts(nums, 1)
	if fuIndex == len(nums) {
		return len(nums)
	}

	zhenIndex := sort.SearchInts(nums, 0)

	if fuIndex+1 > len(nums)-zhenIndex {
		return fuIndex + 1
	}
	return len(nums) - zhenIndex
}
