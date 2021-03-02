package array

func searchInsert(nums []int, target int) int {
	length := len(nums)
	min, max, mid := 0, length-1, 0
	for min <= max {
		mid = (min + max) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return min
}
