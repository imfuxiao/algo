package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	nums := make([]int, 0, 10)
	var midFunc func(*TreeNode)
	midFunc = func(node *TreeNode) {
		if node == nil {
			return
		}
		midFunc(node.Left)
		nums = append(nums, node.Val)
		midFunc(node.Right)
	}
	midFunc(root)
	return nums[k]
}