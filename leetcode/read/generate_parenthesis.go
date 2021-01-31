package read

func generateParenthesis(n int) []string {
	var (
		generate func(int, []byte, int, int, [][]byte) [][]byte
	)

	generate = func(index int, char []byte, open, close int, chars [][]byte) [][]byte {
		if index+1 == n*2 {
			chars = append(chars, char)
			return chars
		}

		if open < n {
			n1 := make([]byte, n*2)
			copy(n1, char)
			n1[index+1] = '('
			chars = generate(index+1, n1, open+1, close, chars)
		}

		if close < n {
			n2 := make([]byte, n*2)
			copy(n2, char)
			n2[index+1] = ')'
			chars = generate(index+1, n2, open, close+1, chars)
		}

		return chars
	}

	var chars [][]byte
	char := make([]byte, n*2)
	char[0] = '('
	chars = generate(0, char, 1, 0, chars)

	var strings []string
	for i := range chars {
		if validate(chars[i]) {
			strings = append(strings, string(chars[i]))
		}
	}
	return strings
}

func validate(char []byte) bool {
	balance := 0
	for i := range char {
		if char[i] == '(' {
			balance++
		} else if char[i] == ')' {
			balance--
		}
		if balance < 0 {
			return false
		}
	}
	return balance == 0
}
