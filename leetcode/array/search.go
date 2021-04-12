package array

func search(nums []int, target int) bool {
	var (
		length      = len(nums)
		left, right = 0, length - 1
	)

	for left <= right {
		midIndex := (left + right) >> 1 // x / 2 == x >> 1
		if nums[midIndex] == target {
			return true
		}

		if nums[left] == nums[midIndex] && nums[midIndex] == nums[right] {
			left++
			right--
			continue
		}

		if nums[left] <= nums[midIndex] {
			if nums[left] <= target && target < nums[midIndex] {
				right = midIndex - 1
			} else {
				left = midIndex + 1
			}
		} else {
			if nums[midIndex] < target && target <= nums[right] {
				left = midIndex + 1
			} else {
				right = midIndex - 1
			}
		}
	}
	return false
}
