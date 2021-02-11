package array

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	var (
		length = len(nums)
		ans = math.MaxInt64
		res = math.MaxInt64
		compare func(third, firstAndSecondSum int)
	)

	compare = func(third, firstAndSecondSum int) {
		v := firstAndSecondSum + nums[third]
		r := abs(target - v)
		if r < res {
			res = r
			ans = v
		}
	}

	// 1. 排序
	sort.Ints(nums)
	for first := 0; first < length; first++ {
		// 重复数据
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second, third := first+1, length-1; second < third; second++ {
			// 重复
			if second > first + 1 && nums[second] == nums[second-1] {
				continue
			}

			value := nums[first] + nums[second]

			if nums[third] + value == target {
				return target
			}

			compare(third, value)

			for nums[third] + value > target && third - 1 > second {
				third--
				compare(third, value)
			}
		}
	}

	return ans
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}


func searchInsert(nums []int, target int) int {
	var length = len(nums)

	min, max, mid := 0, length-1, 0

	for {
		if min > max {
			break
		}
		mid = (min + max) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			max = mid - 1
		} else if nums[mid] < target {
			min = mid + 1
		}
	}
	return min
}

func subsets(nums []int) [][]int {
	var(
		length = len(nums)
		ans = make([][]int, 0, 1<<length)
	)
	for mark := 0; mark < 1<<length; mark++ {
		res := make([]int, 0, length)
		for i, num := range nums {
			if mark >> i & 1 > 0 {
				res = append(res, num)
			}
		}
		ans = append(ans, res)
	}

	return ans
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	for i := range nums2 {
		index := 0
		for j := m-1; j >= 0; j-- {
			if nums1[j] <= nums2[i] {
				index = j+1
				break
			}
			if nums1[j] > nums2[i] {
				nums1[j+1] = nums1[j]
			}
		}
		nums1[index] = nums2[i]
		m = m+1
	}
}