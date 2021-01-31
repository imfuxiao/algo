package binarysearch

func BinarySearch1(n []int, num int) int {
	index := -1

	if len(n) <= 0 {
		return index
	}

	length := len(n)
	// low: 最低位下标, high: 最高位下标 mid: 二分查找中间点下标
	low, high, mid := 0, length-1, length/2

	// (hight + low) >> 1 推导过程:  (hight - low) / 2 + low 等于 (hight + low) / 2 等于 (hight + low) >> 1
	for {
		switch {
		case n[mid] == num:
			return mid
		case n[mid] > num:
			high = mid - 1
			mid = (high + low) >> 1
		case n[mid] < num:
			low = mid + 1
			mid = (high + low) >> 1
		}

		// 注意for循环退出条件
		if low > high {
			return index
		}
	}
}

// 查找第一个值等于给定值的元素, 如:[1, 3, 4, 8, 8, 8, 8, 8, 11, 18], 找 8, 返回下标3
func BinarySearch2(n []int, num int) int {
	index := -1

	if len(n) <= 0 {
		return index
	}

	low, high, mid := 0, len(n)-1, len(n)/2

	for {

		switch {
		case n[mid] < num:
			low = mid + 1
		case n[mid] > num:
			high = mid - 1
		default:
			if mid == 0 || n[mid-1] != num {
				index = mid
				return index
			} else {
				high = mid - 1
			}
		}

		mid = (low + high) >> 1

		if low > high {
			break
		}

	}

	return index
}

// 查找最后一个值等于给定值的元素, 如:[1, 3, 4, 8, 8, 8, 8, 8, 11, 18], 找 8, 返回下标7
func BinarySearch3(n []int, num int) int {
	index := -1

	if len(n) <= 0 {
		return index
	}

	low, high, mid := 0, len(n)-1, len(n)/2

	for {
		switch {
		case n[mid] < num:
			low = mid + 1
		case n[mid] > num:
			high = mid - 1
		default:
			if mid == len(n)-1 || n[mid+1] != num {
				return mid
			} else {
				low = mid + 1
			}
		}

		mid = (low + high) >> 1
		if low > high {
			break
		}
	}

	return index
}

// 查找第一个大于等于给定值的元素
func BinarySearch4(n []int, num int) int {
	index := -1

	low, high, mid := 0, len(n)-1, len(n)/2

	for {
		if n[mid] >= num {
			if mid == 0 || n[mid-1] < num {
				index = mid
				break
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}

		mid = (low + high) >> 1

		if low > high {
			break
		}
	}

	return index
}

// 查找最后一个小于给定值的元素
func BinarySearch5(n []int, num int) int {
	index := -1
	if len(n) <= 0 {
		return index
	}

	low, high, mid := 0, len(n)-1, len(n)/2

	for {

		if n[mid] >= num {
			high = mid - 1
		} else {
			if mid == len(n)-1 || n[mid+1] >= num {
				index = mid
				break
			} else {
				low = mid + 1
			}
		}

		mid = (low + high) / 2

		if low > high {
			break
		}

	}

	return index
}
