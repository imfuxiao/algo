package sort

import (
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
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
				n: []int{4, 5, 6, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			args: args{
				n: []int{3, 5, 4, 1, 2, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			args: args{
				n: []int{1, 2, 3, 4, 5, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.args.n)
			if !reflect.DeepEqual(tt.args.n, tt.want) {
				t.Errorf("you want: %+v, s: %+v", tt.want, tt.args.n)
			}
		})
	}
}

func TestBubbleOptimize(t *testing.T) {
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
				n: []int{4, 5, 6, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			args: args{
				n: []int{3, 5, 4, 1, 2, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			args: args{
				n: []int{1, 2, 3, 4, 5, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleOptimizeSort(tt.args.n)
			if !reflect.DeepEqual(tt.args.n, tt.want) {
				t.Errorf("you want: %+v, s: %+v", tt.want, tt.args.n)
			}
		})
	}
}

func TestBubbleSort2(t *testing.T) {
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
				n: []int{6, 3, 5, 4, 1, 2},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort2(tt.args.n)
			if !reflect.DeepEqual(tt.args.n, tt.want) {
				t.Fatalf("you want: %+v, result: %+v", tt.want, tt.args.n)
			}
		})
	}
}
