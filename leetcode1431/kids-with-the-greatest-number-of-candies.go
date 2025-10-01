package leetcode1431

func kidsWithCandies(candies []int, extraCandies int) []bool {
	// n  candies
	// 额外的给i 后， 如果她是最大(可以相等），则返回true

	// 思路： 先求得当前数组中最大的值， 判断当前值+ extraCandies 是否大于等于即可

	max := 0

	res := make([]bool, len(candies))

	for _, v := range candies {
		if v > max {
			max = v
		}
	}

	for i, v := range candies {

		if v+extraCandies >= max {
			res[i] = true
		} else {
			res[i] = false //  需要注意的是返回值为bool
		}
	}
	return res
}
