package sort

func CountingSort(n []int) {
	if len(n) <= 1 {
		return
	}

	// 1. 求最大值
	max := 0
	for _, v := range n {
		if v > max {
			max = v
		}
	}

	// 计数算法用的桶: 下标为入参数组的值, 值对应入参数组值的个数
	bucket := make([]int, max+1)
	// 填充桶内值
	for i := range n {
		bucket[n[i]] = bucket[n[i]] + 1
	}

	// 依次累加桶内的值
	for i := 1; i <= max; i++ {
		bucket[i] = bucket[i] + bucket[i-1]
	}

	// 排序后的数组
	sortArray := make([]int, len(n))

	// 为排序数组赋值
	for i := len(n) - 1; i >= 0; i-- {
		v := n[i]
		sortArray[bucket[v]-1] = v
		bucket[v] = bucket[v] - 1
	}

	copy(n, sortArray)
}
