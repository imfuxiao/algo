package string

import (
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	var (
		num1Len, num2Len = len(num1), len(num2)
		carry            = 0
		ans              strings.Builder
	)

	if num1Len == 0 || num2Len == 0 {
		return ""
	}

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// 已num2为乘数
	lastValue := "0"
	for i := num2Len - 1; i >= 0; i-- {

		n2 := int(num2[i:][0] - '0')

		for j := num1Len - 1; j >= 0; j-- {
			n1 := int(num1[j:][0] - '0')
			num := n1*n2 + carry
			ans.WriteString(strconv.Itoa(num % 10))
			carry = num / 10
		}

		if carry != 0 {
			ans.WriteString(strconv.Itoa(carry))
			carry = 0
		}

		multValue := recover(ans.String())
		for j := num2Len - 1 - i; j > 0; j-- {
			multValue = multValue + "0"
		}

		lastValue = addStr(multValue, lastValue)
		ans.Reset()
	}

	return lastValue
}

// 反转
func recover(str string) string {
	var ans strings.Builder
	length := len(str)
	for i := length - 1; i >= 0; i-- {
		ans.WriteByte(str[i])
	}
	return ans.String()
}

func addStr(num1, num2 string) string {
	var (
		num1Len, num2Len     = len(num1), len(num2)
		num1Index, num2Index = 0, 0
		carry                = 0
		ans                  strings.Builder
	)

	// 相同位数相加
	for {
		if num1Index >= num1Len || num2Index >= num2Len {
			break
		}

		n1 := int(num1[num1Len-1-num1Index:][0] - '0')
		n2 := int(num2[num2Len-1-num2Index:][0] - '0')

		num := n1 + n2 + carry
		ans.WriteString(strconv.Itoa(num % 10))
		carry = num / 10

		num1Index++
		num2Index++
	}

	// 处理位数不一致的情况
	var otherStr string
	if num1Index >= num1Len && num2Index < num2Len {
		otherStr = num2[:num2Len-num2Index]
	} else if num2Index >= num2Len && num1Index < num1Len {
		otherStr = num1[:num1Len-num1Index]
	}

	length := len(otherStr)
	for i := length - 1; i >= 0; i-- {
		num := int(otherStr[i]-'0') + carry
		ans.WriteString(strconv.Itoa(num % 10))
		carry = num / 10
	}

	if carry != 0 {
		ans.WriteString(strconv.Itoa(carry))
	}

	return recover(ans.String())
}

func spiralOrder(matrix [][]int) []int {
	var (
		rows  = len(matrix)
		cols  = len(matrix[0])
		visit = make([][]bool, rows)
		ans   = make([]int, 0, rows*cols)
		f     func(row, col int)
	)

	// init visit
	for i := 0; i < rows; i++ {
		visit[i] = make([]bool, cols)
	}

	f = func(row, col int) {
		if row >= rows || row < 0 || col >= cols || col < 0 {
			return
		}
		if !visit[row][col] {
			ans = append(ans, matrix[row][col])
			visit[row][col] = true

			if row == 0 || visit[row-1][col] {
				f(row, col+1)
			}
			f(row+1, col)
			f(row, col-1)
			f(row-1, col)
		}
	}

	f(0, 0)
	return ans
}

func generateMatrix(n int) [][]int {
	var (
		ans = make([][]int, n)
		f   func(row, col, num int)
	)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	f = func(row, col, num int) {
		if row < 0 || row >= n || col < 0 || col >= n {
			return
		}
		if ans[row][col] == 0 {
			ans[row][col] = num
			if row == 0 || ans[row-1][col] != 0 {
				f(row, col+1, num+1)
			}
			f(row+1, col, num+1)
			f(row, col-1, num+1)
			f(row-1, col, num+1)
		}
	}

	f(0, 0, 1)

	return ans
}
