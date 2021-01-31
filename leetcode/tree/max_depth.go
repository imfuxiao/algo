// 104. 二叉树的最大深度
// 给定一个二叉树，找出其最大深度。
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例：
// 给定二叉树 [3,9,20,null,null,15,7]，
//
//     3
//    / \
//   9   20
//      /  \
//     15   7
//
// 返回它的最大深度 3 。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree
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
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 1 + maxDepth1(root.Left)
	rightDepth := 1 + maxDepth1(root.Right)
	if rightDepth > max {
		max = rightDepth
	}
	return max
}

// 广度搜索算法
// 利用队列存储每一层的节点
// 有点像按层遍历的思想
func maxDepth2(root *TreeNode) int {
	var (
		deep       = 0
		queue      []*TreeNode // 队列: 保存每层的节点
		queueCount = 0         // 队列中元素个数
		// 入队
		enqueue = func(node *TreeNode) {
			if node != nil {
				queue = append(queue, node)
				queueCount++
			}
		}
		// 出队
		dequeue = func() *TreeNode {
			if queueCount != 0 {
				node := queue[0]
				queueCount--
				queue = queue[1:]
				return node
			}
			return nil
		}
	)

	enqueue(root)
	for queueCount > 0 {
		deep++
		// 遍历当前层全部节点, 并添加下一层全部节点
		for count := queueCount; count > 0; count-- {
			node := dequeue()
			enqueue(node.Left)
			enqueue(node.Right)
		}
	}

	return deep
}
