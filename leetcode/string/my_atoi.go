// 请你来实现一个 atoi 函数，使其能将字符串转换成整数。
//
// 首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。接下来的转化规则如下：
//
// 如果第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字字符组合起来，形成一个有符号整数。
// 假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成一个整数。
// 该字符串在有效的整数部分之后也可能会存在多余的字符，那么这些字符可以被忽略，它们对函数不应该造成影响。
// 假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换，即无法进行有效转换。
//
// 在任何情况下，若函数不能进行有效的转换时，请返回 0 。
//
// 注意：
//
// 本题中的空白字符只包括空格字符 ' ' 。
// 假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231, 231 − 1]。如果数值超过这个范围，请返回  231 − 1 或 −231 。
//
// 示例 1:
//
// 输入: "42"
// 输出: 42
// 示例 2:
//
// 输入: "   -42"
// 输出: -42
// 解释: 第一个非空白字符为 '-', 它是一个负号。
// 我们尽可能将负号与后面所有连续出现的数字组合起来，最后得到 -42 。
//
// 示例3:
//
// 输入: "4193 with words"
// 输出: 4193
// 解释: 转换截止于数字 '3' ，因为它的下一个字符不为数字。
// 示例 4:
//
// 输入: "words and 987"
// 输出: 0
// 解释: 第一个非空字符是 'w', 但它不是数字或正、负号。
// 因此无法执行有效的转换。
// 示例 5:
//
// 输入: "-91283472332"
// 输出: -2147483648
// 解释: 数字 "-91283472332" 超过 32 位有符号整数范围。
// 因此返回 INT_MIN (−231) 。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/string-to-integer-atoi
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package string

import (
	"math"
	"strings"
)

// 好雍容的代码......
func MyAtoi1(s string) int {
	const (
		dot     = '.'
		plus    = '+'
		minus   = '-'
		nine    = '9'
		zero    = '0'
		space   = ' '
		zeroInt = 0
	)

	var (
		isMinus = false // 是否为负数
		num     int     // 最终转换后的数字
		numStr  string  // 代转换前
	)

	// 过滤空白字符
	strs := strings.FieldsFunc(s, func(r rune) bool {
		if r == space {
			return true
		}
		return false
	})

	// 判断是否非空字符
	if len(strs) == zeroInt || len(strs[0]) == 0 {
		return zeroInt
	}

	numStr = strs[0]

	// 判断是否为字母开头
	if 'A' <= numStr[0] && numStr[0] <= 'z' {
		return zeroInt
	}

	// 处理+号, -号
	if strings.HasPrefix(numStr, string(plus)) {
		numStr = numStr[1:]
	} else if strings.HasPrefix(numStr, string(minus)) {
		numStr = numStr[1:]
		isMinus = true
	}

	// 如果第二位不是数字
	if len(numStr) > 0 && !('0' <= numStr[0] && numStr[0] <= '9') {
		return zeroInt
	}

	// 过滤字符串, 只包含数字和.
	strs = strings.FieldsFunc(numStr, func(r rune) bool {
		if (zero <= r && r <= nine) || r == dot {
			return false
		}
		return true
	})

	if len(strs) == zeroInt || len(strs[0]) == zeroInt {
		return zeroInt
	}

	numStr = strs[0]

	// 小数处理
	if index := strings.Index(numStr, string(dot)); index == 0 {
		return zeroInt
	} else if index > 0 {
		numStr = numStr[:index]
	}

	// 转整数
	toNum := func(char byte, index int) int {
		var n int
		switch char {
		case '0':
			n = 0
		case '1':
			n = 1
		case '2':
			n = 2
		case '3':
			n = 3
		case '4':
			n = 4
		case '5':
			n = 5
		case '6':
			n = 6
		case '7':
			n = 7
		case '8':
			n = 8
		case '9':
			n = 9
		}

		for i := 1; i < index; i++ {
			n = n * 10
		}
		return n
	}

	for i := range numStr {
		bitCount := len(numStr) - i
		if numStr[i] > '0' && bitCount > 10 {
			if isMinus {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
		num = num + toNum(numStr[i], bitCount)
	}

	if isMinus {
		num = num * -1
	}

	if num <= math.MinInt32 {
		return math.MinInt32
	} else if num >= math.MaxInt32 {
		return math.MaxInt32
	}

	return num
}

// 有限状态机
func MyAtoi2(s string) int {
	var (
		space    byte = ' '
		plus     byte = '+'
		minus    byte = '-'
		start         = "start"
		signed        = "signed"
		isNumber      = "isNumber"
		end           = "end"
		isMinus       = 1 // 是否负数: 1表示正数, -1表示负数, 不定义为bool类型, 是为了不用判断, 返回值可以直接乘
		ans           = 0 // 返回值
		// 有限状态机
		// 下标0代表输入: ' ' 空格
		// 下标1代表输入: +/- 符号
		// 下标2代表输入: 代表输入数字
		// 下标3代表输入: 其他字符
		statusTable = map[string][]string{
			start:    {start, signed, isNumber, end},
			signed:   {end, end, isNumber, end},
			isNumber: {end, end, isNumber, end},
			end:      {end, end, end, end},
		}
		currentStatus = start // 当前状态
	)

	// 根据字符获取对应的下标
	// 0: 代表' '
	// 1: 代表+/-符号
	// 2: 代表数字
	// 3: 代表其他
	getIndex := func(char byte) int {
		var index int
		switch {
		case char == space:
			index = 0
		case char == plus || char == minus:
			index = 1
		case '0' <= char && char <= '9':
			index = 2
		default:
			index = 3
		}
		return index
	}

	for i := range s {
		currentStatus = statusTable[currentStatus][getIndex(s[i])]
		switch {
		case currentStatus == isNumber:
			ans = ans*10 + int(s[i]-'0') // 累加
			if isMinus == -1 && ans*-1 <= math.MinInt32 {
				return math.MinInt32
			} else if ans > math.MaxInt32 {
				return math.MaxInt32
			}
		case currentStatus == signed:
			if s[i] == minus {
				isMinus = -1
			}
		case currentStatus == end:
			break
		default:
			continue
		}
	}

	return ans * isMinus
}
