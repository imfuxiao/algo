package read

import "math"

func longestCommonPrefix(strs []string) string {
	var (
		strCount = len(strs)
		ans      = ""
	)
	if strCount <= 0 {
		return ans
	}

	min := func(n1, n2 int) int {
		if n1 < n2 {
			return n1
		}
		return n2
	}

	var maxCharIndex = math.MaxInt64
	for i := range strs {
		maxCharIndex = min(maxCharIndex, len(strs[i]))
	}

	if maxCharIndex == 0 {
		return ans
	}

	for charIndex := 0; charIndex < maxCharIndex; charIndex++ {
		char := strs[0][charIndex]
		charStatus := true
		for i := 0; i < strCount; i++ {
			if strs[i][charIndex] != char {
				charStatus = false
			}
		}

		if !charStatus {
			return strs[0][:charIndex]
		}
	}
	return ans
}

func longestCommonPrefix2(strs []string) string {
	var (
		strCount = len(strs)
	)
	if strCount == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < strCount; i++ {
		str := strs[i]

		prefix = func(s1, s2 string) string {
			index, len1, len2 := 0, len(s1), len(s2)
			for {
				if index >= len1 || index >= len2 {
					break
				}
				if s1[index] == s2[index] {
					index++
				} else {
					break
				}
			}
			return s1[:index]
		}(prefix, str)

		if prefix == "" {
			break
		}
	}
	return prefix
}
