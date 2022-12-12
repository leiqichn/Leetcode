package leetcode

func generateMatrix(n int) [][]int {
	top, bottom := 0, n-1
	left, right := 0, n-1 // 注意边界
	num := 1
	tar := n * n
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for n / 2 {
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}

func generateMatrix(n int) [][]int {
	num := 1
	loop := n / 2
	offset := 1
	matrix := make([][]int, n)
	startX, startY := 0, 0
	i, j := 0, 0
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for loop > 0 {
		for j = startY; j < n-offset; j++ { // j 为列的idx
			matrix[startX][j] = num
			num++
		}
		for i = startX; i < n-offset; i++ {
			matrix[i][j] = num
			num++
		}
		for ; j > startY; j-- {
			matrix[i][j] = num
			num++
		}
		for ; i > startX; i-- {
			matrix[i][j] = num
			num++
		}
		loop--
		offset++
		startX++
		startY++

	}

	if n%2 == 1 {
		matrix[startX][startY] = num
	}
	return matrix
}
