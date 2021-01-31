package recursion

import (
	"fmt"
)

func Step(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return Step(n-1) + Step(n-2)
}

// 全排列问题: 输出 data 数组元素的全排列
func AllPermutation(data []int) {
	r := p(data, len(data))
	fmt.Println("AmN total: %n", len(r))
	for i := range r {
		fmt.Printf("%+v\n", r[i])
	}
}

func p(data []int, k int) [][]int {
	var r [][]int
	// 递归终止条件
	if k == 1 {
		n := make([]int, len(data))
		copy(n, data)
		r = append(r, n)
		return r
	}

	for i := 0; i < k; i++ {
		data[i], data[k-1] = data[k-1], data[i]
		r = append(r, p(data, k-1)...)
		data[i], data[k-1] = data[k-1], data[i]
	}

	return r
}

func Fibonacci(n int) int {
	if n == 1 {
		return 2
	}
	if n == 2 {
		return 4
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
