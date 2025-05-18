package leetcode695

func maxAreaOfIsland(grid [][]int) int {
	dic := [][]int{
		{1, 0},
		{-1, 0},
		{0, -1},
		{0, 1},
	}

	var dfs func([][]int, int, int) int // 修改返回类型为int

	dfs = func(grid [][]int, x, y int) int { // 返回当前岛屿面积
		row := len(grid)
		col := len(grid[0])

		// 越界检查
		if x < 0 || x >= row || y < 0 || y >= col {
			return 0
		}

		// 遇到海水或已访问
		if grid[x][y] == 0 {
			return 0
		}

		// 淹没当前点并初始化面积为1
		grid[x][y] = 0
		area := 1

		// 向四个方向扩散并累加面积
		for _, d := range dic {
			area += dfs(grid, x+d[0], y+d[1])
		}
		return area
	}

	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				currentArea := dfs(grid, i, j)
				if currentArea > maxArea {
					maxArea = currentArea
				}
			}
		}
	}
	return maxArea
}
