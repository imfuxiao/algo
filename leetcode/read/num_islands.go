package read

func NumIslands(grid [][]byte) int {
	var (
		ans        = 0
		rows, cols = len(grid), len(grid[0])
		dfs        func(row, col int)
	)

	dfs = func(row, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols {
			return
		}
		if grid[row][col] == '1' {
			grid[row][col] = '0'
			dfs(row+1, col)
			dfs(row-1, col)
			dfs(row, col+1)
			dfs(row, col-1)
		}
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				ans++
				dfs(row, col)
			}
		}
	}

	return ans
}
