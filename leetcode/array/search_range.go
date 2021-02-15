package array

import "reflect"

func searchRange(nums []int, target int) []int {
	var (
		length       = len(nums)
		binarySearch func(nums []int, target int) []int
		noSearch     = []int{-1, -1}
	)

	if length == 0 {
		return noSearch
	}

	binarySearch = func(nums []int, target int) []int {
		answer := make([]int, 0, length)
		min, max := 0, len(nums)-1
		res := -1
		for {
			if min > max {
				break
			}

			mid := (min + max) >> 1

			if nums[mid] == target {
				res = mid
				break
			}

			if nums[mid] > target {
				max = mid - 1
			} else if nums[mid] < target {
				min = mid + 1
			}
		}

		if res == -1 {
			return []int{-1, -1}
		}

		if got := binarySearch(nums[:res], target); !reflect.DeepEqual(got, noSearch) {
			answer = append(answer, got[0])
		} else {
			answer = append(answer, res)
		}

		if got := binarySearch(nums[res+1:], target); !reflect.DeepEqual(got, noSearch) {
			answer = append(answer, got[len(got)-1])
		} else {
			answer = append(answer, res)
		}

		return answer
	}

	return binarySearch(nums, target)
}
