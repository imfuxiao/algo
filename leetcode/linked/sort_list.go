package linked

import "sort"

func sortList(head *ListNode) *ListNode {
	values := make([]*ListNode, 0, 100)
	for head != nil {
		values = append(values, head)
		head = head.Next
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i].Val < values[j].Val
	})
	length := len(values)
	for i:= 0; i < length - 1; i++ {
		values[i].Next = values[i+1]
	}
	return values[0]
}
