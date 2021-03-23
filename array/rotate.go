package array

func rotate(matrix [][]int) {
	var (
		rows      = len(matrix)
		tempArray = make([][]int, len(matrix))
	)

	for i := range matrix {
		tempArray[rows-i-1] = append([]int{}, matrix[i]...)
	}

	for r := range tempArray {
		for col := range tempArray[r] {
			matrix[col][r] = tempArray[r][col]
		}
	}
}
