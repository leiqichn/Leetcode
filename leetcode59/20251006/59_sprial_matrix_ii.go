package _0251006

func generateMatrix(n int) [][]int {
	top, bottom := 0, n-1
	left, right := 0, n-1

	target := n * n

	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	count := 1
	for count <= target {
		for i := left; i <= right; i++ {
			matrix[top][i] = count
			count++
		}
		top++

		for i := top; i <= bottom; i++ {
			matrix[i][right] = count
			count++
		}
		right--

		for i := right; i >= left; i-- {
			matrix[bottom][i] = count
			count++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			matrix[i][left] = count
			count++
		}
		left++
	}

	return matrix
}
