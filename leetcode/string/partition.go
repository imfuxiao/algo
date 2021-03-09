package string

func partition(s string) [][]string {
	var (
		length = len(s)
		dp     = make([][]bool, length)
		ans    = make([][]string, 0, length)
	)

	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}

	for right := 0; right < length; right++ {
		for left := right; left >= 0; left-- {
			if right == left {
				dp[left][right] = true
			} else if left+1 == right && s[left] == s[right] {
				dp[left][right] = true
			} else if dp[left+1][right-1] && s[left] == s[right] {
				dp[left][right] = true
			}
		}
	}

	var dfs func(left int)
	r := make([]string, 0, length)
	dfs = func(left int) {
		if left == length {
			ans = append(ans, append(make([]string, 0, len(r)), r...))
			return
		}
		for right := left; right < length; right++ {
			if dp[left][right] {
				r = append(r, s[left:right+1])
				dfs(right + 1)
				r = r[:len(r)-1]
			}
		}
	}

	dfs(0)

	return ans
}
