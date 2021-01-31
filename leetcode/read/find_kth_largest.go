package read

import (
	"errors"
)

func findKthLargest(nums []int, k int) int {
	length := len(nums)
	miniHeap := NewHeap(k)
	for i := 0; i < length; i++ {
		if i < k {
			miniHeap.Push(nums[i])
			continue
		}
		if nums[i] > miniHeap.Top() {
			miniHeap.Pop()
			miniHeap.Push(nums[i])
		}
	}
	value, _ := miniHeap.Pop()
	return value
}

// 小顶堆
type heap struct {
	elements    []int
	cap, length int
}

func NewHeap(cap int) *heap {
	return &heap{
		elements: make([]int, cap),
		cap:      cap,
		length:   0,
	}
}

// 堆化
func (h *heap) heapfiy() {
	mid := h.length/2 - 1
	for i := mid; i >= 0; i-- {
		leftChildIndex, rightChildIndex := i<<1+1, i<<1+2
		if leftChildIndex < h.length && h.elements[leftChildIndex] < h.elements[i] {
			h.elements[i], h.elements[leftChildIndex] = h.elements[leftChildIndex], h.elements[i]
		}
		if rightChildIndex < h.length && h.elements[rightChildIndex] < h.elements[i] {
			h.elements[i], h.elements[rightChildIndex] = h.elements[rightChildIndex], h.elements[i]
		}
	}
}
func (h *heap) Top() int {
	return h.elements[0]
}

func (h *heap) Push(element int) bool {
	if h.length == h.cap {
		return false
	}

	h.elements[h.length] = element
	h.length++
	h.heapfiy()
	return true
}

func (h *heap) Pop() (int, error) {
	if h.length <= 0 {
		return -1, errors.New("no elements")
	}
	value := h.elements[0]
	h.length--
	if h.length > 0 {
		h.elements[0] = h.elements[h.length]
		h.heapfiy()
	}
	return value, nil
}
