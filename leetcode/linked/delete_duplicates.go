package linked

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	var (
		pre     = &ListNode{}
		p, curr = head, head
	)

	n := pre
	duplicateValue := 999
	for curr != nil {
		if curr.Next != nil && curr.Next.Val == curr.Val {
			duplicateValue = curr.Val
			curr.Next = curr.Next.Next
			continue
		}
		p = curr
		curr = curr.Next
		if p.Val != duplicateValue {
			n.Next = p
			n = n.Next
		} else {
			n.Next = nil
		}
	}

	return pre.Next
}
