package hash

import (
	"reflect"
	"testing"
)

func TestFourSum1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				nums:   []int{1, -2, -5, -4, -3, 3, 3, 5},
				target: -11,
			},
			want: [][]int{
				{-5, -4, -3, 1},
			},
		},
		//{
		//	args: args{
		//		nums:   []int{1, 0, -1, 0, -2, 2},
		//		target: 0,
		//	},
		//	want: [][]int{
		//		{-2, -1, 1, 2},
		//		{-2, 0, 0, 2},
		//		{-1, 0, 0, 1},
		//	},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FourSum1(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FourSum1() = %v, want %v", got, tt.want)
			}
		})
	}
}
