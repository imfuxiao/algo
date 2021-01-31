// 10. 正则表达式匹配
// 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
//
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
// 所谓匹配，是要涵盖整个字符串s的，而不是部分字符串。
//
// 示例 1：
// 输入：s = "aa" p = "a"
// 输出：false
// 解释："a" 无法匹配 "aa" 整个字符串。
//
// 示例 2:
// 输入：s = "aa" p = "a*"
// 输出：true
// 解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
//
// 示例3：
// 输入：s = "ab" p = ".*"
// 输出：true
// 解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
//
// 示例 4：
// 输入：s = "aab" p = "c*a*b"
// 输出：true
// 解释：因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
//
// 示例 5：
// 输入：s = "mississippi" p = "mis*is*p*."
// 输出：false
//
// 提示：
// 0 <= s.length<= 20
// 0 <= p.length<= 30
// s可能为空，且只包含从a-z的小写字母。
// p可能为空，且只包含从a-z的小写字母，以及字符.和*。
// 保证每次出现字符 * 时，前面都匹配到有效的字符
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/regular-expression-matching
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package comprehensive

func IsMatch(s string, p string) bool {
	var (
		srcLength, patternLength      = len(s), len(p)
		dp                            = make([][]bool, srcLength+1) // 注意dp的索引, 0, 表示字符为空, 从1开始为字符的第一个字母
		start, dot               byte = '*', '.'
	)

	for i := 0; i <= srcLength; i++ {
		dp[i] = make([]bool, patternLength+1)
	}
	dp[0][0] = true

	// 初始第一行status, 即s=="", p!=""的情况
	// 因为dp中代表字符的下标从1开始, 所以这里从1开始循环并且长度+1
	// 下面的循环一样
	for i := 1; i < patternLength+1; i++ {
		if p[i-1] == start {
			if i == 1 { // p中第一个*不能消掉, 所以为false
				dp[0][i] = false
			} else {
				dp[0][i] = dp[0][i-2]
			}
		}
	}

	for i := 1; i < srcLength+1; i++ {
		for j := 1; j < patternLength+1; j++ {
			if p[j-1] != start {
				if p[j-1] == dot || p[j-1] == s[i-1] {
					dp[i][j] = dp[i-1][j-1]
				}
				// 因为dp默认为false所以这里可以省略
				//} else {
				//	dp[i][j] = false
				//}
			} else {
				if j == 1 {
					dp[i][j] = false
				} else if p[j-2] == s[i-1] || p[j-2] == dot { // j-2: *前面的字符
					dp[i][j] = dp[i-1][j] || dp[i][j-2]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}

	return dp[srcLength][patternLength]
}
