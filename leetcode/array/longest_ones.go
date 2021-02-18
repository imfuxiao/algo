package array

func longestOnes(A []int, K int) int {
	length := len(A)
	maxWindows := 0
	left, right := 0, 0
	for ; right < length; right++ {
		maxWindows += A[right]
		if right-left+1 > maxWindows+K {
			maxWindows -= A[left]
			left++
		}
	}
	return length - left
}
