package string

import "math"

// 224. 基本计算器
func calculate(s string) int {
	var (
		length = len(s)
		stack  = make([]int, 0, length)
		sign   = 1
		ans    = 0
	)
	stack = append(stack, 1)

	for i := 0; i < length; i++ {
		switch s[i] {
		case ' ':
		case '+':
			sign = stack[len(stack)-1]
		case '-':
			sign = -stack[len(stack)-1]
		case '(':
			stack = append(stack, sign)
		case ')':
			stack = stack[:len(stack)-1]
		default:
			num := 0
			for i < length && '0' <= s[i] && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			i--
			ans += sign * num
		}
	}
	return ans
}

// 227. 基本计算器 II
func calculate2(s string) int {
	var (
		length = len(s)
		stack  = make([]int, 0, length)
		sign   = '+'
	)

	for i := 0; i < length; i++ {
		switch s[i] {
		case ' ':
		case '+':
			sign = '+'
		case '-':
			sign = '-'
		case '*':
			sign = '*'
		case '/':
			sign = '/'
		default:
			num := 0
			for i < length && '0' <= s[i] && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			i--
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack = append(stack[:len(stack)-1], stack[len(stack)-1]*num)
			case '/':
				stack = append(stack[:len(stack)-1], stack[len(stack)-1]/num)
			}
		}
	}

	ans := 0
	for i := range stack {
		ans += stack[i]
	}

	if ans > math.MaxInt32 {
		return math.MaxInt32
	}

	return ans
}

// 772. 基本计算器 III
func calculate3(s string) int {
	var (
		length                = len(s)
		sign             byte = '+'
		stack                 = make([]int, 0, length)
		parentheses           = make([]byte, 0, length)
		parenthesesIndex      = make([]int, 0, length)
	)

	for i := 0; i < length; i++ {
		switch s[i] {
		case ' ':
		case '+':
			sign = '+'
		case '-':
			sign = '-'
		case '*':
			sign = '*'
		case '/':
			sign = '/'
		case '(':
			parentheses = append(parentheses, sign)
			parenthesesIndex = append(parenthesesIndex, len(stack))
			sign = '+'
		case ')':
			sign = parentheses[len(parentheses)-1]
			parentheses = parentheses[:len(parentheses)-1]
			sum := 0
			start := parenthesesIndex[len(parenthesesIndex)-1]
			parenthesesIndex = parenthesesIndex[:len(parenthesesIndex)-1]
			for j := start; j < len(stack); j++ {
				sum += stack[j]
			}
			stack = stack[:start]
			switch sign {
			case '+':
				stack = append(stack, sum)
			case '-':
				stack = append(stack, -sum)
			case '*':
				stack = append(stack[:len(stack)-1], stack[len(stack)-1]*sum)
			case '/':
				stack = append(stack[:len(stack)-1], stack[len(stack)-1]/sum)
			}
		default:
			num := 0
			for i < length && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			i--

			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack = append(stack[:len(stack)-1], stack[len(stack)-1]*num)
			case '/':
				stack = append(stack[:len(stack)-1], stack[len(stack)-1]/num)
			}
		}
	}

	ans := 0
	for i := range stack {
		ans += stack[i]
	}
	if ans > math.MaxInt32 {
		return math.MaxInt32
	}
	return ans
}

// 770. 基本计算器 IV
