package linked

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				head: &ListNode{
					Next: &ListNode{
						Next: &ListNode{
							Next: &ListNode{
								Next: &ListNode{
									Next: nil,
									Val:  5,
								},
								Val: 4,
							},
							Val: 3,
						},
						Val: 2,
					},
					Val: 1,
				},
				n: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		//{
		//	args: args{
		//		head: &ListNode{
		//			Next: &ListNode{
		//				Next: &ListNode{
		//					Next: &ListNode{
		//						Next: &ListNode{
		//							Next: nil,
		//							Val:  5,
		//						},
		//						Val: 4,
		//					},
		//					Val: 3,
		//				},
		//				Val: 2,
		//			},
		//			Val: 1,
		//		},
		//		k: 2,
		//	},
		//},
		{
			args: args{
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
				k:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head *ListNode
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		//{
		//	args: args{
		//		head: NewListNode(3).Add(5),
		//		m:    1,
		//		n:    2,
		//	},
		//},
		{
			args: args{
				head: NewListNode(1).Add(2).Add(3).Add(4).Add(5),
				m:    1,
				n:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
