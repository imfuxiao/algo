package array

func findDisappearedNumbers(nums []int) []int {
	length := len(nums)
	ans := make([]int, 0, length)

	for i := 0; i < length; i++ {
		value := (nums[i] - 1) % length
		nums[value] += length
	}

	for i := 0; i < length; i++ {
		if nums[i] <= length {
			ans = append(ans, i+1)
		}
	}

	return ans
}

func totalFruit(tree []int) int {
	ans := 0
	length := len(tree)
	status := make(map[int]int, length)
	for left, right := 0, 0; right < length; right++ {
		status[tree[right]]++
		for len(status) > 2 {
			if status[tree[left]] == 1 {
				delete(status, tree[left])
			} else {
				status[tree[left]]--
			}
			left++
		}
		ans = max(ans, right - left + 1)
	}
	return ans
}

