package read

import (
	"strconv"
)

func AddStrings(num1 string, num2 string) string {
	var (
		add = 0 // 进位值
		ans = ""
	)

	for firstIndex, secondIndex := len(num1)-1, len(num2)-1; firstIndex >= 0 || secondIndex >= 0 || add != 0; firstIndex, secondIndex = firstIndex-1, secondIndex-1 {
		var first, second int // 因初始值为0, 所以位数不齐时, 直接用初始值0填补

		if firstIndex >= 0 {
			first = int(num1[firstIndex] - '0')
		}

		if secondIndex >= 0 {
			second = int(num2[secondIndex] - '0')
		}

		sum := first + second + add
		add = sum / 10
		ans = strconv.Itoa(sum%10) + ans
	}

	return ans
}
