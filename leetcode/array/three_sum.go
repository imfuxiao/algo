// 15. 三数之和
// 给你一个包含 n 个整数的数组`nums`，判断`nums`中是否存在三个元素 a，b，c ，使得`a + b + c = 0` 请你找出所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
//
// 示例1
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
//
// 示例2
// 输入：nums = []
// 输出：[]
//
// 示例3:
// 输入：nums = [0]
// 输出：[]
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/3sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package array

import (
	"reflect"
	"sort"
)

// 暴力破解, 时间超时
func threeSum1(nums []int) [][]int {
	var (
		res    [][]int
		length = len(nums)
	)
	if len(nums) < 3 {
		return res
	}

	// 从小到大排序
	sort.Ints(nums)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					r := []int{nums[i], nums[j], nums[k]}
					isExists := false
					for i := range res {
						if reflect.DeepEqual(res[i], r) {
							isExists = true
							break
						}
					}
					if !isExists {
						res = append(res, []int{nums[i], nums[j], nums[k]})
					}
				}
			}
		}
	}
	return res
}

func threeSum2(nums []int) [][]int {
	var (
		res    [][]int
		length = len(nums)
	)

	if len(nums) < 3 {
		return res
	}

	// 从小到大排序
	sort.Ints(nums)

	// 排序后第一个数字如果大于0, 则就永不存在 a+b+c==0的情况
	if nums[0] > 0 {
		return res
	}

	for first := 0; first < length; first++ {
		if first-1 >= 0 && nums[first] == nums[first-1] {
			continue
		}

		target := nums[first] * -1 // second + third 的目标值, 一定为负数

		for second, third := first+1, length-1; second < third; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// a + b + c = 0
			// target = -a, 且数组已经排序(从小到大)
			// 所以
			// 当 b+c > target 时, c 应该减少
			// 当 b+c < target 时, b 应该增加
			// 且 b < c
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second >= third {
				break
			}
			if nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}
