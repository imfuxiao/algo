package read

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				intervals: [][]int{
					{1, 4}, {0, 0},
				},
			},
			want: [][]int{
				{0, 0}, {1, 4},
			},
		},
		{
			args: args{
				intervals: [][]int{
					{1, 3}, {2, 6}, {8, 10}, {15, 18},
				},
			},
			want: [][]int{
				{1, 6}, {8, 10}, {15, 18},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
