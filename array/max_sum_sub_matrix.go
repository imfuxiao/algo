package array

import "math"

func maxSumSubmatrix(matrix [][]int, k int) int {
	var (
		rows, cols = len(matrix), len(matrix[0])
		ans        = math.MinInt32
	)

	for r := 0; r < rows; r++ {
		nums := make([]int, cols)

		for i := r; i >= 0; i-- {
			for c := 0; c < cols; c++ {
				nums[c] += matrix[i][c]
			}
			ans = max(ans, maxSumSub(append([]int{}, nums...), k))
		}
	}

	return ans

}

func maxSumSub(nums []int, k int) int {
	var (
		length = len(nums)
		stack  = make([]int, 0, length)
	)

	put := func(num int) {
		if num > k {
			return
		}

		for len(stack) > 0 && stack[len(stack)-1] < num {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, num)
	}

	for i := 1; i < length; i++ {
		nums[i] += nums[i-1]
	}

	for left := 0; left < length; left++ {
		var value = 0
		if left > 0 {
			value = nums[left-1]
		}
		for right := left + 1; right <= length; right++ {
			put(nums[right-1] - value)
		}
	}

	if len(stack) > 0 {
		return stack[0]
	}

	return math.MinInt32
}
