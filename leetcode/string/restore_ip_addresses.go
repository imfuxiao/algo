package string

import "strconv"

func restoreIpAddresses(s string) []string {
	var (
		length = len(s)
		ans    = make([]string, 0, length)
		getIp  func(s, ip string, count int)
	)

	getIp = func(s, ip string, count int) {
		if len(s) == 0 && count == 4 {
			ans = append(ans, ip[:len(ip)-1])
			return
		}

		if count > 4 || (len(s) == 0 && count < 4) {
			return
		}

		for l := 1; l <= 3; l++ {
			if len(s) < l {
				break
			}
			if len(s[:l]) > 1 && s[:l][0] == '0' {
				continue
			}
			value, _ := strconv.Atoi(s[:l])
			if value <= 255 {
				getIp(s[l:], ip+s[:l]+".", count+1)
			}
		}
	}

	getIp(s, "", 0)

	return ans
}
