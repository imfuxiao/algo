package linked

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		newNode = &ListNode{}
		add     = 0
	)
	n := newNode
	for l1 != nil && l2 != nil {
		l1.Val = l1.Val + l2.Val + add
		add = l1.Val / 10
		l1.Val = l1.Val % 10
		n.Next = l1

		n = n.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 == nil && l2 != nil {
		l2.Val = l2.Val + add
		n.Next = l2
		n = n.Next
	} else if l1 != nil && l2 == nil {
		l1.Val = l1.Val + add
		n.Next = l1
		n = n.Next
	} else if add != 0 {
		n.Next = &ListNode{Val: add}
		n = n.Next
	}

	for n.Val >= 10 {
		if n.Next != nil {
			n.Next.Val = n.Val/10 + n.Next.Val
		} else {
			n.Next = &ListNode{
				Val: n.Val / 10,
			}
		}
		n.Val = n.Val % 10
		n = n.Next
	}

	return newNode.Next
}
