package array

func minSwapsCouples(row []int) int {
	ans := 0
	length := len(row)
	graphs := make([][]int, length/2)
	visited := make([]bool, length/2)
	for i := 0; i < length; i += 2 {
		left, right := row[i]/2, row[i+1]/2
		if left != right {
			graphs[left] = append(graphs[left], right)
			graphs[right] = append(graphs[right], left)
		}
	}

	queue := make([]int, 0, length)
	enqueue := func(value int) {
		queue = append(queue, value)
	}
	dequeue := func() int {
		value := queue[0]
		queue = queue[1:]
		return value
	}

	for row := 0; row < length/2; row++ {
		if !visited[row] {
			visited[row] = true

			enqueue(row)
			count := 0

			for len(queue) > 0 {
				r := dequeue()
				count++
				for j := range graphs[r] {
					if !visited[graphs[r][j]] {
						visited[graphs[r][j]] = true
						enqueue(graphs[r][j])
					}
				}
			}

			ans += count - 1
		}
	}

	return ans

}
