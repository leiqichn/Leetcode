package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dpNumPath := make([][]int, m)
	// new 二维数组
	for i := 0; i < m; i++ {
		dpRow := make([]int, n)
		dpNumPath[i] = dpRow
	}
	// 初始化
	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dpNumPath[i][0] = 1
	}
	for j := 0; j < n && obstacleGrid[0][j] == 0; j++ {
		dpNumPath[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dpNumPath[i][j] = dpNumPath[i-1][j] + dpNumPath[i][j-1]
			}
		}
	}
	return dpNumPath[m-1][n-1]
}
