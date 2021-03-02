package array

func nextPermutation(nums []int) {
	length := len(nums)
	index, prevIndex := length-2, length-1

	for index >= 0 && nums[index] >= nums[index+1] {
		index--
	}
	if index >= 0 {
		for prevIndex >= 0 && nums[index] >= nums[prevIndex] {
			prevIndex--
		}
		nums[index], nums[prevIndex] = nums[prevIndex], nums[index]
	}
	reverse(nums[index+1:])
}

func reverse(a []int) {
	length := len(a)
	for i := length/2 - 1; i >= 0; i-- {
		a[i], a[length-i-1] = a[length-i-1], a[i]
	}
}
