package read

import (
	"strings"
)

func DecodeString(s string) string {
	var (
		toString func() string
		index    = 0
	)

	toString = func() string {
		if index >= len(s) || s[index] == ']' {
			return ""
		}
		res := ""
		if s[index] >= '0' && s[index] <= '9' {
			repeatCount, numIndex := toDigital(s, index)
			index = numIndex + 1
			repeatStr := toString()
			res += strings.Repeat(repeatStr, repeatCount)
		} else if s[index] <= 'z' && s[index] >= 'a' || s[index] >= 'A' && s[index] <= 'Z' {
			res += string(s[index])
		}
		index++
		return res + toString()
	}
	return toString()
}

func toDigital(s string, index int) (int, int) {
	num := 0
	for i := index; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
			index = i
		} else {
			break
		}
	}
	return num, index
}
