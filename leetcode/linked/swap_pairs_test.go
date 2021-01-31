package linked

import (
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{head: NewListNode(1).Add(2)},
			want: []int{2, 1},
		},
		{
			args: args{head: NewListNode(1).Add(2).Add(3).Add(4).Add(5)},
			want: []int{2, 1, 4, 3, 5},
		},
		{
			args: args{head: NewListNode(1)},
			want: []int{1},
		},
		{
			args: args{head: NewListNode(2).Add(5).Add(3).Add(4).Add(6).Add(2).Add(2)},
			want: []int{5, 2, 4, 3, 2, 6, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := print(swapPairs(tt.args.head)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swapPairs2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{head: NewListNode(1).Add(2)},
			want: []int{2, 1},
		},
		{
			args: args{head: NewListNode(1).Add(2).Add(3).Add(4).Add(5)},
			want: []int{2, 1, 4, 3, 5},
		},
		{
			args: args{head: NewListNode(1)},
			want: []int{1},
		},
		{
			args: args{head: NewListNode(2).Add(5).Add(3).Add(4).Add(6).Add(2).Add(2)},
			want: []int{5, 2, 4, 3, 2, 6, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := print(swapPairs2(tt.args.head)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs2() = %v, want %v", got, tt.want)
			}
		})
	}
}
