package leetcode167

func twoSum(numbers []int, target int) []int {
	// 有序数组 -> 二分， 想到双指针
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{}
}
