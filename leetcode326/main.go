package leetcode326

import "math"

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	y := math.Logb(float64(n), 3)
	return y == math.Floor(y) // 检查对数是否为整数
}
