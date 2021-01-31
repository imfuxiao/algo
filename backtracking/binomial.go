package backtracking

import "math"

//"杨辉三角": 我们现在对它进行一些改造。每个位置的数字可以随意填写，经过某个数字只能到达下面一层相邻的两个数字。
// 假设你站在第一层，往下移动，我们把移动到最底层所经过的所有数字之和，定义为路径的长度。请你编程求出从最高层移动到最底层的最短路径长度。

// binomialArray[i][j] 中 i 表示层级, j 表示此层级中的元素
func binomial1(binomialArray [][]int) int {

	level := len(binomialArray)

	for i := 1; i < level; i++ {
		currentLevel := binomialArray[i]
		upLevel := binomialArray[i-1]
		for j := 0; j < len(currentLevel); j++ {
			if j-1 < 0 {
				currentLevel[j] = currentLevel[j] + upLevel[j]
			} else if j+1 == len(currentLevel) {
				currentLevel[j] = currentLevel[j] + upLevel[j-1]
			} else {
				value := currentLevel[j] + upLevel[j]
				if v := currentLevel[j] + upLevel[j-1]; value > v {
					value = v
				}
				currentLevel[j] = value
			}

		}

	}

	minValue := math.MaxInt64
	for i := range binomialArray[level-1] {
		if binomialArray[level-1][i] < minValue {
			minValue = binomialArray[level-1][i]
		}
	}

	return minValue
}
