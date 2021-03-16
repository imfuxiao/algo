package binarysearch

func myPow(x float64, n int) float64 {
	if n > 0 {
		return pow(x, n)
	}
	return 1 / pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	y := pow(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

func myPow1(x float64, n int) float64 {
	if n > 0 {
		return pow1(x, n)
	}
	return 1 / pow1(x, n)
}

func pow1(x float64, n int) float64 {
	ans := 1.0
	for i := n; i > 0; i /= 2 {
		if i%2 == 1 {
			ans *= x
		}
		x *= x
	}
	return ans
}
