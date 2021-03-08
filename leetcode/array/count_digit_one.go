package array

func countDigitOne(n int) int {
	var (
		h     = n / 10 // 高位
		s     = 0      // 低位
		cur   = n % 10 // 当前位
		digit = 1      // 10的n次方
		ans   = 0
	)
	for n != 0 {
		switch cur {
		case 0:
			ans += h * digit
		case 1:
			ans += h*digit + s + 1
		default:
			ans += (h + 1) * digit
		}

		s += cur * digit
		cur = h % 10
		h = h / 10
		digit = digit * 10
		n = n / 10
	}
	return ans
}
