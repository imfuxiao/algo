package array

import (
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	sorts(nums)
	if nums[0] == 0 {
		return "0"
	}
	var sb strings.Builder
	for i := range nums {
		sb.WriteString(strconv.Itoa(nums[i]))
	}
	return sb.String()
}

func sorts(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	p := position(nums)
	sorts(nums[:p])
	sorts(nums[p+1:])
}

func position(nums []int) int {
	var (
		index = 0
		value = nums[len(nums)-1]
	)
	for i := 0; i < len(nums)-1; i++ {
		if compare(nums[i], value) {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}

	nums[index], nums[len(nums)-1] = nums[len(nums)-1], nums[index]
	return index
}

func compare(num1, num2 int) bool {
	x, y := 10, 10
	for x <= num1 {
		x *= 10
	}
	for y <= num2 {
		y *= 10
	}
	return num1*y+num2 > num2*x+num1
}
