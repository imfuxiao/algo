package brute_force

// BF 暴力匹配算法
// primary: 主串
// model: 模式串
func BruteForce(primary, model string) bool {
	primaryLength, modelLength := len(primary), len(model)
	if modelLength > primaryLength {
		return false
	}

	check := func(firstStr, secondStr string) bool {
		length := len(firstStr)
		result := true
		for i := 0; i < length; i++ {
			if firstStr[i] != secondStr[i] {
				result = false
				break
			}
		}
		return result
	}

	for primaryStart := 0; primaryStart <= primaryLength-modelLength; primaryStart++ {
		if check(primary[primaryStart:primaryStart+modelLength], model) {
			return true
		}
	}

	return false
}
