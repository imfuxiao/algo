package read

// 滑动窗口
func LengthOfLongestSubstring(s string) int {
	var (
		length     = len(s)
		rightIndex = -1                              // 右侧指针
		status     = make(map[byte]struct{}, length) // 保存已匹配未重复的字符
		answer     = 0                               // 最长子串长度
	)

	// 判断key是否存在
	isExists := func(key byte) bool {
		_, exists := status[key]
		return exists
	}

	max := func(n1, n2 int) int {
		if n1 > n2 {
			return n1
		}
		return n2
	}

	for leftIndex := 0; leftIndex < length; leftIndex++ {
		if leftIndex != 0 {
			// 删除未重复子串的第一个字符
			delete(status, s[leftIndex-1])
		}

		// 右指针向右移动, 通过status判断是否遇到重复字符, 当遇到重复字符, 则退出循环
		for rightIndex+1 < length && !isExists(s[rightIndex+1]) {
			status[s[rightIndex+1]] = struct{}{}
			rightIndex++
		}

		answer = max(answer, rightIndex-leftIndex+1)
	}

	return answer
}
