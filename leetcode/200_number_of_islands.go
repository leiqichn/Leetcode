package leetcode

// 1 岛屿，0 海水， 见到岛屿则使用dfs 淹了这个所有1相连的岛屿，res++
func numIslands(grid [][]byte) int {
	res := 0
	r := len(grid)
	c := len(grid[0])
	visited := make([][]byte, r)
	for i := 0; i < r; i++ {
		visited[i] = make([]byte, c)
	}

	// 遍历二维slice
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			// 每个位置
			dfs(grid, i, j, visited)
			res++
		}
	}
	return res
}

func dfs(grid [][]byte, i int, j int, visited [][]byte) {
	r := len(grid)
	c := len(grid)

	// 边界判断
	if i < 0 || j < 0 || i >= r || j >= c {
		return
	}

	// 判断值，是海水直接返回
	if grid[i][j] == '0' {
		return
	}

	if visited[i][j] == '1' {
		return
	}

	// 改变状态为海水
	grid[i][j] = '0'

	//单层递归逻辑，四面八方递归
	dfs(grid, i+1, j, visited)
	dfs(grid, i, j+1, visited)
	dfs(grid, i-1, j, visited)
	dfs(grid, i, j-1, visited)
}
