package Q2

import "sort"

func maxAlternatingSum(nums []int) int64 {
	// nums 重新排列数据 交替得平方和， 偶数前边的系数是正数， 奇数前边的系数负数
	// 可以随便排序， 需要返回最大的得分 贪心， 将大的放在偶数位置0

	// 绝对值排序zz

	tmpNums := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		tmpNums[i] = nums[i] * nums[i]
	}

	sort.Sort(sort.Reverse(sort.IntSlice(tmpNums)))

	res := 0

	n := len(nums)

	if n%2 == 0 {

		for i := 0; i < (n / 2); i++ {
			res += tmpNums[i]
		}
		for i := n / 2; i < (n); i++ {
			res -= tmpNums[i]
		}

	} else {
		for i := 0; i < (n/2 + 1); i++ {
			res += tmpNums[i]
		}
		for i := n/2 + 1; i < (n); i++ {
			res -= tmpNums[i]
		}

	}
	return int64(res)
}
