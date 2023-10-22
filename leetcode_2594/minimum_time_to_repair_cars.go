package leetcode_2594

import "math"

func repairCars(ranks []int, cars int) int64 {
	l, r := 1, ranks[0]*cars*cars
	var check = func(m int) bool {
		cnt := 0
		for _, x := range ranks {
			cnt += int(math.Sqrt(float64(m / x)))
		}
		return cnt >= cars
	}

	for l < r {
		m := (l + r) >> 1
		if check(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return int64(l)
}
