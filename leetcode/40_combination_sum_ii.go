package leetcode

import "sort"

var (
	res  [][]int
	path []int
	used []bool
)

func combinationSum2(candidates []int, target int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, len(candidates))
	used = make([]bool, len(candidates))
	sort.Ints(candidates) // 排序，为剪枝做准备
	dfs(candidates, 0, target)
	return res
}

func dfs(candidates []int, start int, target int) {
	if target == 0 { // target 不断减小，如果为0说明达到了目标值
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target { // 剪枝，提前返回
			break
		}
		// used[i - 1] == true，说明同一树枝candidates[i - 1]使用过，递归是走函数的，所以是树枝去重
		// used[i - 1] == false，说明同一树层candidates[i - 1]使用过 前一个是进行了回溯的，所以同一层的前一个是恢复过的
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
			continue
		}
		path = append(path, candidates[i])
		used[i] = true
		dfs(candidates, i+1, target-candidates[i])
		used[i] = false
		path = path[:len(path)-1]
	}
}
