/*
 * Copyright (c) 2024 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2024/1/24 上午12:21
 */

package _00

type point struct {
	x int
	y int
}

var dirct = [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}
var visited = map[point]int{}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0
	r := len(grid)
	c := len(grid[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

//r c

func dfs(grid [][]byte, i, j int) {

	r := len(grid)
	c := len(grid[0])
	// 边界条件
	if i < 0 || j < 0 || i >= r || j >= c {
		return
	}
	// 如果是海水
	if grid[i][j] == '0' {
		// 已经是海水了
		return
	}

	// 记录
	// 将陆地变成海水
	grid[i][j] = '0'
	visited[point{i, j}] = 1
	// 调用自身 淹没上下左右的陆地
	for _, item := range dirct {

		myRow := i + item[0]
		myCol := j + item[1]
		if _, ok := visited[point{myRow, myCol}]; !ok {
			dfs(grid, myRow, myCol)
		}
	}
}
