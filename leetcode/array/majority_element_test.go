package array

import "testing"

func TestMajorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{nums: []int{3, 2, 3}},
			want: 3,
		},
		{
			args: args{nums: []int{2, 2, 1, 1, 1, 2, 2}},
			want: 2,
		},
		{
			args: args{nums: []int{-1, 1, 1, 1, 2, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MajorityElement1(tt.args.nums); got != tt.want {
				t.Errorf("MajorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMajorityElement2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{nums: []int{3, 2, 3}},
			want: 3,
		},
		{
			args: args{nums: []int{2, 2, 1, 1, 1, 2, 2}},
			want: 2,
		},
		{
			args: args{nums: []int{-1, 1, 1, 1, 2, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MajorityElement2(tt.args.nums); got != tt.want {
				t.Errorf("MajorityElement2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkMajorityElement1-12    	17103988	        62.7 ns/op
func BenchmarkMajorityElement1(b *testing.B) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	for i := 0; i < b.N; i++ {
		MajorityElement1(nums)
	}
}

// BenchmarkMajorityElement2-12    	29867096	        35.1 ns/op
func BenchmarkMajorityElement2(b *testing.B) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	for i := 0; i < b.N; i++ {
		MajorityElement2(nums)
	}
}

// BenchmarkMajorityElement3-12    	179670356	         6.44 ns/op
func BenchmarkMajorityElement3(b *testing.B) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	for i := 0; i < b.N; i++ {
		MajorityElement3(nums)
	}
}

func TestMajorityElement3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{nums: []int{3, 2, 3}},
			want: 3,
		},
		{
			args: args{nums: []int{2, 2, 1, 1, 1, 2, 2}},
			want: 2,
		},
		{
			args: args{nums: []int{-1, 1, 1, 1, 2, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MajorityElement3(tt.args.nums); got != tt.want {
				t.Errorf("MajorityElement3() = %v, want %v", got, tt.want)
			}
		})
	}
}
