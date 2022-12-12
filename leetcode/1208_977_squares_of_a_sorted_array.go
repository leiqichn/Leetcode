package leetcode

func sortedSquares(nums []int) []int {
	// 有可能是负数，那么平方 则有可能变成较大的值，所以需要用双指针 从两头比较平方，将较大的放在右边，终止条件 左右指针相遇l=right
	left := 0
	right := len(nums) - 1
	n := len(nums)
	k := len(nums) - 1
	ans := make([]int, n) // 使用新数组存放结果
	for left <= right {
		leftSquares := nums[left] * nums[left]
		rightSquares := nums[right] * nums[right]
		if leftSquares > rightSquares {
			ans[k] = leftSquares
			left++
			k--
		} else {
			ans[k] = rightSquares
			right--
			k--
		}
	}
	return ans
}
