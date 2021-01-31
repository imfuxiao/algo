// 32. 最长有效括号
// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
// 示例 1
// 输入：s = "(()"
// 输出：2
// 解释：最长有效括号子串是 "()"
//
// 示例 2
// 输入：s = ")()())"
// 输出：4
// 解释：最长有效括号子串是 "()()"
//
// 示例 3
// 输入：s = ""
// 输出：0
//
// 提示：
// 0 <= s.length <= 3 * 104
// s[i] 为 '(' 或 ')'
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-valid-parentheses
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package stack

func longestValidParentheses1(s string) int {
	length, maxCount := len(s), 0
	if length <= 1 {
		return maxCount
	}

	// leftStack 为左括号'('匹配栈, 栈中元素为'('在字符串中下标, 因为手动添加一个元素, 所以长度要+1
	leftStack, stackLength, left := make([]int, length+1), 0, byte('(')

	// 因为计算最长匹配长度公式: 右括号')'所在字符下标 - "弹栈后"栈顶元素值(注意是弹栈后的栈顶元素的值)
	// 所以为了计算公式计算方便, 手动添加-1的值, 并且下面弹栈的时候, 当栈的元素为空时, 也需要将右括号的下标压入栈中
	leftStack[stackLength] = -1
	stackLength++

	for i := range s {
		if s[i] == left {
			// push stack
			leftStack[stackLength] = i
			stackLength++
		} else {

			// pop stack
			stackLength--

			if stackLength <= 0 {
				// push stack
				leftStack[0] = i
				stackLength = 1
				continue
			}

			// 取栈顶元素, 计算最长匹配长度
			num := i - leftStack[stackLength-1]
			if num > maxCount {
				maxCount = num
			}
		}
	}

	return maxCount
}
