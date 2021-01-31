package backtracking

import "math"

// 我们有一个数字序列包含 n 个不同的数字，如何求出这个序列中的最长递增子序列长度？
// 比如 2, 9, 3, 6, 5, 1, 7 这样一组数字序列，它的最长递增子序列就是 2, 3, 5, 7，所以最长递增子序列的长度是 4。

// 回溯解决方法
// nums: 给定数字序列
// 返回值: 最大子序列长度
func GetMaxSub1(nums []int) int {
	length := len(nums)
	return g1(nums, length, 0, 1)
}

func g1(nums []int, numsLength, beginIndex, subLength int) int {
	if beginIndex >= numsLength {
		return subLength
	}

	value := subLength

	for i := beginIndex + 1; i < len(nums); i++ {
		if nums[i] > nums[beginIndex] {
			v := g1(nums, numsLength, i, subLength+1)
			if v > value {
				value = v
			}
		}
	}
	return value
}

// 动态规划解决方法
// nums: 给定数字序列
// 返回值: 最大子序列长度
func GetMaxSub2(nums []int) int {
	length := len(nums)
	status := make([]int, length)
	for i := range status {
		status[i] = 1
	}

	for i := 0; i < length; i++ {

		// 因为是找递增序列, 所以只需要比较当前元素之后的值
		// 而比当前值小元素, 对应status中的值保持不变
		for j := i + 1; j < length; j++ {
			if nums[i] < nums[j] && status[i] == status[j] { // status[i] == status[j] 保证必须在一个层级才能下钻
				status[j]++
			}
		}
	}

	// 取status中最大的值
	max := math.MinInt64
	for i := range status {
		if status[i] > max {
			max = status[i]
		}
	}
	return max
}
