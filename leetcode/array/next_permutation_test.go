package array

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{nums: []int{3, 2, 1}},
			want: []int{1, 2, 3},
		},
		{
			args: args{nums: []int{1, 2, 3}},
			want: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("want: %+v, got: %+v", tt.want, tt.args.nums)
			}
		})
	}
}
