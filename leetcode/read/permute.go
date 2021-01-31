package read

func Permute(nums []int) [][]int {
	var (
		length = len(nums)
		ansCap = func(length int) int {
			c := 1
			for i := 1; i <= length; i++ {
				c *= i
			}
			return c
		}(length) // cap == length!
		ans       = make([][]int, 0, ansCap)
		backtrack func(numCount int, nums []int)
	)

	backtrack = func(numCount int, nums []int) {
		if numCount >= length {
			newNums := make([]int, length)
			copy(newNums, nums)
			ans = append(ans, newNums)
			return
		}

		for i := numCount; i < length; i++ {
			nums[numCount], nums[i] = nums[i], nums[numCount]
			backtrack(numCount+1, nums)
			nums[numCount], nums[i] = nums[i], nums[numCount]
		}
	}

	for i := 0; i < length; i++ {
		nums[0], nums[i] = nums[i], nums[0]
		backtrack(1, nums)
		nums[0], nums[i] = nums[i], nums[0]
	}

	return ans
}
