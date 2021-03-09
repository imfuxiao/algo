package string

import "math"

func minCut(s string) int {
	var (
		length = len(s)
		dp     = make([][]bool, length)
	)

	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}

	for right := 0; right < length; right++ {
		for left := right; left >= 0; left-- {
			if left == right {
				dp[left][right] = true
			} else if left+1 == right && s[left] == s[right] {
				dp[left][right] = true
			} else if dp[left+1][right-1] && s[left] == s[right] {
				dp[left][right] = true
			}
		}
	}

	var (
		dp2 = make([]int, length)
	)

	for right := 0; right < length; right++ {
		if dp[0][right] == true {
			dp2[right] = 0
		} else {
			dp2[right] = math.MaxInt32
			for j := 0; j < right; j++ {
				if dp[j+1][right] {
					dp2[right] = min(dp2[j]+1, dp2[right])
				}
			}
		}
	}
	return dp2[length-1]
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
