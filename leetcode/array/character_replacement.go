package array

func characterReplacement(s string, k int) int {
	length := len(s)
	if length <= 0 {
		return 0
	}

	ans := 0
	status := [26]int{}
	maxSubStrCount := 0
	for left, right := 0, 0; right < length; right++ {
		index := s[right] - 'A'
		status[index]++

		maxSubStrCount = max(maxSubStrCount, status[index])

		if left < right && right-left+1 > maxSubStrCount+k {
			status[s[left]-'A']--
			left++
		}

		ans = max(ans, right-left+1)
	}

	return ans
}
