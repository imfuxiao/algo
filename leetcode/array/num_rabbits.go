package array

func numRabbits(answers []int) int {
	var (
		length = len(answers)
		set    = make(map[int]int, length)
		ans    = 0
	)
	for i := range answers {
		set[answers[i]+1]++
	}
	for key, value := range set {
		if value%key == 0 {
			ans += key * (value/key)
		} else {
			ans += key * (value/key + 1)
		}
	}
	return ans
}
