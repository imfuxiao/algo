package read

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := range s {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		} else if s[i] == '}' || s[i] == ']' || s[i] == ')' {
			if len(stack) <= 0 {
				return false
			}
			l := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (s[i] == '}' && l == '{') || (s[i] == ']' && l == '[') || (s[i] == ')' && l == '(') {
				continue
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
