package sort

func InsertionSort(n []int) {
	for i, l := 1, len(n); i < l; i++ {
		value, j := n[i], i-1
		for ; j >= 0; j-- {
			if n[j] > value {
				n[j+1] = n[j]
			} else {
				break
			}
		}
		n[j+1] = value
	}
}
