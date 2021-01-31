package read

func LongestPalindrome(s string) string {
	var (
		length = len(s)
		dp     = make([][]bool, length)
		ans    string
	)
	for i := range dp {
		dp[i] = make([]bool, length)
	}

	for maxLength := 0; maxLength < length; maxLength++ {
		for rightIndex := 0; rightIndex+maxLength < length; rightIndex++ {
			leftIndex := rightIndex + maxLength
			if rightIndex == leftIndex {
				dp[rightIndex][leftIndex] = true
			} else if rightIndex+1 == leftIndex {
				dp[rightIndex][leftIndex] = s[rightIndex] == s[leftIndex]
			} else {
				dp[rightIndex][leftIndex] = s[rightIndex] == s[leftIndex] && dp[rightIndex+1][leftIndex-1]
			}

			if dp[rightIndex][leftIndex] && maxLength+1 > len(ans) {
				ans = s[rightIndex : leftIndex+1]
			}
		}
	}

	return ans
}
