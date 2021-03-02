package array

import "strings"

func sumSubarrayMins(A []int) int {
	var (
		n     = len(A)
		stack = make([]int, n) // 单调递增栈
		top   = -1
		count = 0
		mod   = 1000000007
	)

	for i := 0; i < n; i++ {
		if top != -1 && A[stack[top]] >= A[i] {
			index := stack[top]
			top--
			leftIndex := -1
			if top != -1 {
				leftIndex = stack[top]
			}
			rightIndex := i
			count += (index - leftIndex) * (rightIndex - index) * A[index]
			count %= mod
		}
		top++
		stack[top] = i
	}

	for top != -1 {
		index := stack[top]
		top--
		leftIndex := -1
		if top != -1 {
			leftIndex = stack[top]
		}
		rightIndex := n
		count += (index - leftIndex) * (rightIndex - index) * A[index]
		count %= mod
	}

	return count

}

const (
	one  byte = '1'
	zero byte = '0'
)

func addBinary(a string, b string) string {
	carry := zero
	lengthAIndex, lengthBIndex := len(a)-1, len(b)-1

	numsIndex := max(lengthAIndex, lengthBIndex)
	nums := make([]byte, numsIndex+1)
	for lengthAIndex >= 0 && lengthBIndex >= 0 {
		num, c1 := add(a[lengthAIndex], b[lengthBIndex])
		num, c2 := add(num, carry)

		nums[numsIndex] = num
		carry, _ = add(c1, c2)

		numsIndex--
		lengthAIndex--
		lengthBIndex--
	}

	for lengthAIndex >= 0 {
		num, c1 := add(a[lengthAIndex], carry)

		nums[numsIndex] = num
		carry = c1

		numsIndex--
		lengthAIndex--
	}

	for lengthBIndex >= 0 {
		num, c1 := add(b[lengthBIndex], carry)

		nums[numsIndex] = num
		carry = c1

		numsIndex--
		lengthBIndex--
	}

	var ans strings.Builder
	if carry == one {
		ans.WriteByte(one)
	}

	for _, num := range nums {
		ans.WriteByte(num)
	}

	return ans.String()
}

func add(a, b byte) (byte, byte) {
	if a == one && b == one {
		return zero, one
	}
	if (a == one && b == zero) || (b == one && a == zero) {
		return one, zero
	}
	return zero, zero
}
