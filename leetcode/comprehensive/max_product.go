// 152. 乘积最大子数组
// 给你一个整数数组 nums，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
// 示例 1:
// 输入: [2,3,-2,4]
// 输出: 6
// 解释:子数组 [2,3] 有最大乘积 6。
//
// 示例 2:
// 输入: [-2,0,-1]
// 输出: 0
// 解释:结果不能为 2, 因为 [-2,-1] 不是子数组。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/maximum-product-subarray
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package comprehensive

import "math"

// 时间超时
func MaxProduct1(nums []int) int {
	length := len(nums)
	result := math.MinInt64
	multiply := func(nums []int) int {
		res := 1
		for i := range nums {
			res = res * nums[i]
		}
		return res
	}
	max := func(num1, num2 int) int {
		if num1 > num2 {
			return num1
		}
		return num2
	}
	for step := 1; step <= length; step++ {
		for i := 0; i <= length-step; i++ {
			result = max(result, multiply(nums[i:i+step]))
		}
	}
	return result
}

// 我的动态规划思路
func MaxProduct2(nums []int) int {
	length := len(nums)
	result := math.MinInt32
	max := func(num1, num2 int) int {
		if num1 > num2 {
			return num1
		}
		return num2
	}

	status := make([]int, length)
	for i := 0; i < length; i++ {
		status[i] = nums[i]
		result = max(nums[i], result)
	}

	for step := 1; step < length; step++ {
		// 必须从大往小,否则会被覆盖
		for i := length - 1; i >= step; i-- {
			status[i] = nums[i] * status[i-1]
			result = max(result, status[i])
		}
	}

	return result
}

func MaxProduct3(nums []int) int {
	maxmum, minimum, length, answer := nums[0], nums[0], len(nums), nums[0]

	max, min := func(num1, num2 int) int {
		if num1 > num2 {
			return num1
		}
		return num2
	}, func(num1, num2 int) int {
		if num1 < num2 {
			return num1
		}
		return num2
	}

	for i := 1; i < length; i++ {
		m, n := maxmum, minimum
		maxmum = max(nums[i], max(nums[i]*m, nums[i]*n))
		minimum = min(nums[i], min(nums[i]*n, nums[i]*n))
		answer = max(maxmum, answer)
	}

	return answer
}
