// 64. 最小路径和
// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
//
// 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
// 输出：7
// 解释：因为路径 1→3→1→1→1 的总和最小。

// 示例 2：
// 输入：grid = [[1,2,3],[4,5,6]]
// 输出：12
//
// 提示：
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 200
// 0 <= grid[i][j] <= 100
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/minimum-path-sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package comprehensive

// 动态规划
func MinPathSum1(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	// 状态保存, 用一行, 压缩状态
	status := make([]int, cols)
	status[0] = grid[0][0]

	// init first row
	for i := 1; i < cols; i++ {
		status[i] = status[i-1] + grid[0][i]
	}

	for row := 1; row < rows; row++ {
		for col := 0; col < cols; col++ {
			downValue := grid[row][col] + status[col]
			if col > 0 {
				// 注意这里: grid[row][col] + status[col-1], 不是grid[row][col-1]
				leftValue := grid[row][col] + status[col-1]
				if downValue > leftValue {
					downValue = leftValue
				}
			}
			status[col] = downValue
		}
	}

	return status[cols-1]
}
