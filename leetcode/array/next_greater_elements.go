package array

func nextGreaterElements(nums []int) []int {
	var (
		length = len(nums)
		stack  = make([]int, 0, length)
		ans    = make([]int, 0, length)
	)

	if length == 0 {
		return ans
	}

	stack = append(stack, 0)

	for i := 1; i < length; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		if nums[stack[len(stack)-1]] > nums[i-1] {
			ans = append(ans, nums[stack[len(stack)-1]])
		} else {
			ans = append(ans, -1)
		}
	}

	if length == 1 {
		ans = append(ans, -1)
	} else {
		ans = append(ans, ans[0])
	}

	return ans
}
