package leetcode

func removeElement(nums []int, val int) int {
	// 暴力解法； 遍历 遇到该值则将后边的数都向前移动一位，否则 count++ reurn count
	var count int
	for _, i := range nums {
		if i == val {
			for j := i; j < len(nums)+1; j++ {
				nums[j] = nums[j+1]
			}
		}
		i--
		count++
	}
	return count
}

func removeElement2(nums []int, val int) int {
	// 暴力解法； 遍历 遇到该值则将后边的数都向前移动一位，否则 count++ reurn count
	var count int = len(nums)
	//numsBak := nums
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			for j := i; j < len(nums)-1; j++ {
				nums[j] = nums[j+1]
			}
			i--
			count--
		}

	}
	return count
}

// 快慢指针
func removeElement3(nums []int, val int) int {
	lens := len(nums)
	slow := 0

	for fast := 0; fast < lens; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast] // 将快指针的值 赋值给slow
			slow++
		}
	}
	return slow

}
