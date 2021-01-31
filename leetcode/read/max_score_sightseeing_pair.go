package read

func MaxScoreSightseeingPair(A []int) int {
	var (
		length = len(A)
		max    = func(n1, n2 int) int {
			if n1 > n2 {
				return n1
			}
			return n2
		}
		ans         = 0
		preMaxValue = 0
	)
	if length <= 0 {
		return preMaxValue
	}

	for i := 0; i < length; i++ {
		ans = max(preMaxValue+A[i]-i, ans)
		preMaxValue = max(preMaxValue, A[i]+i) // 注意这里的是+, 不是减, 因为公式为: A[i]+i+A[j]-j, preMaxValue = A[i]+i
	}

	return ans
}
