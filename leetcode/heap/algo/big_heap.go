package algo

// 大顶堆: 除叶子节点外, 任何节点都比自己小
type BigStack struct {
	values []int
	cap    int
	length int
}

func NewBigStack(cap int) *BigStack {
	return &BigStack{
		values: make([]int, cap),
		cap:    cap,
		length: 0,
	}
}

// 堆化
func (s *BigStack) heapfiy() {
	// 堆化: 从 (n-1)/2 非叶子节点索引处. 向上堆化
	for index := (s.length - 1) >> 1; index >= 0; index-- {
		leftIndex := index<<1 + 1
		rightIndex := index<<1 + 2
		if leftIndex < s.length && s.values[leftIndex] > s.values[index] {
			s.values[index], s.values[leftIndex] = s.values[leftIndex], s.values[index]
		}
		if rightIndex < s.length && s.values[rightIndex] > s.values[index] {
			s.values[index], s.values[rightIndex] = s.values[rightIndex], s.values[index]
		}
	}
}

func (s *BigStack) Insert(value int) bool {
	if s.length == s.cap {
		return false
	}
	s.values[s.length] = value
	s.length++
	s.heapfiy()
	return true
}

func (s *BigStack) Delete() (int, bool) {
	if s.length == 0 {
		return 0, false
	}
	value := s.values[0]
	s.length--
	s.values[0] = s.values[s.length]
	s.values = s.values[:s.length]
	s.heapfiy()
	return value, true
}
