package _0250613

func numSubarrayProductLessThanK(nums []int, k int) int {
	slow := 0
	res := 0
	cheng := 1

	if k <= 1 {
		return 0
	}

	for fast, val := range nums {
		cheng *= val

		for cheng >= k { // 不符合条件，才缩小
			cheng /= nums[slow]
			slow++
		}
		res += fast - slow + 1 // 【left,right], [left +1， right], ...[right ,right]

	}

	return res
}
