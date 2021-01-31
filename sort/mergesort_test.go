package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				n: []int{6, 4, 5},
			},
			want: []int{4, 5, 6},
		},
		{
			args: args{
				n: []int{6, 4, 5, 1, 2, 3},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.n)
			if !reflect.DeepEqual(tt.args.n, tt.want) {
				t.Fatalf("you want: %+v, result: %+v", tt.want, tt.args.n)
			}
		})
	}
}
