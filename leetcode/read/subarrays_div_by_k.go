package read

func SubArraysDivByK(A []int, K int) int {
	var (
		preSum, answer = 0, 0
		status         = make([]int, K)
	)
	status[0] = 1
	for i := range A {
		preSum += A[i]
		mod := (preSum%K + K) % K
		answer += status[mod]
		status[mod]++
	}
	return answer
}
