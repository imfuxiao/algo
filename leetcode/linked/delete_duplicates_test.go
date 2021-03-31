package linked

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{head: NewListNode(1).Add(2).Add(2)},
		},
		{
			args: args{
				head: &ListNode{
					Next: &ListNode{
						Next: &ListNode{
							Next: &ListNode{
								Next: &ListNode{
									Next: &ListNode{
										Next: &ListNode{
											Next: nil,
											Val:  5,
										},
										Val: 4,
									},
									Val: 4,
								},
								Val: 3,
							},
							Val: 3,
						},
						Val: 2,
					},
					Val: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
