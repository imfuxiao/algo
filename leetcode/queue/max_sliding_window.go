// 239. 滑动窗口最大值
// 给你一个整数数组 nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。
// 你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
// 返回滑动窗口中的最大值。
//
// 示例 1：
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
// 解释：
// 滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7      3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
// 示例 2：
//
// 输入：nums = [1], k = 1
// 输出：[1]
//
// 示例 3：
//
// 输入：nums = [1,-1], k = 1
// 输出：[1,-1]
//
// 示例 4：
//
// 输入：nums = [9,11], k = 2
// 输出：[11]
//
// 示例 5：
//
// 输入：nums = [4,-2], k = 2
// 输出：[4]
//
// 提示：
//
// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// 1 <= k <= nums.length
//
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/sliding-window-maximum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package queue

import (
	"math"
)

// 暴力解题, 会超时
func MaxSlidingWindow1(nums []int, k int) []int {
	if k == 1 {
		return nums
	}

	maxFun := func(n []int) int {
		max := math.MinInt64
		for i := range n {
			if n[i] > max {
				max = n[i]
			}
		}
		return max
	}

	length := len(nums)
	res := make([]int, 0, length)

	for i := range nums {
		beginIndex, endIndex := i, i+k
		if endIndex > length {
			endIndex = length
		}
		res = append(res, maxFun(nums[beginIndex:endIndex]))
		if endIndex == length {
			break
		}
	}
	return res
}

// 使用双向队列
func MaxSlidingWindow2(nums []int, k int) []int {
	if k == 1 || len(nums) < k {
		return nums
	}

	// queue为双向队列, 用这个队列存储nums的下标, 队列中下标对应的值从左到右依次递减
	var queue []int
	// 从队列右侧push索引, 当索引对应nums中的值, 比queue中最后一个值大, 则queue会弹出这个元素, 直到index对应的值比queue中最后一个元素小或相等
	rightPush := func(index int) {
		length := len(queue)
		for length > 0 && nums[index] > nums[queue[length-1]] {
			queue = queue[:length-1]
			length = len(queue)
		}
		queue = append(queue, index)
	}

	for i := 0; i < k; i++ {
		rightPush(i)
	}

	result := make([]int, len(nums)-k+1)
	index := 0
	result[index] = nums[queue[0]]
	index++

	for i := k; i < len(nums); i++ {
		rightPush(i)
		for j := queue[0]; j < i-k+1; {
			queue = queue[1:]
			j = queue[0]
		}
		result[index] = nums[queue[0]]
		index++
	}
	return result
}
