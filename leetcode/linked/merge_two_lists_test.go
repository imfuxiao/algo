package linked

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				l1: NewListNode(1).Add(2).Add(4),
				l2: NewListNode(1).Add(3).Add(4),
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			args: args{
				l1: nil,
				l2: nil,
			},
			want: nil,
		},
		{
			args: args{
				l1: nil,
				l2: NewListNode(0),
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := print(mergeTwoLists(tt.args.l1, tt.args.l2)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeTwoLists1(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				l1: NewListNode(1).Add(2).Add(4),
				l2: NewListNode(1).Add(3).Add(4),
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			args: args{
				l1: nil,
				l2: nil,
			},
			want: nil,
		},
		{
			args: args{
				l1: nil,
				l2: NewListNode(0),
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := print(mergeTwoLists(tt.args.l1, tt.args.l2)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeTwoLists2(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists2(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists2() = %v, want %v", got, tt.want)
			}
		})
	}
}
