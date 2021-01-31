package algo

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{nums: []int{1}},
			want: []int{1},
		},
		{
			args: args{nums: []int{1, 2, 3}},
			want: []int{1, 2, 3},
		},
		{
			args: args{nums: []int{2, 1, 9, 7, 3}},
			want: []int{1, 2, 3, 7, 9},
		},
		{
			args: args{nums: []int{2, 1, 9, 7}},
			want: []int{1, 2, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if QuickSort(tt.args.nums); !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Fatalf("got: %+v, want: %+v", tt.args.nums, tt.want)
			}
		})
	}
}
