// 36. 有效的数独
// 判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。
//
// 数字1-9在每一行只能出现一次。
// 数字1-9在每一列只能出现一次。
// 数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。
//
// 说明:
//
// 一个有效的数独（部分已被填充）不一定是可解的。
// 只需要根据以上规则，验证已经填入的数字是否有效即可。
// 给定数独序列只包含数字1-9和字符'.'。
// 给定数独永远是9x9形式的。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/valid-sudoku
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package graph

func isValidSudoku1(board [][]byte) bool {
	var dot byte = '.'
	rows := make([][]bool, 9)  // 行
	cols := make([][]bool, 9)  // 列
	boxes := make([][]bool, 9) // 9个box

	// 初始化
	for i := 0; i < 9; i++ {
		rows[i] = make([]bool, 10)
		cols[i] = make([]bool, 10)
		boxes[i] = make([]bool, 10)
	}

	getBoxIndex := func(row, col int) int {
		return row/3*3 + col/3
	}

	for row := range board {
		for col := range board[row] {
			value := board[row][col]

			if value == dot {
				continue
			}

			value = value - '0'

			// 判断行是否符合1~9的要求
			if rows[row][int(value)] == false {
				rows[row][int(value)] = true
			} else {
				return false
			}

			// 判断列是否符合要求
			if cols[col][int(value)] == false {
				cols[col][int(value)] = true
			} else {
				return false
			}

			// 判断3*3宫内是否符合要求
			cell := getBoxIndex(row, col)
			if boxes[cell][int(value)] == false {
				boxes[cell][int(value)] = true
			} else {
				return false
			}

		}
	}
	return true
}
