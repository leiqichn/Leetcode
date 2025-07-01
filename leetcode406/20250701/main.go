package _0250701

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {

	// sort
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	res := [][]int{}

	for _, info := range people {
		idx := info[1]
		res = append(res[:idx], append([][]int{info}, res[idx:]...)...)
	}

	return res
}
