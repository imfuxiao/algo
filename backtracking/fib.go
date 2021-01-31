package backtracking

// 最原始的递归
// n 为fib数列的第几位
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// 添加备忘录递归
func fib2(n int) int {
	var dict = make([]int, n+1)
	return fibDic(n, dict)
}

func fibDic(n int, dict []int) int {
	if n == 1 || n == 2 {
		return 1
	}
	if v := dict[n]; v != 0 {
		return v
	}
	dict[n] = fibDic(n-1, dict) + fibDic(n-2, dict)
	return dict[n]
}

// 通过dp数组求fib
func fib3(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i:=3; i<=n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fib4(n int) int {
	prev1, prev2 := 1, 1
	if n == 1 || n == 2 {
		return prev2
	}
	for i:=3; i<=n; i++ {
		prev2 = prev1 + prev2
	}
	return prev2
}