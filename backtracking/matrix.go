package backtracking

import "math"

// matrix: 必须为n*n的矩阵
func matrix1(matrix [][]int) int {
	maxRow := len(matrix)
	return m(matrix, maxRow, maxRow, 0, 0, matrix[0][0])
}

func m(matrix [][]int, maxRow, maxColumn, row, column, step int) int {
	if row+1 == maxRow && column+1 == maxColumn {
		return step
	}

	t := math.MaxInt64

	if row+1 < maxRow {
		// 向下一行移动
		stepRow := m(matrix, maxRow, maxColumn, row+1, column, step+matrix[row+1][column])
		if t > stepRow {
			t = stepRow
		}
	}
	if column+1 < maxColumn {
		// 向右移动一列
		stepColumn := m(matrix, maxRow, maxColumn, row, column+1, step+matrix[row][column+1])
		if t > stepColumn {
			t = stepColumn
		}
	}

	return t
}

func matrix2(matrix [][]int) int {
	n := len(matrix) // 方行矩阵维度
	status := make([]int, n)
	status[0] = matrix[0][0]

	// 初始化第一行
	for i := range matrix[0] {
		if i+1 < n {
			status[i+1] = matrix[0][i+1] + status[i]
		}
	}

	for i := 1; i < n; i++ {
		status[0] = matrix[i][0] + status[0]
		for j := 1; j < n; j++ {
			status[j] = matrix[i][j] + int(math.Min(float64(status[j-1]), float64(status[j])))
		}

	}
	return status[n-1]
}
