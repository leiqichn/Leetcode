package leetcode283

func moveZeroes(nums []int) {
	left, n := 0, len(nums)
	for right := 0; right < n; right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
