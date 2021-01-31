package backtracking

import "testing"

// BenchmarkFib-12    	   51253	     21928 ns/op
func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(20)
	}
}
// BenchmarkFib2-12    	 7674883	       148 ns/op
func BenchmarkFib2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib2(20)
	}
}

//BenchmarkFib3-12    	17795608	        68.3 ns/op
func BenchmarkFib3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib3(20)
	}
}

// BenchmarkFib4-12    	166816629	         6.73 ns/op
func BenchmarkFib4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib4(20)
	}
}

