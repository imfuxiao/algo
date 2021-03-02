package array

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	ans := math.MaxInt32
	length := len(nums)

	sum := 0
	for left, right := 0, 0; right < length; right++ {
		if nums[right] == target {
			return 1
		}

		sum += nums[right]

		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if ans == math.MaxInt32 {
		return 0
	}

	return ans
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
