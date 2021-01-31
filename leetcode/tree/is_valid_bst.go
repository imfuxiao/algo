// 98. 验证二叉搜索树
// 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
// 假设一个二叉搜索树具有如下特征：
//
// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
//
// 示例 1:
// 输入:
//   2
//  / \
// 1   3
// 输出: true
//
// 示例 2:
//
// 输入:
//   5
//  / \
// 1   4
//    / \
//   3   6
// 输出: false
//
// 解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/validate-binary-search-tree
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package tree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归
func isValidBST1(root *TreeNode) bool {
	// 比较函数
	var compareFunc func(node *TreeNode, min, max int) bool

	compareFunc = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}
		return compareFunc(node.Left, min, node.Val) && compareFunc(node.Right, node.Val, max)
	}

	return compareFunc(root, math.MinInt64, math.MaxInt64)
}

// 利用中序遍历
func isValidBST2(root *TreeNode) bool {
	var (
		sortsNode []*TreeNode                      // 保存中序遍历后的数据
		midFunc   func(root *TreeNode) []*TreeNode // 中序遍历
	)

	midFunc = func(root *TreeNode) []*TreeNode {
		var result []*TreeNode
		if root != nil {
			result = append(result, midFunc(root.Left)...)
			result = append(result, root)
			result = append(result, midFunc(root.Right)...)
		}
		return result
	}

	sortsNode = midFunc(root)

	for i, length := 1, len(sortsNode); i < length; i++ {
		if sortsNode[i].Val <= sortsNode[i-1].Val {
			return false
		}
	}
	return true
}
