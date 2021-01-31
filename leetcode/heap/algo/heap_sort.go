package algo

func HeapSort(nums []int) {
	length := len(nums)
	// 1. 建堆
	heapfiy(nums, 0, length-1)

	// 排序
	for i := length - 1; i > 0; i-- {
		// 将大元素放在后面
		nums[0], nums[i] = nums[i], nums[0]
		// 建堆
		heapfiy(nums, 0, i-1)
	}
}

// 大顶堆
func heapfiy(nums []int, startIndex, endIndex int) {
	if startIndex == endIndex {
		return
	}
	for index := (endIndex - startIndex) / 2; index >= 0; index-- {
		leftIndex, rightIndex := index>>1+1, index>>1+2
		if leftIndex <= endIndex && nums[leftIndex] > nums[index] {
			nums[index], nums[leftIndex] = nums[leftIndex], nums[index]
		}
		if rightIndex <= endIndex && nums[rightIndex] > nums[index] {
			nums[index], nums[rightIndex] = nums[rightIndex], nums[index]
		}
	}
}
