package leetcode258

import "strconv"

func addDigits(num int) int {
	// 迭代
	stringNums := strconv.Itoa(num)
	if len(stringNums) == 1 {
		return num
	}
	tmpRes := 0
	for _, runeNum := range stringNums {
		tmpRes += int(runeNum - '0')
	}
	return addDigits(tmpRes)
}
