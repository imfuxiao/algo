package comprehensive

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{nums: []int{-4, -3}},
			want: 12,
		},
		{
			args: args{nums: []int{-2}},
			want: -2,
		},
		{
			args: args{nums: []int{2, 3, -2, 4}},
			want: 6,
		},
		{
			args: args{nums: []int{-2, 0, -1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProduct1(tt.args.nums); got != tt.want {
				t.Errorf("MaxProduct1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxProduct2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{nums: []int{2, -5, -2, -4, 3}},
			want: 24,
		},
		{
			args: args{nums: []int{-4, -3}},
			want: 12,
		},
		{
			args: args{nums: []int{-2}},
			want: -2,
		},
		{
			args: args{nums: []int{2, 3, -2, 4}},
			want: 6,
		},
		{
			args: args{nums: []int{-2, 0, -1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProduct2(tt.args.nums); got != tt.want {
				t.Errorf("MaxProduct2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxProduct3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{nums: []int{0, -5, -2, -4, 3}},
			want: 24,
		},
		{
			args: args{nums: []int{2, -5, -2, -4, 3}},
			want: 24,
		},
		{
			args: args{nums: []int{-4, -3}},
			want: 12,
		},
		{
			args: args{nums: []int{-2}},
			want: -2,
		},
		{
			args: args{nums: []int{2, 3, -2, 4}},
			want: 6,
		},
		{
			args: args{nums: []int{-2, 0, -1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProduct3(tt.args.nums); got != tt.want {
				t.Errorf("MaxProduct2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkMaxProduct2-12    	30338176	        37.0 ns/op
func BenchmarkMaxProduct2(b *testing.B) {
	nums := []int{2, -5, -2, -4, 3}
	for i := 0; i < b.N; i++ {
		MaxProduct2(nums)
	}
}

// BenchmarkMaxProduct3-12    	142296064	         7.57 ns/op
func BenchmarkMaxProduct3(b *testing.B) {
	nums := []int{2, -5, -2, -4, 3}
	for i := 0; i < b.N; i++ {
		MaxProduct3(nums)
	}
}
