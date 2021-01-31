// 242. 有效的字母异位词
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
// 示例 1:
//
// 输入: s = "anagram", t = "nagaram"
// 输出: true
// 示例 2:
//
// 输入: s = "rat", t = "car"
// 输出: false
// 说明:
// 你可以假设字符串只包含小写字母。
//
// 进阶:
// 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/valid-anagram
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package hash

import (
	"reflect"
	"sort"
)

// 用hashMap来实现
func isAnagram1(s string, t string) bool {
	lengthS := len(s)
	if lengthS != len(t) {
		return false
	}

	set := make(map[byte]int, len(s))
	for i := range s {
		set[s[i]] = set[s[i]] + 1
	}
	for j := range t {
		if _, exists := set[t[j]]; !exists {
			return false
		} else {
			set[t[j]] = set[t[j]] - 1
		}
	}
	for _, v := range set {
		if v != 0 {
			return false
		}
	}
	return true
}

// 用排序法
func isAnagram2(s string, t string) bool {
	s1 := []byte(s)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})

	t1 := []byte(t)
	sort.Slice(t1, func(i, j int) bool {
		return t1[i] < t1[j]
	})
	return reflect.DeepEqual(s1, t1)
}
