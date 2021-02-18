package array

import "sort"

func arrayPairSum(nums []int) int {
	ans := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		ans += nums[i]
	}
	return ans
}
