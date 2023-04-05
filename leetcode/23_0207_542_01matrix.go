package leetcode

type point struct {
	x int
	y int
}

func updateMatrix(mat [][]int) [][]int {
	startX := 0
	startY := 0
	m, n := len(mat), len(mat[0])
	queue := []point{}
	queue = append(queue, point{startX, startY})
	// step 是再循环里边定义
	dirs := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	resMatrix := make([][]int, m)

	for i := 0; i < m; i++ {
		resMatrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			resMatrix[i][j] = 0
		}
	}
	for len(queue) != 0 {
		for _, q := range queue {
			step := 0
			for _, dir := range dirs {
				// 遍历方向
				newPointX := q.x + dir.x
				newPointY := q.y + dir.y
				// 确定符合边界条件，并且没有 visited 过
				if newPointX >= 0 && newPointX < n && newPointY >= 0 && newPointY <= m && resMatrix[newPointX][newPointY] == 0 && mat[newPointX][newPointX] == 0 {
					resMatrix[newPointX][newPointY] = step
					queue = append(queue, point{newPointX, newPointY})
				}
			}
			step++
		}
	}
	return resMatrix
}
