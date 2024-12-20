package LCP06

import "math"

func minCount(coins []int) int {
	res := 0
	for i := 0; i < len(coins); i++ {
		res += int(math.Ceil(float64(coins[i] / 2)))
	}
	return res
}
