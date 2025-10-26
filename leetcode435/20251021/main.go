package _0251021

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := 1

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[i-1][1] {

			res++

		} else {
			intervals[i][1] = min(intervals[i-1][1], intervals[i][1])
		}
	}

	return len(intervals) - res

}
