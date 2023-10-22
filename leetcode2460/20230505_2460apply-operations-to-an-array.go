package leetcode2460

func applyOperations(nums []int) []int {
	var res []int
	res = make([]int, len(nums))
	index := 0
	// 第一次遍历
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			res[index] = nums[i]
			index++
		}
	}
	return res
}
