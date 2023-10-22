package _0230827_56_merge_intervals

import "sort"

func merge(intervals [][]int) [][]int {
	// 遍历每个节点， 判断节点 的后一个节点 的idx 0 是否在上个节点的idx 0-1 之间，如果在，则合并，否则到到下个节点判断
	// 存在有多个合并的情况 [1,2] [2,3] [3,4]  => [1,3] 需要使用front 中间变量
	// 终止条件 下个节点为空，则直接append 当前节点到res；若当前节点为空，则直接结束。
	// 先按照idx 0 进行大小排序

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})

	front := []int{}
	front = intervals[0]
	res := [][]int{}
	lens := len(intervals)
	if lens == 1 {
		return intervals
	}
	for i := 0; i+1 < lens; i++ {

		if intervals[i+1][0] <= front[1] { // 要和前边的对比 而不是index i
			if intervals[i+1][1] > front[1] {
				front[1] = intervals[i+1][1]
			}
		} else {
			res = append(res, front)
			front = intervals[i+1]
		}
	}
	res = append(res, front)
	return res
}
