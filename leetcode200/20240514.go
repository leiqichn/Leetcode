/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/6/3 下午11:41
 */

package leetcode200

func numIslands(grid [][]byte) int {

	dic := [][]int{
		{1, 0},
		{-1, 0},
		{0, -1},
		{0, 1},
	}
	var dfs func([][]byte, int, int)
	// 淹掉陆地
	dfs = func(grid [][]byte, x, y int) {
		row := len(grid)
		col := len(grid[0])

		if x < 0 || x >= row || y < 0 || y >= col {
			return
		}

		if grid[x][y] == '0' {
			return
		}

		grid[x][y] = '0'
		for _, d := range dic {
			xNext := x + d[0]
			yNext := y + d[1]
			dfs(grid, xNext, yNext)

		}
	}

	if len(grid) == 0 {
		return 0
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}
