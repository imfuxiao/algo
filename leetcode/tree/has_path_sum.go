// 112. 路径总和
// 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例:
// 给定如下二叉树，以及目标和 sum = 22，
//
//       5
//      / \
//     4   8
//    /   / \
//   11  13  4
//  /  \      \
// 7    2      1
// 返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/path-sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归
func hasPathSum1(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Right == nil && root.Left == nil {
		return root.Val == sum
	}
	sum = sum - root.Val
	return hasPathSum1(root.Left, sum) || hasPathSum1(root.Right, sum)
}

// 广度优先
// 广度优先需要利用queue
func hasPathSum2(root *TreeNode, sum int) bool {
	var (
		queue = struct {
			nodes []*TreeNode
			vals  []int
			count int
		}{}
		enqueue = func(node *TreeNode, val int) {
			if node != nil {
				queue.nodes = append(queue.nodes, node)
				queue.vals = append(queue.vals, val)
				queue.count++
			}
		}
		dequeue = func() (node *TreeNode, val int) {
			if queue.count > 0 {
				node = queue.nodes[0]
				queue.nodes = queue.nodes[1:]

				val = queue.vals[0]
				queue.vals = queue.vals[1:]

				queue.count--
			}
			return
		}
	)

	// 广度优先
	enqueue(root, 0)
	for queue.count > 0 {
		node, val := dequeue()
		val = val + node.Val
		if node.Left == nil && node.Right == nil && val == sum {
			return true
		}
		enqueue(node.Left, val)
		enqueue(node.Right, val)
	}
	return false
}
