package array

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	mono_stack := []int{}
	for i := 0; i < n; i++ {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			right[mono_stack[len(mono_stack)-1]] = i
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func maximalRectangle(matrix [][]byte) int {
	var (
		rows, cols = len(matrix), 0
		ans        int
	)

	if rows <= 0 {
		return 0
	}

	cols = len(matrix[0])
	dp := make([]int, cols)

	for _, row := range matrix {
		for i, v := range row {
			if v == '0' {
				dp[i] = 0
			} else {
				dp[i] += 1
			}
		}

		ans = max(ans, maximal(dp))
	}

	return ans

}

func maximal(heights []int) int {
	var (
		length = len(heights)
		stack  = make([]int, 0, length) // 单调递增栈
		// heights[i] 左侧, 右侧的index
		left, right = make([]int, length), make([]int, length)
		ans         = 0
	)

	// init reight value
	for i := range right {
		right[i] = length
	}

	for i := 0; i < length; i++ {
		// hights[i] 右侧最低的index
		for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}

		// height[i] 左侧最低的index
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	for i := 0; i < length; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}

	return ans
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
