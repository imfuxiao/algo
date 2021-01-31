// 206. 反转链表
// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
// 示例
// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/reverse-linked-list
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package linked

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	trail := &ListNode{
		Val: head.Val,
	}
	for {
		if head.Next == nil {
			break
		}
		head = head.Next
		trail = &ListNode{
			Val:  head.Val,
			Next: trail,
		}
	}
	return trail
}

func reverseList3(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		curr.Next, prev, curr = prev, curr, curr.Next
	}
	return prev
}
