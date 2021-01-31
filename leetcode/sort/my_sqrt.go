// 69. x 的平方根
// 实现 int sqrt(int x) 函数。
//
// 计算并返回x的平方根，其中x是非负整数。
//
// 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
//
// 示例 1:
//
// 输入: 4
// 输出: 2
//
// 示例 2:
//
// 输入: 8
// 输出: 2
// 说明: 8 的平方根是 2.82842...,
//     由于返回类型是整数，小数部分将被舍去。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/sqrtx
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package sort

import "math"

// 二分查找
func mySqrt1(x int) int {
	max, min, num := x, 0, -1
	for min <= max {
		mid := (max + min) >> 1 // (x+1)>>1 = (x+1)/2 = (x-1) / 2 - 1
		sqrt := mid * mid

		if sqrt > x {
			max = mid - 1
		} else {
			min = mid + 1
			num = mid
		}
	}
	return num
}

// 数学计算法: 根号x = x^(1/2) = e^ln(x)^/2 = e^(0.5*ln(x))
func mySqrt2(x int) int {
	sqrtResult := int(math.Exp(0.5 * math.Log(float64(x))))
	if (sqrtResult+1)*(sqrtResult+1) <= x {
		sqrtResult++
	}
	return sqrtResult
}
