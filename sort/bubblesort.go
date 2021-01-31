package sort

// 错误的冒泡
func BubbleSort2(n []int) {
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n)-i; j++ {
			if n[i] > n[j] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
}

func BubbleSort(n []int) {
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n)-i-1; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}
}

func BubbleOptimizeSort(n []int) {
	for i, l := 0, len(n); i < l; i++ {
		isBreak := false
		for j := 0; j < l-i-1; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
				isBreak = true
			}
		}
		if !isBreak {
			break
		}
	}
}
