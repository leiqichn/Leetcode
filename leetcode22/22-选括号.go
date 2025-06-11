package leetcode22

import "strings"

func generateParenthesis(n int) []string {
	//n 对括号 只要最后左右总和一样，
	m := n * 2
	res := []string{}

	var dfs func(i, open int)
	path := []string{}
	dfs = func(i, open int) {
		if i == m {
			res = append(res, strings.Join(path, ""))
		}

		// 选左边
		if open < n {
			path = append(path, "(")
			dfs(i+1, open+1)
			path = path[:len(path)-1]
		}

		// 选右边
		if i-open < open { // 右边要小于左边才能加
			path = append(path, ")")
			dfs(i+1, open)
			path = path[:len(path)-1]
		}

	}
	dfs(0, 0) // 入口函数需要添加
	return res
}
