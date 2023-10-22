package leetcode

type point struct {
	x int
	y int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	res := 0
	dirs := [][]int{{1, 1}, {1, -1}, {1, 0}, {0, -1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}}
	var help func(i, j int)
	//curPoint := &point{0, 0}

	help = func(i, j int) {
		tmp := 0
		n := len(grid)
		visited := make(map[point]bool)
		// 判断是否yue
		if i < 0 || i > n-1 || j < 0 || j > n-1 {
			return
		}
		curPoint := point{i, j}
		if val, ok := visited[curPoint]; ok && val {
			return
		}

		for _, item := range dirs {
			x := item[0]
			y := item[1]
			i += x
			j += y

			if curPoint.x == n-1 && curPoint.y == n-1 {
				if tmp > res {
					res = tmp
					return
				}
			}
			if i < 0 || i > n-1 || j < 0 || j > n-1 {
				return
			}
			if grid[i][j] == 0 {
				tmp++
				help(i, j)
			}
		}
	}
	help(0, 0)
	return res
}
