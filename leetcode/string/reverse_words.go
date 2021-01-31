// 151. 翻转字符串里的单词
// 给定一个字符串，逐个翻转字符串中的每个单词。
//
// 说明：
//
// 无空格字符构成一个 单词 。
// 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
// 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
//
// 示例 1：
//
// 输入："the sky is blue"
// 输出："blue is sky the"
// 示例 2：
//
// 输入：" hello world! "
// 输出："world! hello"
// 解释：输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
// 示例 3：
//
// 输入："a good example"
// 输出："example good a"
// 解释：如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
// 示例 4：
//
// 输入：s = "  Bob    Loves  Alice   "
// 输出："Alice Loves Bob"
// 示例 5：
//
// 输入：s = "Alice does not even like bob"
// 输出："bob like even not does Alice"
//
// 提示：
//
// 1 <= s.length <= 104
// s 包含英文大小写字母、数字和空格 ' '
// s 中 至少存在一个 单词
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/reverse-words-in-a-string
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package string

import (
	"strings"
)

// 用栈实现
func ReverseWords1(s string) string {
	var (
		result             []byte
		length                  = len(s)
		stack                   = make([]byte, length)
		stackElementsCount      = 0
		space              byte = ' '

		// 入栈
		stackPush = func(char byte) {
			stack[stackElementsCount] = char
			stackElementsCount++
		}

		// 出栈
		stackPop = func() byte {
			char := stack[stackElementsCount-1]
			stackElementsCount--
			return char
		}

		// 单词反转
		reverse = func(chars []byte, start, end int) {
			for left, right := start, end; left < right; left, right = left+1, right-1 {
				chars[left], chars[right] = chars[right], chars[left]
			}
		}
	)

	// 入栈
	wordStart := 0 // 记录单词开始下标, 为单词反转使用
	for i := range s {
		if s[i] == space {
			// 当遇到空格时, 判断栈顶是否也是空格, 如果不是, 则先反转栈中的单词, 反转后添加一个空格
			if stackElementsCount > 0 && stack[stackElementsCount-1] != space {
				reverse(stack, wordStart, stackElementsCount-1)
				wordStart = stackElementsCount + 1
				stackPush(space)
			}
			continue
		}
		stackPush(s[i])
	}
	// 反转最后一个单词
	reverse(stack, wordStart, stackElementsCount-1)

	// 出栈
	for stackElementsCount > 0 {
		char := stackPop()
		if char == space && len(result) == 0 {
			continue
		}
		result = append(result, char)
	}

	return string(result)
}

// 原生字符处理
func ReverseWords2(s string) string {
	spaceByte := ' '
	// 过滤空格
	words := strings.FieldsFunc(s, func(r rune) bool {
		if r == spaceByte {
			return true
		}
		return false
	})

	// 反转单词顺序
	func(words []string) {
		for left, right := 0, len(words)-1; left < right; left, right = left+1, right-1 {
			words[left], words[right] = words[right], words[left]
		}
	}(words)

	// 拼接字符
	return strings.Join(words, string(spaceByte))
}

// 双队列实现
func ReverseWords3(s string) string {
	panic("not implements")
}
