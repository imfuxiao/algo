package read

func ClimbStairs(n int) int {
	var (
		dp = make([]int, n+3)
	)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
