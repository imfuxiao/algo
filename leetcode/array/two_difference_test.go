package array

import (
	"reflect"
	"testing"
)

func TestTwoDifference(t *testing.T) {
	type args struct {
		num    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				num:    []int{1, 1, 9, 3, 8, 7, 6, 5, -9},
				target: 5,
			},
			want: [][]int{
				{6, 1},
				{8, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoDifference(tt.args.num, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
