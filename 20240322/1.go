package _0240322

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	temp, max := 1, 0

	for i := 0; i < len(nums); i++ {
		if i != 0 {
			if nums[i] == nums[i-1] {
				continue
			} else if nums[i] == nums[i-1]+1 {
				temp++
			} else {
				temp = 1
			}
		}

		if temp > max {
			max = temp
		}

	}
	return max
}
