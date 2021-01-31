package linked

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	n21 := &ListNode{
		Val: 3,
	}
	n22 := &ListNode{
		Val: 2,
	}
	n23 := &ListNode{
		Val: 0,
	}
	n24 := &ListNode{
		Val: -4,
	}
	n21.Next = n22
	n22.Next = n23
	n23.Next = n24

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{head: n21},
			want: n23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_middleNode2(t *testing.T) {
	n21 := &ListNode{
		Val: 3,
	}
	n22 := &ListNode{
		Val: 2,
	}
	n23 := &ListNode{
		Val: 0,
	}
	n24 := &ListNode{
		Val: -4,
	}
	n25 := &ListNode{
		Val: -5,
	}
	n21.Next = n22
	n22.Next = n23
	n23.Next = n24
	n24.Next = n25

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{head: n21},
			want: n23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode2() = %v, want %v", got, tt.want)
			}
		})
	}
}
