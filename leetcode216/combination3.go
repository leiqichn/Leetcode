package leetcode216

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	path := []int{}

	var dfs func(start, sum int)
	dfs = func(start, sum int) {

		// 找到一个合法组合
		if len(path) == k && sum == n {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		// 剪枝：剩余数字不足以填满 path 或无法达到 sum
		for i := start; i >= 1; i-- {
			if i < k-len(path) { // 确保剩余数字足够填满 path
				break
			}
			if sum+i > n { // 提前终止，避免无效递归  如果 sum + i > n，说明当前 i 太大，不能选它。但 更小的 i 仍然可能满足 sum + i <= n，所以不能直接 break（否则会漏掉可能的解）。
				continue
			}
			path = append(path, i)
			dfs(i-1, sum+i)
			path = path[:len(path)-1]
		}
	}

	dfs(9, 0) // 数字范围是 1~9
	return res
}
