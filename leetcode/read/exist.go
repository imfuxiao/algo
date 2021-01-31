package read

func Exist(board [][]byte, word string) bool {
	var (
		rows, cols = len(board), len(board[0])
		wordLength = len(word)
		ans        = false
		dfs        func(row, col, wordIndex int) bool
		visited    = make([][]bool, rows)
	)

	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	dfs = func(row, col, wordIndex int) bool {
		if word[wordIndex] != board[row][col] {
			return false
		}

		if wordIndex+1 == wordLength {
			return true
		}

		visited[row][col] = true
		defer func() {
			visited[row][col] = false
		}()

		check := false
		if row+1 < rows && !visited[row+1][col] && dfs(row+1, col, wordIndex+1) {
			check = true
		}
		if !check && 0 <= row-1 && !visited[row-1][col] && dfs(row-1, col, wordIndex+1) {
			check = true
		}
		if !check && col+1 < cols && !visited[row][col+1] && dfs(row, col+1, wordIndex+1) {
			check = true
		}
		if !check && 0 <= col-1 && !visited[row][col-1] && dfs(row, col-1, wordIndex+1) {
			check = true
		}
		return check
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if ans = dfs(row, col, 0); ans {
				return ans
			}
		}
	}

	return ans
}
