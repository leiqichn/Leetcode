package leetcode1470

func shuffle(nums []int, n int) []int {
	// 重新排列数组
	// 将index n：2n-1 的值， 分为前后两个数组，然后新建一个大数组

	front := nums[:n]
	back := nums[n:]

	res := []int{}

	for index, val := range front {
		if index%2 == 0 {
			res = append(res, val)
		} else {
			res = append(res, back[index])
		}
	}
	return res
}
