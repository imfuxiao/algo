package heap

import "testing"

func TestFindKthLargest1(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{-1, 2, 0},
				k:    2,
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{2, 1},
				k:    2,
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKthLargest1(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("FindKthLargest1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindKthLargest2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{-1, 2, 0},
				k:    2,
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{2, 1},
				k:    2,
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKthLargest2(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("FindKthLargest2() = %v, want %v", got, tt.want)
			}
		})
	}
}
