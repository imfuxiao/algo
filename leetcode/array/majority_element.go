// 169. 多数元素
// 给定一个大小为 n 的数组，找到其中的多数元素。
// 多数元素是指在数组中出现次数 大于[ n/2 ] 的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。

// 示例 1
// 输入：[3,2,3]
// 输出：3
//
// 示例 2：
// 输入：[2,2,1,1,1,2,2]
// 输出：2
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/majority-element
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package array

import (
	"math"
	"sort"
)

func MajorityElement1(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 只限入参不能为负数时
func MajorityElement2(nums []int) int {
	max := math.MinInt64
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}
	status := make([]int, max+1)
	for i := range nums {
		status[nums[i]]++
	}
	max = math.MinInt64
	index := 0
	for i := range status {
		if status[i] > max {
			max = status[i]
			index = i
		}
	}
	return index
}

// BM子符串匹配算法
func MajorityElement3(nums []int) int {
	count, candidate := 0, 0
	for i := range nums {
		if count == 0 {
			candidate = nums[i]
		}
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
