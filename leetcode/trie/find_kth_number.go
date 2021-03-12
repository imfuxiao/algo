package trie

func findKthNumber(n int, k int) int {
	var (
		// 获取同层节点数, 比如1, 2之间的节点数
		getCount func(cur, next int) int
	)

	getCount = func(cur, next int) int {
		count := 0
		for cur <= n {
			count += min(n-cur+1, next-cur)
			cur *= 10
			next *= 10
		}
		return count
	}

	ans := 1
	k--
	for k > 0 {
		count := getCount(ans, ans+1)
		if k < count { // 去下一层级
			ans *= 10
			k--
		} else { // 去同层级的下个数字
			k -= count
			ans += 1
		}
	}

	return ans
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
