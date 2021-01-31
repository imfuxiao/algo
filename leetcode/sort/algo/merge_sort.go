package algo

// 归并排序
// 平均时间复杂度 O(nlogn)
// 平均空间复杂度 O(n)
func MergeSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	newNums := merge(nums)
	for i := range nums {
		nums[i] = newNums[i]
	}
}

func merge(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	mid := length / 2
	first := merge(nums[:mid])
	second := merge(nums[mid:])

	return func(num1, num2 []int) []int {
		indexNum1, num1Length, indexNum2, num2Length := 0, len(num1), 0, len(num2)
		newNums := make([]int, 0, num1Length+num2Length)
		for indexNum1 < num1Length && indexNum2 < num2Length {
			if num1[indexNum1] < num2[indexNum2] {
				newNums = append(newNums, num1[indexNum1])
				indexNum1++
			} else if num1[indexNum1] > num2[indexNum2] {
				newNums = append(newNums, num2[indexNum2])
				indexNum2++
			} else {
				newNums = append(newNums, num2[indexNum2], num1[num1Length])
				indexNum1++
				indexNum2++
			}
		}

		if indexNum1 == num1Length && indexNum2 < num2Length {
			newNums = append(newNums, num2[indexNum2:]...)
		} else if indexNum2 == num2Length && indexNum1 < num1Length {
			newNums = append(newNums, num1[indexNum1:]...)
		}
		return newNums
	}(first, second)
}
