package linked

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	linkedLength := 0
	h := head
	for head != nil {
		head = head.Next
		linkedLength++
	}

	if linkedLength <= 1 {
		return true
	}

	mid := linkedLength / 2
	head = h
	var midNode *ListNode

	var pre *ListNode

	for head != nil {
		mid--

		next := head.Next
		head.Next = pre
		pre = head
		head = next
		if mid == 0 {
			midNode = next
			break
		}
	}

	ans := true
	if linkedLength%2 != 0 {
		midNode = midNode.Next
	}

	for {
		if midNode == nil {
			break
		}

		if pre.Val != midNode.Val {
			ans = false
			break
		}
		pre = pre.Next
		midNode = midNode.Next
	}

	return ans

}
