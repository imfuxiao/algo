package read

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	var (
		ans     = math.MinInt64
		nodeSum func(node *TreeNode) int
	)

	nodeSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 屏蔽值为负数的节点
		leftNodeValue := max(nodeSum(node.Left), 0)
		rightNodeValue := max(nodeSum(node.Right), 0)

		ans = max(ans, node.Val+leftNodeValue+rightNodeValue) // 遍历时动态更新路径最大值

		return node.Val + max(leftNodeValue, rightNodeValue) // 节点值 = 节点自己的值 + max(左节点值, 右节点值)
	}
	nodeSum(root)
	return ans
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
