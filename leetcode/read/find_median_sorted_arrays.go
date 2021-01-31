package read

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	firstNumLength, firstIndex, secondNumLength, secondIndex := len(nums1), 0, len(nums2), 0
	allNumLength := firstNumLength + secondNumLength
	allNum := make([]int, 0, allNumLength)

	// 数组合并
	for {
		if firstIndex >= firstNumLength {
			if secondIndex < secondNumLength {
				allNum = append(allNum, nums2[secondIndex:]...)
			}
			break
		}
		if secondIndex >= secondNumLength {
			if firstIndex < firstNumLength {
				allNum = append(allNum, nums1[firstIndex:]...)
			}
			break
		}
		if nums1[firstIndex] < nums2[secondIndex] {
			allNum = append(allNum, nums1[firstIndex])
			firstIndex++
		} else if nums1[firstIndex] > nums2[secondIndex] {
			allNum = append(allNum, nums2[secondIndex])
			secondIndex++
		} else {
			allNum = append(allNum, nums1[firstIndex], nums2[secondIndex])
			firstIndex++
			secondIndex++
		}
	}

	mid, midNum := allNumLength/2, 0.0
	if allNumLength%2 == 0 { // 偶数
		midNum = float64(allNum[mid]+allNum[mid-1]) / 2
	} else {
		midNum = float64(allNum[mid])
	}

	return midNum
}
