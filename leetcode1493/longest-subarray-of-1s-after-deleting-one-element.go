package leetcode1493

func longestSubarray(nums []int) int {
	// 01 删除一个元素
	// 全1， 的最长数组
	// 如果不存在则返回0

	// 计算0的长度， 当0 的长度大于1， 收缩

	res := 0

	left := 0
	count := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			count++
		}
		for count > 1 {
			if nums[left] == 0 {
				count--
			}
			left++
		}

		res = max(res, right-left) // 我们计算的是删除一个元素后的最长连续1的子数组长度。
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
