package leetcode

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0] // 身高由大到小排列

	})
	// 再按照K进行插入排序，优先插入K小的
	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i+1]) // 空出一个位置
		people[p[1]] = p
	}
	return people

}
