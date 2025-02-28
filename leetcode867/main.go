package leetcode867

func transpose(matrix [][]int) [][]int {
	// 获取原矩阵的行数和列数
	rows := len(matrix)
	cols := len(matrix[0])

	// 初始化转置矩阵，行数为原矩阵的列数，列数为原矩阵的行数
	newMatrix := make([][]int, cols)
	for i := range newMatrix {
		newMatrix[i] = make([]int, rows)
	}

	// 将原矩阵的元素转移到转置矩阵的相应位置
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			newMatrix[j][i] = matrix[i][j]
		}
	}

	return newMatrix
}
