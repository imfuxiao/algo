package read

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RightSideView(root *TreeNode) []int {
	var (
		ans     = make([]int, 0, 100)
		queue   = make([]*TreeNode, 0, 100)
		enqueue = func(node *TreeNode) {
			if node != nil {
				queue = append(queue, node)
			}
		}
		dequeue = func() *TreeNode {
			node := queue[0]
			queue = queue[1:]
			return node
		}
	)
	enqueue(root)
	for len(queue) > 0 {
		for i, length := 0, len(queue); i < length; i++ {
			node := dequeue()
			enqueue(node.Left)
			enqueue(node.Right)
			if i+1 == length {
				ans = append(ans, node.Val)
			}
		}
	}
	return ans
}
