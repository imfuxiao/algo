package linked

import "testing"

func Test_deleteNode(t *testing.T) {
	type args struct {
		node *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				node: &ListNode{
					Next: &ListNode{
						Next: &ListNode{
							//Next: &ListNode{
							//	Next: nil,
							//	Val:  9,
							//},
							Val: 1,
						},
						Val: 5,
					},
					Val: 4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%+v", tt.args)
			deleteNode(tt.args.node)
			t.Logf("%+v", tt.args)
		})
	}
}
