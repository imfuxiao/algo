// 23. 合并K个升序链表
// 给你一个链表数组，每个链表都已经按升序排列。
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
// 示例 1:
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
// 1->4->5,
// 1->3->4,
// 2->6
// ]
// 将它们合并到一个有序链表中得到。
// 1->1->2->3->4->4->5->6
//
// 示例 2:
// 输入：lists = []
// 输出：[]
//
// 示例 3:
// 输入：lists = [[]]
// 输出: []
//
// 提示：
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package linked

import "math"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	mergeList := &ListNode{}
	trailNode := mergeList

	getMin := func(lists []*ListNode) (row int, value int) {
		value = math.MaxInt64
		for i := range lists {
			if lists[i] != nil && lists[i].Val < value {
				row, value = i, lists[i].Val
			}
		}
		return
	}

	isAllNil := func(lists []*ListNode) bool {
		for i := range lists {
			if lists[i] != nil {
				return false
			}
		}
		return true
	}

	for {
		// colIndex 内 全部为nil, 则不在
		if isAllNil(lists) {
			break
		}

		row, value := getMin(lists)
		lists[row] = lists[row].Next

		node := &ListNode{
			Val: value,
		}
		trailNode.Next = node
		trailNode = node
	}

	return mergeList.Next
}

func print(listNode *ListNode) []int {
	var nodes []int
	for {
		if listNode == nil {
			break
		}
		nodes = append(nodes, listNode.Val)
		listNode = listNode.Next
	}
	return nodes
}
