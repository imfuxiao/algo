package linked

import (
	"reflect"
	"testing"
)

func Test_reverseList1(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{head: NewListNode(1).Add(2).Add(3).Add(4).Add(5)},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := print(reverseList1(tt.args.head)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			// TODO: Add test cases.
			args: args{head: NewListNode(1).Add(2).Add(3).Add(4).Add(5)},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := print(reverseList(tt.args.head)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseList3(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			// TODO: Add test cases.
			args: args{head: NewListNode(1).Add(2).Add(3).Add(4).Add(5)},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := print(reverseList3(tt.args.head)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList3() = %v, want %v", got, tt.want)
			}
		})
	}
}
