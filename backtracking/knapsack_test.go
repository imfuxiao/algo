package backtracking

import "testing"

func Test_knapsack(t *testing.T) {
	type args struct {
		weight []int
		w      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				weight: []int{2, 5, 10, 3},
				w:      9,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knapsack(tt.args.weight, tt.args.w); got != tt.want {
				t.Errorf("knapsack() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_knapsack0(t *testing.T) {
	type args struct {
		weight []int
		w      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				weight: []int{2, 2, 4, 6, 10, 3},
				w:      9,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knapsack0(tt.args.weight, tt.args.w); got != tt.want {
				t.Errorf("knapsack0() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkKnapsack-12    	 9714747	       113 ns/op
// BenchmarkKnapsack-12    	  930126	      1256 ns/op
func BenchmarkKnapsack(b *testing.B) {
	weight := []int{2, 2, 4, 6, 10, 3, 1, 8, 7}
	w := 20
	for i := 0; i < b.N; i++ {
		knapsack(weight, w)
	}
}

// BenchmarkKnapsack0-12    	 3733075	       283 ns/op
// BenchmarkKnapsack0-12    	 2005975	       539 ns/op
func BenchmarkKnapsack0(b *testing.B) {
	weight := []int{2, 2, 4, 6, 10, 3, 1, 8, 7}
	w := 20
	for i := 0; i < b.N; i++ {
		knapsack0(weight, w)
	}
}

func Test_knapsack2(t *testing.T) {
	type args struct {
		weight []int
		w      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				weight: []int{2, 2, 4, 6, 10, 3},
				w:      9,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knapsack2(tt.args.weight, tt.args.w); got != tt.want {
				t.Errorf("knapsack2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkKnapsack2-12    	 9633274	       112 ns/op
func BenchmarkKnapsack2(b *testing.B) {
	weight := []int{2, 2, 4, 6, 10, 3, 1, 8, 7}
	w := 20
	for i := 0; i < b.N; i++ {
		knapsack2(weight, w)
	}
}

func Test_knapsack3(t *testing.T) {
	type args struct {
		weights   []int
		values    []int
		maxWeight int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				weights:   []int{2, 2, 4, 6, 3},
				values:    []int{3, 4, 8, 9, 6},
				maxWeight: 9,
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knapsack3(tt.args.weights, tt.args.values, tt.args.maxWeight); got != tt.want {
				t.Errorf("knapsack3() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkKnapsack3-12    	 6156502	       177 ns/op
func BenchmarkKnapsack3(b *testing.B) {
	weights := []int{2, 2, 4, 6, 3}
	values := []int{3, 4, 8, 9, 6}
	maxWeight := 9
	for i := 0; i < b.N; i++ {
		knapsack3(weights, values, maxWeight)
	}
}

func Test_knapsack4(t *testing.T) {
	type args struct {
		weights   []int
		values    []int
		maxWeight int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				weights:   []int{2, 2, 4, 6, 3},
				values:    []int{3, 4, 8, 9, 6},
				maxWeight: 9,
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knapsack4(tt.args.weights, tt.args.values, tt.args.maxWeight); got != tt.want {
				t.Errorf("knapsack4() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkKnapsack4-12    	17273131	        66.5 ns/op
func BenchmarkKnapsack4(b *testing.B) {
	weights := []int{2, 2, 4, 6, 3}
	values := []int{3, 4, 8, 9, 6}
	maxWeight := 9
	for i := 0; i < b.N; i++ {
		knapsack4(weights, values, maxWeight)
	}
}
