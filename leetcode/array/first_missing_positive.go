// 41. 缺失的第一个正数
// 给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。
//
// 示例 1
// 输入: [1, 2, 0]
// 输出: 3
//
// 示例 2
// 输入: [3, 4, -1, 1]
// 输出: 2
//
// 示例 3
// 输入: [7, 8, 9, 11, 12]
// 输出: 1
// 提示：
//
// 0 <= nums.length <= 300
// -231 <= nums[i] <= 231 - 1
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/first-missing-positive
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package array

import "math"

// 暴力解题
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
func FirstMissingPositive1(nums []int) int {
	min := 1
	for {
		isMin := true
		for i := range nums {
			if nums[i] == min {
				min++
				isMin = false
				continue
			}
		}
		if isMin {
			return min
		}
	}
}

// 标记法
func FirstMissingPositive2(nums []int) int {
	length := len(nums)

	// 将不符合[1, length]区间内的数据标记为N+1
	for i := 0; i < length; i++ {
		if nums[i] < 1 || nums[i] > length {
			nums[i] = length + 1
		}
	}

	// 将属于[1, length]区间内的数打标记
	// 如何打标记?
	// 如果num[i]的值在[1, length]的区间内, 那么就可以用num数组的下标作为num[i]的值.
	// 而因为数组下标从0开始, 则需要将num[i]-1做为num[i]的值的下标
	// 注意, 这里标记的时候有个小技巧, 不能直接给对应的下标赋值
	// 例如num[0] = 3, num[2] = 1, 如果直接赋值, 那么num[2] = MinInt8 把原来的值1覆盖了, 就会产生bug
	// 所以这里使用了临时变量 n 来保存 num[num[i]-1]的值
	// 同时, 值判断必须使用for, 而不能使用if. if只能判断一次.
	// num[0] = 3, num[2] = 1, 第一次判断后, num[2] = -3, n = 1, 如果不继续判断, 如何安放原来num[2] = 1的值呢?
	// 所以需要使用for一直判断下去, 直到n的值不在[1, length]区间内为止.
	for i := 0; i < length; i++ {
		n := nums[i]
		for n >= 1 && n <= length {
			nums[n-1], n = math.MinInt8, nums[n-1]
		}
	}

	// 遍历nums, 找出第一个标记为0的数, 那么该数的下标则为缺失的最小正整数
	// 如果没有一个数符合, 则说明缺失的数为length+1
	for i := 0; i < length; i++ {
		if nums[i] != math.MinInt8 {
			return i + 1
		}
	}
	return length + 1
}

// 置换法
func FirstMissingPositive3(nums []int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for nums[i] >= 1 && nums[i] <= length && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return length + 1
}
