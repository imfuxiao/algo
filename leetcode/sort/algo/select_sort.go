package algo

// 选择排序
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
func SelectSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	for i := 0; i < length; i++ {
		min, index := nums[i], i
		for j := i + 1; j < length; j++ {
			if nums[j] < min {
				min, index = nums[j], j
			}
		}
		nums[i], nums[index] = nums[index], nums[i]
	}
}
