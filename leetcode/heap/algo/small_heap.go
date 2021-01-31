package algo

// 小顶堆: 除叶子节点外, 任何节点的子节点都比自己大
type SmallHeap struct {
	values []int
	length int
	cap    int
}

func NewSmallHeap(cap int) *SmallHeap {
	return &SmallHeap{
		values: make([]int, cap),
		length: 0,
		cap:    cap,
	}
}

func (h *SmallHeap) heapfiy() {
	for index := (h.length - 1) >> 1; index >= 0; index-- {
		leftIndex, rightIndex := index<<1+1, index<<1+2
		if leftIndex < h.length && h.values[leftIndex] < h.values[index] {
			h.values[index], h.values[leftIndex] = h.values[leftIndex], h.values[index]
		}
		if rightIndex < h.length && h.values[rightIndex] < h.values[index] {
			h.values[index], h.values[rightIndex] = h.values[rightIndex], h.values[index]
		}
	}
}

func (h *SmallHeap) Insert(value int) bool {
	if h.length >= h.cap {
		return false
	}
	h.values[h.length] = value
	h.length++
	h.heapfiy()
	return true
}

func (h *SmallHeap) Delete() (int, bool) {
	if h.length <= 0 {
		return 0, false
	}
	value := h.values[0]
	h.length--
	h.values[0] = h.values[h.length]
	h.heapfiy()
	return value, true
}
