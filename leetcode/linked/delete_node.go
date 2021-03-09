package linked

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	p := &ListNode{
		Next: node,
	}
	// 快慢指针, 快每次2步, 慢每次一步, 等快的到达终点, 则慢指针指向中间节点
	fast, slow := p, p
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		if fast == nil || fast.Next == nil {
			slow.Next = slow.Next.Next
		} else {
			slow = slow.Next
		}
	}
}
