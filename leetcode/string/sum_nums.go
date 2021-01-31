package string

func sumNums(n int) int {
	var (
		sum func(int) bool
		ana int
	)
	sum = func(num int) bool {
		ana += num
		return num > 0 && sum(num-1)
	}
	sum(n)
	return ana
}
