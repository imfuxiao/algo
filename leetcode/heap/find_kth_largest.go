// 215. 数组中的第K个最大元素
// 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 示例 1:
//
// 输入: [3,2,1,5,6,4] 和 k = 2
// 输出: 5
// 示例 2:
//
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
// 输出: 4
// 说明:
//
// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package heap

func FindKthLargest1(nums []int, k int) int {
	var (
		heap      = make([]int, 0, k) // 小顶堆
		heapCount = 0                 // 堆元素大小
		heapfiy   = func() {
			for index := (heapCount - 1) >> 1; index >= 0; index-- {
				leftIndex, rightIndex := index<<1+1, index<<1+2
				if leftIndex < heapCount && heap[leftIndex] < heap[index] {
					heap[leftIndex], heap[index] = heap[index], heap[leftIndex]
				}
				if rightIndex < heapCount && heap[rightIndex] < heap[index] {
					heap[rightIndex], heap[index] = heap[index], heap[rightIndex]
				}
			}
		}
		Insert = func(value int) {
			heapCount++
			heap = append(heap, value)
			heapfiy()
		}
		Delete = func() int {
			heapCount--
			value := heap[0]
			heap = heap[1:]
			heapfiy()
			return value
		}
	)

	for i := range nums {
		if heapCount < k {
			Insert(nums[i])
			continue
		}
		if heapCount == k && nums[i] > heap[0] {
			Delete()
			Insert(nums[i])
		}
	}
	return heap[0]
}

// 快速排序
func FindKthLargest2(nums []int, k int) int {
	var (
		quickSort func(nums []int)
		position  func(nums []int, length int) int
	)

	quickSort = func(nums []int) {
		length := len(nums)
		if length <= 1 {
			return
		}
		p := position(nums, length)
		quickSort(nums[:p])
		quickSort(nums[p+1:])
	}

	position = func(nums []int, length int) int {
		index, value := 0, nums[length-1]
		for i := 0; i < length-1; i++ {
			if nums[i] > value {
				nums[i], nums[index] = nums[index], nums[i]
				index++
			}
		}
		nums[index], nums[length-1] = nums[length-1], nums[index]
		return index
	}

	quickSort(nums)
	return nums[k-1]
}
