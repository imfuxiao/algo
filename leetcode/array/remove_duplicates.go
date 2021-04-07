package array

func removeDuplicates(nums []int) int {
	var (
		length = len(nums)
		slow = 2
	)

	for fast := slow; fast < length; fast++ {
		if nums[slow-2] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow
}
