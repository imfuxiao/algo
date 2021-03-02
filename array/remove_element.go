package array

func removeElement(nums []int, val int) int {
	length := len(nums)
	i := 0
	for i < length {
		for nums[i] == val && i < length {
			nums[i], nums[length-1] = nums[length-1], nums[i]
			length--
		}
		i++
	}
	return length
}
