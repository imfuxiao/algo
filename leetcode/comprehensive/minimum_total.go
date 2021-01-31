// 120. 三角形最小路径和
// 给定一个三角形 triangle ，找出自顶向下的最小路径和。
//
// 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
// 也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
//
//
// 示例 1：
//
// 输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
// 输出：11
// 解释：如下面简图所示：
// 2
// 3 4
// 6 5 7
// 4 1 8 3
// 自顶向下的最小路径和为11（即，2+3+5+1= 11）。
// 示例 2：
//
// 输入：triangle = [[-10]]
// 输出：-10
//
//
// 提示：
//
// 1 <= triangle.length <= 200
// triangle[0].length == 1
// triangle[i].length == triangle[i - 1].length + 1
// -104 <= triangle[i][j] <= 104
//
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/triangle
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package comprehensive

import "math"

func MinimumTotal1(triangle [][]int) int {
	var (
		rows   = len(triangle)
		status = make([]int, len(triangle[rows-1])) // 动态规划存储中间状态, 使用一维数组压缩空间
		// 获取两个数字中的最小值
		min = func(num1, num2 int) int {
			if num1 > num2 {
				return num2
			}
			return num1
		}
	)

	// init status
	status[0] = triangle[0][0]

	for row := 1; row < rows; row++ {
		cols := len(triangle[row]) - 1
		// 从高往低处理, 否则会覆盖低位的值
		status[cols] = triangle[row][cols] + status[cols-1]
		for col := cols - 1; col > 0; col-- {
			status[col] = triangle[row][col] + min(status[col], status[col-1])
		}
		status[0] = triangle[row][0] + status[0]
	}

	result := math.MaxInt64
	for i := range status {
		if status[i] < result {
			result = status[i]
		}
	}
	return result
}
