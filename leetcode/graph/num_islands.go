// 200. 岛屿数量
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
// 示例 1：
// 输入：grid = [
//   ["1","1","1","1","0"],
//   ["1","1","0","1","0"],
//   ["1","1","0","0","0"],
//   ["0","0","0","0","0"]
// ]
// 输出：1
//
// 示例 2：
// 输入：grid = [
//   ["1","1","0","0","0"],
//   ["1","1","0","0","0"],
//   ["0","0","1","0","0"],
//   ["0","0","0","1","1"]
// ]
// 输出：3
//
//
// 提示：
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] 的值为 '0' 或 '1'
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/number-of-islands
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package graph

// 广度优先搜索
func numIslands1(grid [][]byte) int {
	var (
		rowLength  = len(grid)
		colLength  = 0
		visited    = make([][]bool, rowLength) // 记录当前节点是否已经被访问过
		queueCount = 0                         // 队列中元素数量
		queue      = make([]struct {
			row int
			col int
		}, 0, len(grid))
		islandCount = 0 // 岛屿数量

		// 广度优先搜索
		rfs func(row, col int)
	)

	if rowLength <= 0 {
		return 0
	} else {
		colLength = len(grid[0])
	}

	// 状态初始化
	for i := range visited {
		visited[i] = make([]bool, colLength)
	}

	// 入队
	enqueue := func(row, col int) {
		if row >= 0 && row < rowLength && col >= 0 && col < colLength {
			queueCount++
			queue = append(queue, struct {
				row int
				col int
			}{row: row, col: col})
		}
	}

	// 出队
	dequeue := func() (row, col int) {
		point := queue[0]
		queue = queue[1:]
		queueCount--
		return point.row, point.col
	}

	// 广度优先搜索
	rfs = func(row, col int) {
		enqueue(row, col)
		visited[row][col] = true

		for queueCount > 0 {
			for i, length := 0, queueCount; i < length; i++ {
				r, c := dequeue()
				// 当前节点如果为1, 则改变当前节点的值并添加当前节点的上, 下, 左, 右节点
				if grid[r][c] == '1' {
					grid[r][c] = '0'
					enqueue(r-1, c)
					enqueue(r+1, c)
					enqueue(r, c+1)
					enqueue(r, c-1)
				}
			}
		}
	}

	for row := 0; row < rowLength; row++ {
		for col := 0; col < colLength; col++ {
			if !visited[row][col] && grid[row][col] == '1' {
				rfs(row, col)
				islandCount++
			}
		}
	}
	return islandCount
}

// 深度优先搜索
func numIslands2(grid [][]byte) int {
	var (
		rowLength   = len(grid)
		colLength   = 0
		visited     = make([][]bool, rowLength) // 保存访问过的节点
		dfs         func(row, col int)          // 广度优先搜索
		islandCount = 0
	)
	if rowLength <= 0 {
		return 0
	} else {
		colLength = len(grid[0])
	}

	for i := range visited {
		visited[i] = make([]bool, colLength)
	}

	dfs = func(row, col int) {
		if row < 0 || row >= rowLength || col < 0 || col >= colLength {
			return
		}
		visited[row][col] = true
		if grid[row][col] == '1' {
			grid[row][col] = '0'
			dfs(row, col+1)
			dfs(row, col-1)
			dfs(row+1, col)
			dfs(row-1, col)
		}
	}

	for r := 0; r < rowLength; r++ {
		for c := 0; c < colLength; c++ {
			if !visited[r][c] && grid[r][c] == '1' {
				islandCount++
				dfs(r, c)
			}
		}
	}

	return islandCount
}
