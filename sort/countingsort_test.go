package sort

import (
	"reflect"
	"testing"
)

func TestCountingSort(t *testing.T) {
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
				n: []int{6, 2, 3, 5, 4, 1},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CountingSort(tt.args.n)
			if !reflect.DeepEqual(tt.args.n, tt.want) {
				t.Fatalf("you want %+v, result %+v", tt.want, tt.args.n)
			}
		})
	}
}
