package backtracking

import "math"

// 莱文斯坦距离
func Lwst1(str1, str2 string) int {
	return lwst1(str1, str2, len(str1), len(str2), 0, 0, 0)
}

func lwst1(str1, str2 string, lengthStr1, lengthStr2, indexStr1, indexStr2, count int) int {
	if indexStr1 >= lengthStr1 || indexStr2 >= lengthStr2 {
		if indexStr1 >= lengthStr1 && indexStr2 < lengthStr2 {
			count = count + len(str2[indexStr2:])
		}
		if indexStr2 >= lengthStr2 && indexStr1 < lengthStr1 {
			count = count + len(str2[indexStr1:])
		}
		return count
	}

	c := math.MaxInt64

	if str1[indexStr1] == str2[indexStr2] {
		c1 := lwst1(str1, str2, lengthStr1, lengthStr2, indexStr1+1, indexStr2+1, count)
		if c1 < c {
			c = c1
		}

	} else {
		// 删除a[i]或者b[j]前添加一个字符
		c1 := lwst1(str1, str2, lengthStr1, lengthStr2, indexStr1, indexStr2+1, count+1)
		c2 := lwst1(str1, str2, lengthStr1, lengthStr2, indexStr1+1, indexStr2, count+1)
		c3 := lwst1(str1, str2, lengthStr1, lengthStr2, indexStr1+1, indexStr2+1, count+1)

		if c1 < c {
			c = c1
		}
		if c2 < c {
			c = c2
		}
		if c3 < c {
			c = c3
		}
	}

	return c
}

// 莱文斯坦距离
func Lwst2(str1, str2 string) int {
	lengthStr1, lengthStr2 := len(str1), len(str2)

	status := make([][]int, lengthStr1)
	for i := range status {
		status[i] = make([]int, lengthStr2)
	}

	// 初始化第一行, 以a[0]与b比较
	for i := 0; i < lengthStr2; i++ {
		if str1[0] == str2[i] {
			status[0][i] = i
		} else if i != 0 {
			status[0][i] = status[0][i-1] + 1
		} else {
			status[0][i] = 1
		}
	}

	// 初始化第1列, 已b[0]与a比较
	for i := 1; i < lengthStr1; i++ {
		if str2[0] == str1[i] {
			status[i][0] = i
		} else {
			status[i][0] = status[i-1][0] + 1
		}
	}

	min := func(c1, c2, c3 int) int {
		m := math.MaxInt64
		if c1 < m {
			m = c1
		}
		if c2 < m {
			m = c2
		}
		if c3 < m {
			m = c3
		}
		return m
	}

	// 填充矩阵数据
	for i := 1; i < lengthStr1; i++ {
		for j := 1; j < lengthStr2; j++ {
			if str1[i] != str2[j] {
				status[i][j] = min(status[i-1][j]+1, status[i][j-1]+1, status[i-1][j-1]+1)
			} else {
				status[i][j] = min(status[i-1][j]+1, status[i][j-1]+1, status[i-1][j-1])
			}
		}
	}
	return status[lengthStr1-1][lengthStr2-1]
}

// 最长子串匹配
func LangSubString1(str1, str2 string) int {
	lengthStr1, lengthStr2 := len(str1), len(str2)
	return ls1(str1, str2, lengthStr1, lengthStr2, 0, 0, 0)
}

func ls1(str1, str2 string, lengthStr1, lengthStr2, indexStr1, indexStr2 int, subCount int) int {
	if indexStr1 >= lengthStr1 || indexStr2 >= lengthStr2 {
		return subCount
	}

	count := math.MinInt64

	if str1[indexStr1] == str2[indexStr2] {
		c := ls1(str1, str2, lengthStr1, lengthStr2, indexStr1+1, indexStr2+1, subCount+1)
		if c > count {
			count = c
		}
	} else {
		c1 := ls1(str1, str2, lengthStr1, lengthStr2, indexStr1+1, indexStr2, subCount)
		c2 := ls1(str1, str2, lengthStr1, lengthStr2, indexStr1, indexStr2+1, subCount)
		if c1 > c2 {
			count = c1
		} else {
			count = c2
		}
	}
	return count
}

func LangSubString2(str1, str2 string) int {
	lenStr1, lenStr2 := len(str1), len(str2)

	// str1长度为行, str2长度为列
	status := make([][]int, lenStr1)
	for i := range status {
		status[i] = make([]int, lenStr2)
	}

	// 初始化第1行, 用str1[0] 与str2比较
	for i := 0; i < lenStr2; i++ {
		if str1[0] == str2[i] {
			status[0][i] = 1
		} else if i != 0 {
			status[0][i] = status[0][i-1]
		}
	}

	// 初始化第1列, 用str2[0]与str1比较
	for i := 1; i < lenStr1; i++ {
		if str2[0] == str1[i] {
			status[i][0] = 1
		} else if i != 0 {
			status[i][0] = status[i-1][0]
		}
	}

	max := func(s1, s2, s3 int) int {
		s := math.MinInt64
		if s1 > s {
			s = s1
		}
		if s2 > s {
			s = s2
		}
		if s3 > s {
			s = s3
		}
		return s
	}

	// 填充矩阵
	for i := 1; i < lenStr1; i++ {
		for j := 1; j < lenStr2; j++ {
			if str1[i] == str2[j] {
				status[i][j] = max(status[i-1][j], status[i][j-1], status[i-1][j-1]+1)
			} else {
				status[i][j] = max(status[i-1][j], status[i][j-1], status[i-1][j-1])
			}
		}
	}
	return status[lenStr1-1][lenStr2-1]
}
