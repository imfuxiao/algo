package backtracking

import (
	"fmt"
)

// 根据row, column判断此位置是否可以放置皇后, false 不满足放置条件
// 因为放置皇后是从第1行开始, 所以检测时只需要检测[1, row-1]行即可
// queens: 棋子queen所在位置, 下标为行数, 值为列数
func checkPositionIsOk(queens *[9]int, row, column int) bool {
	if row == 1 {
		return false
	}

	// 左上对角列坐标, 右上对角列坐标
	leftUpColumn, rightUpColumn := column-1, column+1
	for rowIndex := row - 1; rowIndex > 0; rowIndex-- {
		// 判断所在列是否已放置皇后
		if queens[rowIndex] == column {
			return false
		}
		// 判断左上对角线
		if leftUpColumn >= 1 && queens[rowIndex] == leftUpColumn {
			return false
		} else {
			leftUpColumn--
		}
		// 判断右上对角线
		if rightUpColumn <= 8 && queens[rowIndex] == rightUpColumn {
			return false
		} else {
			rightUpColumn--
		}
	}
	return true
}

func Queen() {
	queues := [9]int{}
	for column := 1; column < 9; column++ {
		queues[1] = column
		put(&queues, 1)
		if queues[8] != 0 {
			fmt.Println(queues)
		}
		queues = [9]int{}
	}
}

func put(queues *[9]int, row int) {
	if row > 8 {
		return
	}

	for column := 1; column < 9; column++ {
		if checkPositionIsOk(queues, row, column) {
			queues[row] = column
		}
	}
	put(queues, row+1)
}
