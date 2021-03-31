package array

func find132pattern(nums []int) bool {
	var (
		length = len(nums)
		stack  = make([]int, 0, length)
	)
	for i := 0; i < length; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			if len(stack) > 1 && nums[stack[0]] < nums[i] {
				return true
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return false
}
