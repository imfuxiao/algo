package array

var (
	ac  = 2
	ans = 1
	sum = 1
)

func minimumBoxes(n int) int {
	if n == 1 {
		return sum
	}

	if sum >= n {
		return ans
	}

	ans += ac
	sum += ans
	ac++

	minimumBoxes(n)

	if sum == n || ans == n {
		return ans
	}

	for sum > n {
		ac--
		sum -= ans
		ans -= ac
	}

	return ans + (n - sum)
}
