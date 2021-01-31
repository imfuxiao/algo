package read

func topKFrequent(nums []int, k int) []int {

	var (
		length = len(nums)
		status = make(map[int]int, length)
	)
	type miniHeap struct {
		nums        []int
		cap, length int
	}
	heap := &miniHeap{
		nums:   make([]int, k),
		cap:    k,
		length: 0,
	}
	heapfiy := func() {
		mid := heap.length/2 - 1
		for i := mid; i >= 0; i-- {
			leftChildIndex, rightChildIndex := i<<1+1, i<<1+2
			if leftChildIndex < heap.length && status[heap.nums[leftChildIndex]] < status[heap.nums[i]] {
				heap.nums[i], heap.nums[leftChildIndex] = heap.nums[leftChildIndex], heap.nums[i]
			}
			if rightChildIndex < heap.length && status[heap.nums[rightChildIndex]] < status[heap.nums[i]] {
				heap.nums[i], heap.nums[rightChildIndex] = heap.nums[rightChildIndex], heap.nums[i]
			}
		}
	}
	push := func(num int) {
		if heap.length >= heap.cap {
			return
		}
		heap.nums[heap.length] = num
		heap.length++
		heapfiy()
	}
	pop := func() int {
		if heap.length <= 0 {
			return -1
		}

		value := heap.nums[0]
		if heap.length > 0 {
			heap.nums[0] = heap.nums[heap.length-1]
		}
		heap.length--
		heapfiy()
		return value
	}
	for i := range nums {
		status[nums[i]]++
	}

	for key, value := range status {
		if heap.length < k {
			push(key)
			continue
		}
		if value > status[heap.nums[0]] {
			pop()
			push(key)
		}
	}
	return heap.nums
}
