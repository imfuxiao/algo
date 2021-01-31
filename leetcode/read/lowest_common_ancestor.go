package read

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	leftChild := lowestCommonAncestor(root.Left, p, q)
	rightChild := lowestCommonAncestor(root.Right, p, q)
	if leftChild != nil && rightChild != nil {
		return root
	}
	if leftChild == nil {
		return rightChild
	}
	return leftChild
}
