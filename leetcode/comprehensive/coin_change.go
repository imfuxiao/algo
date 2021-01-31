// 322. 零钱兑换
// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回-1。
// 你可以认为每种硬币的数量是无限的。
//
// 示例 1：
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1
//
// 示例 2：
// 输入：coins = [2], amount = 3
// 输出：-1
//
// 示例 3：
// 输入：coins = [1], amount = 0
// 输出：0
//
// 示例 4：
// 输入：coins = [1], amount = 1
// 输出：1
//
// 提示：
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 231 - 1
// 0 <= amount <= 104
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/coin-change
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package comprehensive

import "math"

// 递归, leetcode会超时
func CoinChange(coins []int, amount int) int {
	var (
		calc   func(amount, count int)
		status = math.MaxInt64
	)

	calc = func(amount, count int) {
		if amount < 0 {
			return
		}
		if amount == 0 && count < status {
			status = count
			return
		}
		for i := range coins {
			calc(amount-coins[i], count+1)
		}
	}

	calc(amount, 0)
	if status == math.MaxInt64 {
		status = -1
	}
	return status
}

// 动态规划
func CoinChange1(coins []int, amount int) int {
	status := make([]int, amount+1)
	for i := range status {
		status[i] = amount + 1
	}

	status[0] = 0

	min := func(num1, num2 int) int {
		if num1 > num2 {
			return num2
		}
		return num1
	}
	for i := 1; i <= amount; i++ {
		for j := range coins {
			if i-coins[j] < 0 {
				continue
			}
			status[i] = min(status[i], 1+status[i-coins[j]])
		}
	}
	if status[amount] > amount {
		status[amount] = -1
	}
	return status[amount]
}
