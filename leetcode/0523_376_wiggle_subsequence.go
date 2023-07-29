package leetcode

func wiggleMaxLength(nums []int) int {
	tmp := []int{}
	var res int
	for i := 0; i < len(nums)-1; i++ {
		curDiff := nums[i+1] - nums[i]
		tmp = append(tmp, curDiff)
	}

	for i := 0; i < len(tmp); i++ {
		if tmp[i]*tmp[i+1] < 0 {
			res++
		} else {
			continue
		}
	}
	return res
}
