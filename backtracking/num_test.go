package backtracking

import "testing"

func TestGetMaxSub(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{nums: []int{2, 9, 3, 6, 5, 1, 7}},
			want: 4,
		},
		{
			args: args{nums: []int{4, 6, 8, 9, 7}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMaxSub1(tt.args.nums); got != tt.want {
				t.Errorf("GetMaxSub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMaxSub2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{nums: []int{2, 9, 3, 6, 5, 1, 7}},
			want: 4,
		},
		{
			args: args{nums: []int{4, 6, 8, 9, 7}},
			want: 4,
		},
		{
			args: args{nums: []int{2, 8, 4, 6, 5, 9, 1, 7}},
			want: 4,
		},
		{
			args: args{nums: []int{2, 11, 6, 12, 10, 16, 14, 8, 18}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMaxSub2(tt.args.nums); got != tt.want {
				t.Errorf("GetMaxSub2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkGetMaxSub1-12    	 5491318	       196 ns/op
func BenchmarkGetMaxSub1(b *testing.B) {
	nums := []int{2, 11, 6, 12, 10, 16, 14, 8, 18}
	for i := 0; i < b.N; i++ {
		GetMaxSub1(nums)
	}
}

//BenchmarkGetMaxSub2-12    	15412532	        75.0 ns/op
func BenchmarkGetMaxSub2(b *testing.B) {
	nums := []int{2, 11, 6, 12, 10, 16, 14, 8, 18}
	for i := 0; i < b.N; i++ {
		GetMaxSub2(nums)
	}
}
