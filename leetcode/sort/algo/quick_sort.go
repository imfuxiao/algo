package algo

// 快速排序
// 时间复杂度: O(nlogn)
// 空间复杂度: O(1)
func QuickSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	index := position(nums, length)
	QuickSort(nums[:index])
	QuickSort(nums[index+1:])
}

func position(nums []int, length int) int {
	num := nums[length-1]
	index := 0
	for i := 0; i < length-1; i++ {
		if nums[i] < num {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
	nums[index], nums[length-1] = nums[length-1], nums[index]
	return index
}
