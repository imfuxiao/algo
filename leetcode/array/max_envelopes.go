package array

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	ans := 0
	length := len(envelopes)
	if length <= 1 {
		return length
	}

	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] < envelopes[j][1])
	})

	// 求height的最长递增子序
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		ans = max(dp[i], ans)
	}

	return ans
}

//func max(num1, num2 int) int {
//	if num1 > num2 {
//		return num1
//	}
//	return num2
//}
