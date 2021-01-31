package read

import "testing"

func TestSubArraySum(t *testing.T) {
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
				nums: []int{1, 1, 1},
				k:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubArraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("SubArraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
