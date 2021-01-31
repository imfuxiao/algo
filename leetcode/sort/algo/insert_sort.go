package algo

// 插入排序
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
func InsertSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	// 未排序区域, 从1开始
	for i := 1; i < length; i++ {
		value, index := nums[i], i

		// 已排序区域, 初始只有一个0元素
		// 必须从高位开始
		for j := i - 1; j >= 0; j-- {
			if nums[j] <= value {
				break
			}
			nums[j+1] = nums[j]
			index = j
		}
		nums[index] = value
	}
}
