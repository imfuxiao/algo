package read

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	var (
		newNode = &ListNode{
			Next: head,
		}
		index = 0
	)

	cur, end, pre := newNode.Next, newNode.Next, newNode
	for end != nil {
		index++
		if index%k == 0 {
			pre.Next = end
			cur, pre = reverse(cur, end.Next), cur
			end = cur
			continue
		}
		end = end.Next
	}

	return newNode.Next
}

func reverse(cur *ListNode, end *ListNode) *ListNode {
	p := end
	for cur != end {
		next := cur.Next
		cur.Next = p
		p = cur
		cur = next
	}
	return cur
}
