package linked

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	// 得到中间节点
	fast, slow := head, head
	for fast != nil && slow != nil {
		slow = slow.Next

		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	// 从中间节点开始逆序
	var pre *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = pre
		pre = slow
		slow = next
	}

	h := head
	for h != nil && pre != nil {
		next := h.Next
		p := pre.Next

		h.Next = pre
		pre.Next = next

		pre = p
		h = next
	}
	if h != nil {
		h.Next = nil
	}
}
