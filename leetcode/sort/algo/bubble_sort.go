package algo

// 冒泡排序
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
func BubbleSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
