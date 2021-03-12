package linked

import "testing"

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		//{
		//	args: args{
		//		head: &ListNode{
		//			Val: 1,
		//			Next: &ListNode{
		//				Val: 2,
		//				Next: &ListNode{
		//					Val: 3,
		//					Next: &ListNode{
		//						Val: 4,
		//					},
		//				},
		//			},
		//		},
		//	},
		//},
		{
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
		})
	}
}
