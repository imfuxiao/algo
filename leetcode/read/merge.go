package read

import "sort"

func Merge(intervals [][]int) [][]int {
	var (
		rows        = len(intervals)
		ans         = make([][]int, 0, rows)
		ansRowIndex = 0
	)

	// 必须先排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	if rows > 0 {
		ans = append(ans, intervals[0])
	}
	for i := 1; i < rows; i++ {
		if intervals[i][0] <= ans[ansRowIndex][1] {
			secondElement := intervals[i][1]
			if secondElement < ans[ansRowIndex][1] {
				secondElement = ans[ansRowIndex][1]
			}
			ans = append(ans[0:ansRowIndex], []int{ans[ansRowIndex][0], secondElement})
		} else {
			ans = append(ans, intervals[i])
			ansRowIndex++
		}
	}
	return ans
}
