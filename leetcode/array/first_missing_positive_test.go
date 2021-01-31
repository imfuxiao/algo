package array

import "testing"

func TestFirstMissingPositive1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{nums: []int{1, 2, 0}},
			want: 3,
		},
		{
			args: args{nums: []int{3, 4, -1, 1}},
			want: 2,
		},
		{
			args: args{nums: []int{7, 8, 9, 11, 12}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstMissingPositive1(tt.args.nums); got != tt.want {
				t.Errorf("FirstMissingPositive1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstMissingPositive2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{nums: []int{1}},
			want: 2,
		},
		{
			args: args{nums: []int{1, 2, 0}},
			want: 3,
		},
		{
			args: args{nums: []int{3, 4, -1, 1}},
			want: 2,
		},
		{
			args: args{nums: []int{7, 8, 9, 11, 12}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstMissingPositive2(tt.args.nums); got != tt.want {
				t.Errorf("FirstMissingPositive2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkFirstMissingPositive1-12    	121787240	         9.50 ns/op
func BenchmarkFirstMissingPositive1(b *testing.B) {
	num := []int{3, 4, -1, 1}
	for i := 0; i < b.N; i++ {
		FirstMissingPositive1(num)
	}
}

// BenchmarkFirstMissingPositive2-12    	84687568	        13.5 ns/op
func BenchmarkFirstMissingPositive2(b *testing.B) {
	num := []int{3, 4, -1, 1}
	for i := 0; i < b.N; i++ {
		FirstMissingPositive2(num)
	}
}

// BenchmarkFirstMissingPositive3-12    	194350959	         5.50 ns/op
func BenchmarkFirstMissingPositive3(b *testing.B) {
	num := []int{3, 4, -1, 1}
	for i := 0; i < b.N; i++ {
		FirstMissingPositive3(num)
	}
}

func TestFirstMissingPositive3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{nums: []int{1}},
			want: 2,
		},
		{
			args: args{nums: []int{1, 2, 0}},
			want: 3,
		},
		{
			args: args{nums: []int{3, 4, -1, 1}},
			want: 2,
		},
		{
			args: args{nums: []int{7, 8, 9, 11, 12}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstMissingPositive3(tt.args.nums); got != tt.want {
				t.Errorf("FirstMissingPositive3() = %v, want %v", got, tt.want)
			}
		})
	}
}
