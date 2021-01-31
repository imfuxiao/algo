package read

import "sort"

func ThreeSum(nums []int) [][]int {
	var (
		length = len(nums)
		result = make([][]int, 0, length)
	)

	if length < 3 {
		return result
	}

	// 排序
	sort.Ints(nums)

	if nums[0] > 0 {
		return result
	}

	for firstIndex := 0; firstIndex < length; firstIndex++ {
		// 去重
		if firstIndex > 0 && nums[firstIndex] == nums[firstIndex-1] {
			continue
		}

		for secondIndex, thrivedIndex := firstIndex+1, length-1; secondIndex < length; secondIndex++ {
			//  去重
			if secondIndex > firstIndex+1 && nums[secondIndex] == nums[secondIndex-1] {
				continue
			}

			target := 0 - firstIndex - nums[secondIndex]

			for secondIndex < thrivedIndex && nums[thrivedIndex] > target {
				thrivedIndex--
			}

			if secondIndex < thrivedIndex && nums[thrivedIndex] == target {
				result = append(result, []int{nums[firstIndex], nums[secondIndex], nums[thrivedIndex]})
			}
		}
	}
	return result
}
