package backtracking

import "testing"

func Test_matrix1(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{matrix: [][]int{
				{1, 3, 5, 9},
				{2, 1, 3, 4},
				{5, 2, 6, 7},
				{6, 8, 4, 3},
			}},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrix1(tt.args.matrix); got != tt.want {
				t.Errorf("matrix1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkMatrix1-12    	 4285246	       248 ns/op
func BenchmarkMatrix1(b *testing.B) {
	matrix := [][]int{
		{1, 3, 5, 9},
		{2, 1, 3, 4},
		{5, 2, 6, 7},
		{6, 8, 4, 3},
	}
	for i := 0; i < b.N; i++ {
		matrix1(matrix)
	}
}

// BenchmarkMatrix2-12    	15005472	        72.9 ns/op
func BenchmarkMatrix2(b *testing.B) {
	matrix := [][]int{
		{1, 3, 5, 9},
		{2, 1, 3, 4},
		{5, 2, 6, 7},
		{6, 8, 4, 3},
	}
	for i := 0; i < b.N; i++ {
		matrix2(matrix)
	}
}

func Test_matrix2(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{matrix: [][]int{
				{1, 3, 5, 9},
				{2, 1, 3, 4},
				{5, 2, 6, 7},
				{6, 8, 4, 3},
			}},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrix2(tt.args.matrix); got != tt.want {
				t.Errorf("matrix2() = %v, want %v", got, tt.want)
			}
		})
	}
}
