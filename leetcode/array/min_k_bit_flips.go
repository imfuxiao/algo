package array

func minKBitFlips(A []int, K int) int {
	length := len(A)
	ans := 0
	for right := 0; right < length; right++ {
		if A[right] == 1 {
			continue
		}
		if right+K > length {
			return -1
		}
		for step := 0; step < K; step++ {
			A[right+step] ^= 1
		}
		ans++
	}
	return ans
}
