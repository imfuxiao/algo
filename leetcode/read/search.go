package read

func Search(nums []int, target int) int {
	var (
		length = len(nums)
		mid    = length >> 1
	)

	if nums[mid] == target {
		return mid
	}

	// 二分查找
	// add是增量, 右侧区域二分查找结束后需要给index加上右侧minIndex值
	binarySearch := func(nums []int, target, add int) int {
		var (
			length        = len(nums)
			min, max, mid = 0, length - 1, length >> 1
		)
		for {
			if min > max {
				return -1
			}
			if target == nums[mid] {
				return mid + add
			}
			if target < nums[mid] {
				max = mid - 1
			} else {
				min = mid + 1
			}
			mid = (max + min) >> 1
		}
	}

	// 寻找旋转点
	leftMax, rightMin := mid, mid
	if nums[0] <= nums[leftMax] { // 左侧有序
		for {
			if leftMax+1 < length && nums[leftMax+1] > nums[leftMax] {
				leftMax++
			} else {
				break
			}
		}
		rightMin = leftMax + 1
	} else { // 右侧有序
		for {
			if rightMin-1 > 0 && nums[rightMin-1] < nums[rightMin] {
				rightMin--
			} else {
				break
			}
		}
		leftMax = rightMin - 1
	}

	if target >= nums[0] && target <= nums[leftMax] {
		return binarySearch(nums[:leftMax+1], target, 0)
	}
	return binarySearch(nums[rightMin:], target, rightMin)
}
