package math

import "fmt"

func PrintPrime(n int) {
	var (
		prime = make([]bool, n+1)
	)
	if n >= 1 {
		fmt.Println(1)
	}
	for i := 2; i <= n; i++ {
		for index := i * 2; index <= n; index += i {
			prime[index] = true
		}
		if !prime[i] {
			fmt.Println(i)
		}
	}
}
