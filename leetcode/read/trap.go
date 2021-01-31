package read

// BF
func Trap1(height []int) int {
	var (
		length = len(height)
		ans    = 0
		max    = func(n1, n2 int) int {
			if n1 > n2 {
				return n1
			}
			return n2
		}
		min = func(n1, n2 int) int {
			if n1 < n2 {
				return n1
			}
			return n2
		}
	)

	for i := 1; i < length; i++ {
		// left max value
		leftMaxValue := height[i]
		for leftIndex := i; leftIndex >= 0; leftIndex-- {
			leftMaxValue = max(leftMaxValue, height[leftIndex])
		}

		// right max value
		rightMaxValue := height[i]
		for rightIndex := i; rightIndex < length; rightIndex++ {
			rightMaxValue = max(rightMaxValue, height[rightIndex])
		}

		ans += min(rightMaxValue, leftMaxValue) - height[i]
	}

	return ans
}

// 动态编程
func Trap2(height []int) int {
	var (
		length   = len(height)
		leftMax  = make([]int, length)
		rightMax = make([]int, length)
		ans      = 0
		max, min = func(n1, n2 int) int {
			if n1 > n2 {
				return n1
			}
			return n2

		}, func(n1, n2 int) int {
			if n1 < n2 {
				return n1
			}
			return n2
		}
	)

	if length <= 0 {
		return ans
	}

	// 缓存左侧最高值
	leftMax[0] = height[0]
	for i := 1; i < length; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	// 缓存右侧最高值
	rightMax[length-1] = height[length-1]
	for i := length - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	// 累加
	for i := 1; i < length; i++ {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}

	return ans
}
