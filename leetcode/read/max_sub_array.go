package read

import "math"

func MaxSubArray(nums []int) int {
	var (
		length = len(nums)
		dp     = nums[0]
		ans    = math.MinInt64
		max    = func(n1, n2 int) int {
			if n1 > n2 {
				return n1
			}
			return n2
		}
	)
	for i := 1; i < length; i++ {
		dp = max(dp+nums[i], nums[i])
		ans = max(dp, ans)
	}

	return ans
}
