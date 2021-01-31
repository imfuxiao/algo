package comprehensive

import "testing"

func TestMaxSubArray1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArray1(tt.args.nums); got != tt.want {
				t.Errorf("MaxSubArray1() = %v, want %v", got, tt.want)
			}
		})
	}
}
