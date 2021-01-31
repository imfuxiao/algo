package read

func SubArraySum(nums []int, k int) int {
	status := make(map[int]int, len(nums))
	status[k] = 1
	ans := 0
	preSum := 0
	for i := range nums {
		preSum += nums[i]
		ans += status[preSum]
		status[k+preSum]++
	}
	return ans
}
