package read

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{3, 2, 1},
				{3, 1, 2},
				{2, 1, 3},
				{2, 3, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
