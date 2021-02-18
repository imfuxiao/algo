package array

func findMaxConsecutiveOnes(nums []int) int {
	length := len(nums)
	dp0, dp1 := 0, 0
	ans := 0
	for i := 0; i < length; i++ {
		if nums[i] == 0 {
			dp1 = dp0+1
			dp0 = 0
		} else {
			dp0++
			dp1++
		}
		ans = max(dp0, dp1)
	}
	return ans
}

