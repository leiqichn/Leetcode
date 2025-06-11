package leetcode77

func combine2(n int, k int) [][]int {
	// 组合
	if k > n {
		return [][]int{}
	}
	res := [][]int{}
	path := []int{}
	var dfs func(i int) // 下一个位置的组合
	dfs = func(i int) {
		// base case
		if len(path) == k {
			tmp := make([]int, k)
			copy(path, tmp)
			res = append(res, path)
		}

		//path 还需要 n - len(path)
		for i := n; i >= 1; i-- {
			if i < n-len(path) {
				continue
			}
			path = append(path, i)
			dfs(i - 1)
			path = path[:len(path)-1]
		}
	}

	dfs(n)

	return res
}
