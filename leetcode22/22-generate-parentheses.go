package leetcode22

func generateParenthesis(n int) (ans []string) {
	// 选或者不选的思路
	m := n * 2
	path := make([]byte, m)
	var dfs func(int, int)
	dfs = func(i, open int) { // open设定为左括号的数量
		if i == m {
			ans = append(ans, string(path))
			return
		}
		if open < n { // 可以填左括号， 左括号小于一半，则可以添加
			path[i] = '('
			dfs(i+1, open+1) //到下一个位置，同时左括号添加1
		}
		if i-open < open { // 可以填右括号
			path[i] = ')'
			dfs(i+1, open)
		}
	}
	dfs(0, 0)
	return
}
