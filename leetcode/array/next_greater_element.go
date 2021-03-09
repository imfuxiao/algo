package array

import (
	"math"
	"sort"
	"strconv"
)

func nextGreaterElement(n int) int {
	var (
		bytes       = []byte(strconv.Itoa(n))
		length      = len(bytes)
		left, right = 0, 0
		stack       = make([]int, 0, length)
	)
	// 从右向左, 寻找第一个降序的index
	for i := length - 1; i >= 0; i-- {
		if len(stack) > 0 && bytes[stack[len(stack)-1]] > bytes[i] {
			left = i
			break
		}
		stack = append(stack, i)
	}

	// 从右向左, 寻找第一个大于bytes[left]的index
	for i := length - 1; i >= 0; i-- {
		if bytes[i] > bytes[left] {
			right = i
			break
		}
	}

	if left < right {
		bytes[left], bytes[right] = bytes[right], bytes[left]

		b := bytes[left+1:]
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
	}

	value := byteToInt(bytes)
	if value == n || value > math.MaxInt32 {
		return -1
	}
	return value
}

func byteToInt(bytes []byte) int {
	value, _ := strconv.Atoi(string(bytes))
	return value
}
