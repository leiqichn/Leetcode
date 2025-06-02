package leetcode

import "sort"

// 申请全局变量

// 注意点 数字可以重复（start 不用+1），没有明确k层数的限制（深度）
var (
	path []int
	res  [][]int
)

func combinationSum(candidates []int, target int) [][]int {
	path, res = make([]int, 0), make([][]int, 0)
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	backTracking(candidates, target, 0, 0)
	return res

}

func backTracking(candidates []int, target int, sum int, startIndex int) {
	// 如果sum > target 直接return 不会记录结果
	if sum > target {
		return
	}
	// sum == target 则记录结果到res
	if sum == target {
		var tmp []int
		tmp = make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}
	// for 循环 遍历
	for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {
		// 单次循环的操作
		sum += candidates[i]
		path = append(path, candidates[i])
		backTracking(candidates, target, sum, i)
		path = path[:len(path)-1]
		sum -= candidates[i]
	}
}
