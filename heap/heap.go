package heap

import (
	"errors"
	"fmt"
)

type Heap struct {
	elements     []int // 用户存储元素从下标1开始, 方便计算
	elementTotal int
	elementMax   int
}

func NewHeap(max int) *Heap {
	if max <= 0 {
		max = 16
	}
	elements := make([]int, 0, max+1)
	elements = append(elements, -1)
	return &Heap{
		elements:     elements,
		elementTotal: 0,
		elementMax:   max,
	}
}

func (h *Heap) String() string {
	return fmt.Sprintf("%+v", h.elements[1:])
}

func (h *Heap) Add(element int) error {
	if h.elementTotal == h.elementMax {
		return errors.New("heap overflow")
	}

	h.elementTotal++
	h.elements = append(h.elements, element)

	// 自下向上堆化
	// parent 父节点索引 = index / 2, 即右移一位
	for index, parent := h.elementTotal, h.elementTotal>>1; parent > 0; parent = parent >> 1 {
		if index > 1 && h.elements[index] > h.elements[parent] {
			h.elements[index], h.elements[parent] = h.elements[parent], h.elements[index]
			index = parent
		}
	}

	return nil
}

func (h *Heap) RemoveHeapTop() {
	if h.elementTotal == 0 {
		return
	}
	// 将最后一个元算放到堆顶
	h.elements[1] = h.elements[h.elementTotal]
	h.elements = h.elements[:h.elementTotal]
	h.elementTotal--

	// 自上向下堆化
	for index, maxElementIndex := 1, 1; ; {

		// 左子节点, 右子节点索引 = 当前节点索引 * 2, 当前节点索引 * 2 + 1, 乘2等于左移一位
		leftChildIndex, rightChildIndex := index<<1, (index<<1)+1

		// 找出当前节点与两个左右子节点的最大值的索引
		if leftChildIndex <= h.elementTotal && h.elements[index] < h.elements[leftChildIndex] {
			maxElementIndex = leftChildIndex
		}
		if rightChildIndex <= h.elementTotal && h.elements[maxElementIndex] < h.elements[rightChildIndex] {
			maxElementIndex = rightChildIndex
		}

		// 终止条件: 符合完全二叉树定义, 即当前节点的值大于等于其左右子节点的值
		if maxElementIndex == index {
			return
		}

		h.elements[maxElementIndex], h.elements[index] = h.elements[index], h.elements[maxElementIndex]
		index = maxElementIndex
	}
}

// 堆排序: 1. 建堆, 2. 排序: 排序利用堆顶为最大值的特性
func HeapSort(elements []int) {
	total := len(elements)
	if total <= 1 {
		return
	}

	// 建堆: 自下向上
	// 因为[n/2+1, n]区间为叶子节点, 所以只需要对[1, n/2]区间数据做比较
	// 因为下标从0开始, 而非从1开始, 所以父节点索引为: (n - 1) / 2
	for index := (total - 1) / 2; index >= 0; index-- {
		heapfiy(elements, index, total)
	}

	// 排序: 把0与max交换, 并对[0, max-1]区间重新堆化
	for maxElementIndex := total - 1; maxElementIndex > 0; maxElementIndex-- {
		elements[0], elements[maxElementIndex] = elements[maxElementIndex], elements[0]
		heapfiy(elements[:maxElementIndex], 0, maxElementIndex)
	}
}

// 自上向下堆化: 需要比较节点和两个子节点大小, 拿到最大值节点与节点值进行交换, 并根据最大值节点下标往下一级节点递进
// elements: 代建堆
// index: 起始下标
// maxIndex: 最大下标
func heapfiy(elements []int, index, maxIndex int) {
	for {
		tempIndex := index

		leftChildIndex, rightChildIndex := tempIndex<<1+1, tempIndex<<1+2

		if leftChildIndex < maxIndex && elements[index] < elements[leftChildIndex] {
			tempIndex = leftChildIndex
		}
		if rightChildIndex < maxIndex && elements[tempIndex] < elements[rightChildIndex] {
			tempIndex = rightChildIndex
		}

		if tempIndex == index {
			break
		}

		elements[tempIndex], elements[index] = elements[index], elements[tempIndex]
		index = tempIndex
	}
}
