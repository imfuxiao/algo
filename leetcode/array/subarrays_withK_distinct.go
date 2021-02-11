package array

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 {
		return k
	}
	ans := 0
	length := len(s)
	status := make(map[byte]int, length)

	for left, right := 0, 0; right < length; right++ {
		status[s[right]]++

		for left < right && len(status) > k {
			if _, exists := status[s[left]]; exists && status[s[left]] == 1 {
				delete(status, s[left])
			} else if exists && status[s[left]] > 1 {
				status[s[left]]--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func subarraysWithKDistinct(A []int, K int) int {
	return distinct(A, K) - distinct(A, K-1)
}

func distinct(a []int, k int) int {
	ans := 0
	length := len(a)
	status := make(map[int]int, length)

	for left, right := 0, 0; right < length; right++ {
		status[a[right]]++
		for len(status) > k {
			if _, exists := status[a[left]]; exists && status[a[left]] == 1 {
				delete(status, a[left])
			} else {
				status[a[left]]--
			}
			left++
		}
		ans += right - left
	}
	return ans
}
