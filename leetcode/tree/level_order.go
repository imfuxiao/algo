package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var (
		queue = make([]*TreeNode, 0, 10)
		ans   = make([][]int, 0, 10)
	)

	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		length := len(queue)
		level := make([]int, 0, length)
		for i := 0; i < length; i++ {

			// 出栈
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			// 入栈
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, level)
	}

	return ans
}
