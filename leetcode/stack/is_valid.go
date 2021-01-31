// 20. 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串，判断字符串是否有效。
// 有效字符串需满足：
// 1. 左括号必须用相同类型的右括号闭合。
// 2. 左括号必须以正确的顺序闭合。
// 3. 注意空字符串可被认为是有效字符串。
//
// 示例 1
// 输入: "()"
// 输出: true
//
// 示例 2
// 输入: "()[]{}"
// 输出: true
//
// 示例 3
// 输入: "(]"
// 输出: false
//
// 示例 4
// 输入: "([)]"
// 输出: false
//
// 示例 5
// 输入: "{[]}"
// 输出: true
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/valid-parentheses
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package stack

func isValid1(s string) bool {
	length := len(s)
	if length == 0 || length%2 != 0 {
		return false
	}
	pairs := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack, stackIndex := make([]byte, length), 0
	for i := range s {
		if v, exists := pairs[s[i]]; exists {
			// push stack
			stack[stackIndex] = v
			stackIndex++
		} else {
			// pop stack check stackIndex
			if stackIndex <= 0 {
				return false
			}
			// pop stack
			if stack[stackIndex-1] != s[i] {
				return false
			}
			stackIndex--
		}
	}
	return stackIndex == 0
}
