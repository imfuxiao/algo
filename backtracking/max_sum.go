package backtracking

import (
	"math"
	"strings"
)

// 如果给你一个包含正数和负数的整数数组，你能否找到一个具有最大和的连续子数组，
// 其中子数组最少包含一个元素，返回其最大和。
// 比如输入的子数组是 [-2, 1, -3, 1, -1, 6, 2, -5, 4]，输出是 8，因为连续子数组 [1, -1, 6, 2] 的和最大。
func MaxSum(num []int) int {
	var (
		length = len(num)
	)
	if length <= 0 {
		return -1
	}
	dp, ans := num[0], math.MinInt64
	for i := 1; i < length; i++ {
		dp = max(num[i], dp+num[i])
		ans = max(dp, ans)
	}
	return ans
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

// 给你一个整数数组 numbers，找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），返回该子数组的乘积。
// 示例1：输入: [2,7,-2,4]输出: 14解释: 子数组 [2,7] 有最大乘积 14。
func MaxMultiply(nums []int) int {
	var (
		length = len(nums)
	)
	if length <= 0 {
		return 0
	}
	dp, ans := nums[0], math.MinInt64
	for i := 1; i < length; i++ {
		dp = max(dp*nums[i], nums[i])
		ans = max(dp, ans)
	}
	return ans
}

func numSimilarGroups(strs []string) int {
	// TODO
	panic("not implements")
}

func countSubstrings(s string) int {
	var (
		length = len(s)
		dp     = make([][]bool, length)
	)
	if length <= 1 {
		return length
	}

	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}

	ans := 0
	for maxLength := 0; maxLength < length; maxLength++ {
		for left := 0; left+maxLength < length; left++ {
			right := left + maxLength
			if left == right {
				dp[left][right] = true
				ans++
			} else if left+1 == right {
				dp[left][right] = s[left] == s[right]
				if dp[left][right] {
					ans++
				}
			} else if dp[left+1][right-1] && s[left] == s[right] {
				dp[left][right] = true
				ans++
			}
		}
	}

	return ans
}

func isPalindrome(s string) bool {
	var (
		length = len(s)
		strs   = make([]byte, 0, length)
	)
	s = strings.ToLower(s)
	for i := range s {
		if ('0' <= s[i] && s[i] <= '9') || ('a' <= s[i] && s[i] <= 'z') {
			strs = append(strs, s[i])
		}
	}
	length = len(strs)
	for i := 0; i < length/2; i++ {
		if strs[i] != strs[length-i-1] {
			return false
		}
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome2(head *ListNode) bool {
	var (
		slow, fast, pre *ListNode
	)
	slow, fast = head, head
	for slow != nil && fast != nil && fast.Next != nil {

		fast = fast.Next.Next

		next := slow.Next
		slow.Next = pre
		pre = slow
		slow = next
	}

	// 只有一个元素
	if slow != nil && fast != nil && pre == nil {
		return true
	} else if fast != nil && fast.Next == nil { // 说明为奇数, 则慢指针需要往前走一步
		slow = slow.Next
	}

	for slow != nil && pre != nil {
		if slow.Val != pre.Val {
			return false
		}
		slow = slow.Next
		pre = pre.Next
	}
	if slow == nil && pre == nil {
		return true
	}

	return false
}
