package leetcode77

func combine(n int, k int) [][]int {
	// 组合
	if k > n {
		return [][]int{}
	}
	res := [][]int{}
	path := []int{}
	var dfs func(start int) // 下一个位置的组合
	dfs = func(start int) {
		// base case
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path) // copy 目标值在前边
			res = append(res, tmp)
		}

		for i := start; i >= 1; i-- { // 倒序方便
			//path 还需要 k - len(path) 提前减枝
			if i < k-len(path) { // 注意是k - len(path)
				break
			}

			path = append(path, i)
			dfs(i - 1)
			path = path[:len(path)-1]
		}
	}
	dfs(n)
	return res
}
