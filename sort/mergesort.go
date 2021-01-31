package sort

func MergeSort(n []int) {
	if len(n) <= 1 {
		return
	}
	copy(n, merge(n))
}

func merge(n []int) []int {
	if len(n) <= 1 {
		return n
	}

	middle := len(n) / 2

	m1, m2 := merge(n[:middle]), merge(n[middle:])

	return func(m1, m2 []int) []int {
		r := make([]int, 0, len(m1)+len(m2))
		i, j := 0, 0
		var s []int
		for {
			v1, v2 := m1[i], m2[j]
			switch {
			case v1 > v2:
				j++
				r = append(r, v2)
			case v2 > v1:
				i++
				r = append(r, v1)
			default:
				i++
				j++
				r = append(r, v1, v2)
			}
			if i >= len(m1) {
				s = m2[j:]
				break
			}
			if j >= len(m2) {
				s = m1[i:]
				break
			}
		}
		r = append(r, s...)
		return r
	}(m1, m2)
}
