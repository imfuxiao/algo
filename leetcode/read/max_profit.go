package read

func MaxProfit(prices []int) int {
	var (
		length   = len(prices)
		ans      = 0
		max, min = func(n1, n2 int) int {
			if n1 > n2 {
				return n1
			}
			return n2
		}, func(n1, n2 int) int {
			if n1 < n2 {
				return n1
			}
			return n2
		}
	)

	if length <= 0 {
		return 0
	}

	preMinValue := prices[0]
	for i := 1; i < length; i++ {
		preMinValue = min(preMinValue, prices[i])
		ans = max(ans, prices[i]-preMinValue)
	}
	return ans
}
