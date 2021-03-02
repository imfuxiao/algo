package array

import "sort"

func TwoDifference(num []int, target int) [][]int {
	var (
		length = len(num)
		ans    = make([][]int, 0, length)
	)

	if length < 2 {
		return ans
	}

	// 排序
	sort.Ints(num)

	left, right := 0, 1

	for ; right < length; right++ {
		// 去除重复
		if left > 0 && num[left] == num[left-1] {
			continue
		}

		if target+num[left] > num[right] {
			continue
		}

		for left < right && num[left]+target < num[right] {
			left++
		}

		if target+num[left] == num[right] {
			ans = append(ans, []int{num[right], num[left]})
		}
	}

	return ans
}
