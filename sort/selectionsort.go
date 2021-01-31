package sort

func SelectionSort(n []int) {
	for i, l := 0, len(n); i < l; i++ {
		value, index := n[i], i
		for j := i; j < l; j++ {
			if value > n[j] {
				value = n[j]
				index = j
			}
		}
		n[i], n[index] = value, n[i]
	}
}
