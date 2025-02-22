package leetcode454

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	maps := make(map[int]int, len(nums1))
	// 不去重， 暴力做法：需要遍历四个数值。分治，先算两个数组中的和，然后在求对应两个列表的数值之和。
	for _, v := range nums1 {
		for _, v2 := range nums2 {
			maps[v+v2] += 1
		}
	}
	count := 0
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			if val, ok := maps[-v3-v4]; ok {
				count += val
			}
		}
	}
	return count
}
