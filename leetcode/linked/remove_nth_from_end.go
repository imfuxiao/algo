package linked

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		pre *ListNode
	)
	for head != nil {
		next := head.Next
		head.Next, pre, head = pre, head, next
	}
	head, pre = pre, nil
	for head != nil {
		n--
		if n == 0 {
			head = head.Next
			continue
		}
		next := head.Next
		head.Next, pre, head = pre, head, next
	}
	return pre
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	queue := make([]*ListNode, 0, 100)
	for head != nil {
		queue = append(queue, head)
		head = head.Next
	}

	length := len(queue)
	mod := k % length
	if mod == 0 {
		return queue[0]
	}

	queue[length-1].Next = queue[0]
	index := length - mod
	queue[index-1].Next = nil
	return queue[index]

}

func deleteDuplicates(head *ListNode) *ListNode {
	var (
		newNode = &ListNode{}
	)

	node, isRepeat := newNode, false
	for head != nil {
		for head.Next != nil && head.Next.Val == head.Val {
			isRepeat = true
			head.Next = head.Next.Next
		}
		next := head.Next
		if !isRepeat {
			node.Next = head
			node = node.Next
		} else {
			node.Next = next
		}
		head = next
		isRepeat = false
	}

	return newNode.Next
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var (
		newHead = &ListNode{
			Next: head,
		}
		reverse         func(begin, end, pre *ListNode)
		index           = 0
		begin, end, pre *ListNode
	)

	reverse = func(begin, end, pre *ListNode) {
		e := end.Next
		p := pre
		b := begin
		for begin != nil && begin != e {
			next := begin.Next
			begin.Next = pre
			pre = begin
			begin = next
		}
		p.Next = pre
		b.Next = e
	}

	h := newHead
	for h != nil && h.Next != nil {
		if index+1 == m {
			pre = h
			begin = h.Next
		}
		if index+1 == n {
			end = h.Next
			reverse(begin, end, pre)
			break
		}
		index++
		h = h.Next
	}

	return newHead.Next
}
