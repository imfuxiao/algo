package read

import (
	"sort"
)

func FindBestValue(arr []int, target int) int {
	// 数组先排序
	sort.Ints(arr)
	length := len(arr)

	// 前缀和
	preSum := make([]int, length+1)
	for i := 1; i <= length; i++ {
		preSum[i] = arr[i-1] + preSum[i-1]
	}

	// 二分查找
	min, max, ans := 0, arr[length-1], -1
	for min <= max {
		mid := (max + min) >> 1

		// mid在arr中不存在, 则返回 最接近mid的值的索引
		preIndex := sort.SearchInts(arr, mid)
		if preIndex < 0 {
			preIndex = abs(preIndex) - 1
		}

		sum := preSum[preIndex] + (length-preIndex)*mid
		if sum <= target {
			ans = mid
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	choseSmall := check(arr, ans)
	choseBig := check(arr, ans+1)
	if abs(choseSmall-target) > abs(choseBig-target) {
		ans++
	}

	return ans
}

func check(arr []int, x int) int {
	ret := 0
	for _, num := range arr {
		ret += minFunc(num, x)
	}
	return ret
}

func minFunc(n1, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}
