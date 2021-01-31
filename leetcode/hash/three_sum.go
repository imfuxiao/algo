// 15. 三数之和
// 给你一个包含 n 个整数的数组 nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
//
// 示例 1
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
//
// 示例 2
// 输入：nums = []
// 输出：[]
//
// 示例 3
// 输入：nums = [0]
// 输出：[]
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/3sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package hash

import (
	"math"
	"sort"
)

// 用hashMap省略一层循环
func ThreeSum1(nums []int) [][]int {
	var result [][]int

	length := len(nums)
	if length < 3 {
		return result
	}

	sort.Ints(nums)
	if nums[0] > 0 {
		return result
	}

	targetStatus := make(map[[3]int][][3]int, length)
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			target := 0 - nums[i] - nums[j]
			key := [3]int{target, math.MinInt64, math.MinInt64}
			targetStatus[key] = append(targetStatus[key], [3]int{i, j})
		}
	}

	for i, v := range nums {
		key := [3]int{v, math.MinInt64, math.MinInt64}
		if values, exists := targetStatus[key]; exists {
			for _, indexs := range values {
				if i != indexs[0] && i != indexs[1] && nums[indexs[0]]+nums[indexs[1]]+v == 0 {
					r := []int{nums[indexs[0]], nums[indexs[1]], v}
					sort.Ints(r)
					newKey := [3]int{r[0], r[1], r[2]}
					if _, e := targetStatus[newKey]; !e {
						result = append(result, r)
						targetStatus[newKey] = nil
					}
				}
			}
		}
	}
	return result
}

// 双指针线性逼近
func ThreeSum2(nums []int) [][]int {
	var (
		result [][]int
		length = len(nums)
	)
	if length < 3 {
		return result
	}

	// sort: 方便后续处理重复值
	sort.Ints(nums)
	// 排序后从小到大, 且需要 a + b + c = 0, 所以如果第一值大于0, 那么就不存在 a + b + c == 0 的情况了
	if nums[0] > 0 {
		return result
	}

	for i := 0; i < length; i++ {
		// 过滤重复值
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for secondIndex, thirdIndex := i+1, length-1; secondIndex < thirdIndex; secondIndex++ {
			if secondIndex > i+1 && nums[secondIndex] == nums[secondIndex-1] {
				continue
			}
			target := 0 - nums[i] - nums[secondIndex]
			for nums[thirdIndex] > target && thirdIndex > secondIndex {
				thirdIndex--
			}
			if nums[thirdIndex] == target && thirdIndex > secondIndex {
				result = append(result, []int{nums[i], nums[secondIndex], nums[thirdIndex]})
			}
		}
	}

	return result
}
