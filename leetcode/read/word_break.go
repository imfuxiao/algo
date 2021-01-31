package read

func WordBreak(s string, wordDict []string) bool {
	var (
		length = len(s)
		dp     = make([]bool, length+1)
		dict   = make(map[string]bool, len(wordDict))
	)
	for i := 0; i < len(wordDict); i++ {
		dict[wordDict[i]] = true
	}

	dp[0] = true // 空字符表示匹配
	for i := 1; i < length+1; i++ {
		for j := 0; j < i; j++ {
			if dp[i] = dp[j] && dict[s[j:i]]; dp[i] {
				break
			}
		}
	}
	return dp[length]
}
