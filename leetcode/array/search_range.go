package array

func searchRange(nums []int, target int) []int {
	var (
		length       = len(nums)
		binarySearch func(nums []int, target int, lower bool) int
		noSearch     = []int{-1, -1}
	)

	if length == 0 {
		return noSearch
	}

	binarySearch = func(nums []int, target int, lower bool) int {
		min, max := 0, len(nums)-1
		res := len(nums)
		for min <= max{
			mid := (min + max) >> 1

			if nums[mid] > target || (lower && nums[mid] >= target) {
				max = mid - 1
				res = mid
			} else {
				min = mid + 1
			}
		}
		return res
	}

	firstIndex := binarySearch(nums, target, true)
	if firstIndex == -1 || firstIndex == len(nums) || nums[firstIndex] != target{
		return noSearch
	}
	lastIndex := binarySearch(nums, target, false) - 1
	if lastIndex == -1 {
		return noSearch
	}

	return []int{firstIndex, lastIndex}
}
