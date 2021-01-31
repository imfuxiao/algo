package sort

func QuickSort(n []int) {
	if len(n) <= 1 {
		return
	}
	p := position(n)
	QuickSort(n[:p])
	QuickSort(n[p+1:])
}

func position(n []int) int {
	maxIndex := len(n) - 1
	index := 0
	for i := 0; i < maxIndex; i++ {
		if n[i] < n[maxIndex] {
			n[index], n[i] = n[i], n[index]
			index++
		}
	}
	n[index], n[maxIndex] = n[maxIndex], n[index]
	return index
}
