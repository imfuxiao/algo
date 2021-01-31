package algo

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{nums: []int{10, 3, 4, 9, 7}},
			want: []int{3, 4, 7, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if HeapSort(tt.args.nums); !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Fatalf("got: %+v, want: %+v", tt.args.nums, tt.want)
			}
		})
	}
}
