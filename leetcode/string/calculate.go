package string

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
