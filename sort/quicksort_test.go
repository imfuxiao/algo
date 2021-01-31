package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
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
				[]int{6, 4, 5, 1, 2, 3},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			args: args{
				[]int{6, 4, 4, 3, 2, 2},
			},
			want: []int{2, 2, 3, 4, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.n)
			if !reflect.DeepEqual(tt.args.n, tt.want) {
				t.Fatalf("you want:%+v, result: %+v", tt.want, tt.args.n)
			}
		})
	}
}

func Test_position(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: []int{6, 2, 5, 1, 3, 4},
			},
			want: 3,
		},
		{
			args: args{
				n: []int{2, 4, 4, 3, 2},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := position(tt.args.n); got != tt.want {
				t.Errorf("position() = %v, want %v", got, tt.want)
			}
			t.Logf("n: %+v\n", tt.args.n)
		})
	}
}
