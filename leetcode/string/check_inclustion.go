package string

func checkInclusion(s1 string, s2 string) bool {
	lenS1, lenS2 := len(s1), len(s2)
	if lenS1 > lenS2 {
		return false
	}

	s1Array := [26]int{}
	for i := range s1 {
		s1Array[int(s1[i]-'a')]++
	}

	// 滑动窗口
	left, right := 0, 0
	existsArray := [26]int{}
	for right < lenS2 {
		rightIndex := int(s2[right] - 'a')
		if s1Array[rightIndex] > 0 {
			existsArray[rightIndex]++
		}
		right++

		if existsArray == s1Array {
			return true
		}

		if right-left == lenS1 {
			leftIndex := int(s2[left] - 'a')
			if s1Array[leftIndex] > 0 {
				existsArray[leftIndex]--
			}
			left++
		}
	}

	return false
}

func minWindow(s string, t string) string {
	ans := ""
	lenS, lenT := len(s), len(t)
	if lenT > lenS {
		return ans
	}

	sArray, tArray := [72]int{}, [72]int{}
	for i := range t {
		tArray[t[i]-'A']++
	}

	check := func() bool {
		for i := range t {
			index := t[i] - 'A'
			if sArray[index] < tArray[index] {
				return false
			}
		}
		return true

	}

	for l, r := 0, 0; r < lenS; r++ {
		rightIndex, leftIndex := s[r]-'A', s[l]-'A'
		if tArray[rightIndex] > 0 {
			sArray[rightIndex]++
		}

		if tArray[leftIndex] == 0 {
			l++
		}

		for r-l+1 >= lenT && check() {
			if ans == "" || (r-l+1 < len(ans)) {
				ans = s[l : r+1]
			}
			leftIndex := s[l] - 'A'
			if sArray[leftIndex] > 0 {
				sArray[leftIndex]--
			}
			l++
		}
	}

	return ans
}

func findAnagrams(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	ans := make([]int, 0, sLen)

	if pLen > sLen {
		return ans
	}

	sArray, pArray := [72]int{}, [72]int{}
	for i := range p {
		pArray[p[i]-'A']++
	}

	for left, right := 0, 0; right < sLen; right++ {
		rightIndex, leftIndex := s[right]-'A', s[left]-'A'

		if pArray[rightIndex] > 0 {
			sArray[rightIndex]++
		}

		if sArray == pArray {
			ans = append(ans, left)
		}

		if right-left+1 >= pLen {
			if pArray[leftIndex] > 0 {
				sArray[leftIndex]--
			}
			left++
		}
	}

	return ans
}
