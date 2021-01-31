// 21. 合并两个有序链表
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 示例 1
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
//
// 示例 2
// 输入：l1 = [], l2 = []
// 输出：[]
//
// 示例 3
// 输入：l1 = [], l2 = [0]
// 输出：[0]
//
// 提示：
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package linked

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	newListNode := &ListNode{}
	n := newListNode
	for {
		switch {
		case l1.Val < l2.Val:
			n.Next = l1
			l1 = l1.Next
		case l2.Val < l1.Val:
			n.Next = l2
			l2 = l2.Next
		default:
			n.Next = l1
			l1 = l1.Next
			n = n.Next
			n.Next = l2
			l2 = l2.Next
		}
		n = n.Next
		if l1 == nil {
			n.Next = l2
			break
		} else if l2 == nil {
			n.Next = l1
			break
		}
	}
	return newListNode.Next
}
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	newNode := &ListNode{}
	n := newNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			n.Next = l1
			l1 = l1.Next
		} else if l2.Val < l1.Val {
			n.Next = l2
			l2 = l2.Next
		} else {
			n.Next = l1
			l1 = l1.Next
			n = n.Next
			n.Next = l2
			l2 = l2.Next
		}
		n = n.Next
	}
	if l1 == nil {
		n.Next = l2
	} else if l2 == nil {
		n.Next = l1
	}
	return newNode.Next
}
