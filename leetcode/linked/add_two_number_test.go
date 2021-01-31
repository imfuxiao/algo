package linked

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				l1: NewListNode(2).Add(4).Add(3),
				l2: NewListNode(5).Add(6).Add(4),
			},
			want: NewListNode(0).Add(1),
		},
		{
			args: args{
				l1: NewListNode(5),
				l2: NewListNode(5),
			},
			want: NewListNode(0).Add(1),
		},
		{
			args: args{
				l1: NewListNode(9).Add(9).Add(9).Add(9).Add(9).Add(9),
				l2: NewListNode(9).Add(9),
			},
			want: NewListNode(0).Add(0).Add(0).Add(1),
		},
		{
			args: args{
				l1: NewListNode(1).Add(2).Add(3),
				l2: NewListNode(9).Add(8).Add(7),
			},
			want: NewListNode(0).Add(0).Add(0).Add(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
