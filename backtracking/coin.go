package backtracking

import "math"

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

// 递归求解
func CoinChange1(coins []int, amount int) int {
	var (
		coinChange func([]int, int, int) int
	)
	if amount <= 0 {
		return -1
	}

	coinChange = func(coins []int, amount int, count int) int {
		if amount == 0 {
			return count
		} else if amount < 0 {
			return -1
		}
		var ans = math.MaxInt64
		for _, c := range coins {
			minCount := coinChange(coins, amount-c, count+1)
			if minCount != -1 {
				ans = min(minCount, ans)
			}
		}
		if ans == math.MaxInt64 {
			return -1
		}
		return ans
	}

	return coinChange(coins, amount, 0)
}

// 动态规划01版
func CoinChange3(coins []int, amount int) int {
	if amount <= 0 {
		return -1
	}

	// dp数组: 下标表示amount金额, 值表示对应硬币的数量
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // 硬币数量如果按1元的硬币数量算最多为amount个, 所以这里初始值设置为amount+1
	}

	// i表示金额
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			value := i - c
			if value < 0 { // 硬币面值太大, 不能使用
				continue
			}
			dp[i] = min(dp[value]+1, dp[i])
		}
	}
	if dp[amount] >= amount+1 {
		return -1
	}
	return dp[amount]
}
