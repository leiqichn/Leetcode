package leetcode

// 1 岛屿，0 海水， 见到岛屿则使用dfs 淹了这个所有1相连的岛屿，res++
func numIslands(grid [][]byte) int {
	res := 0
	row := len(grid)
	col := len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		m := len(grid)
		n := len(grid[0])
		// 判断边界
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}

		if grid[i][j] == '0' {
			return
		}

		grid[i][j] = '0'

		dfs(i+1, j)
		dfs(i, j+1)
		dfs(i-1, j)
		dfs(i, j-1)

	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' { // 这里需要判断
				dfs(i, j)
				res++
			}
		}
	}

	return res
}
