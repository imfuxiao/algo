package array

func letterCombinations(digits string) []string {
	length := len(digits)

	ans := make([]string, 0, length)

	chars := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	for _, char := range digits {
		if len(ans) == 0 {
			for _, c := range chars[byte(char)] {
				ans = append(ans, c)
			}
		} else {
			other := ans
			ans = ans[len(ans):]
			for _, b := range other {
				for _, c := range chars[byte(char)] {
					ans = append(ans, b+c)
				}
			}
		}
	}

	return ans
}
