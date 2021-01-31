// 18. 四数之和
// 给定一个包含n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d，使得a + b + c + d的值与 target 相等？
// 找出所有满足条件且不重复的四元组。
// 注意：
// 答案中不可以包含重复的四元组。
//
// 示例：
//
// 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
//
// 满足要求的四元组集合为：
// [
//  [-1,  0, 0, 1],
//  [-2, -1, 1, 2],
//  [-2,  0, 0, 2]
// ]
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/4sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package hash

import "sort"

func FourSum1(nums []int, target int) [][]int {
	var (
		result [][]int
		length = len(nums)
	)

	// 排序
	sort.Ints(nums)

	for firstIndex := 0; firstIndex < length; firstIndex++ {
		// 去重
		if firstIndex > 0 && nums[firstIndex] == nums[firstIndex-1] {
			continue
		}

		// 变为三数之和
		for secondIndex := firstIndex + 1; secondIndex < length; secondIndex++ {
			// 去重
			if secondIndex > firstIndex+1 && nums[secondIndex] == nums[secondIndex-1] {
				continue
			}

			for thirdIndex, fourIndex := secondIndex+1, length-1; thirdIndex < fourIndex; thirdIndex++ {
				// 去重
				if thirdIndex > secondIndex+1 && nums[thirdIndex] == nums[thirdIndex-1] {
					continue
				}

				fourTarget := target - nums[firstIndex] - nums[secondIndex] - nums[thirdIndex]
				for nums[fourIndex] > fourTarget && thirdIndex < fourIndex {
					fourIndex--
				}

				if nums[fourIndex] == fourTarget && thirdIndex < fourIndex {
					result = append(result, []int{nums[firstIndex], nums[secondIndex], nums[thirdIndex], nums[fourIndex]})
				}
			}

		}
	}

	return result
}
