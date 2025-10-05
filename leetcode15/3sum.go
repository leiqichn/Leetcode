package leetcode15

import "sort"

func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums) // 先排序可使用双指针
	n := len(nums)

	for i := 0; i < n-2; i++ {
		x := nums[i]

		if i > 0 && x == nums[i-1] { // 重复下一个
			continue
		}
		j := i + 1
		k := n - 1

		for j < k {

			sum := x + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else { // 三数之和等于0
				ans = append(ans, []int{x, nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {

				}
				for k--; j < k && nums[k] == nums[k-1]; k-- { // 条件需要判断为k+1

				}
			}

		}

	}
	return
}
