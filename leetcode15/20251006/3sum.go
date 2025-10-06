package _0251006

import (
	"slices"
	"sort"
)

func threeSum(nums []int) [][]int {
	// 先排序
	// 输出的顺序和三元组的顺序并不重要
	// i<j<k

	sort.Ints(nums)
	slices.Sort(nums) // 支持其它整型的排序

	n := len(nums)
	res := [][]int{}
	for i, x := range nums {

		if i > 0 && nums[i-1] == x {
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
			} else {
				res = append(res, []int{x, nums[j], nums[k]})
				// j 去重

				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}

				// k 去重
				for k--; j < k && nums[k] == nums[k+1]; k-- {
				}
			}

		}

	}
	return res
}
