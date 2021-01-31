package linked

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{head: NewListNode(1).Add(2).Add(3).Add(2).Add(1)},
			want: true,
		},
		{
			args: args{head: NewListNode(1).Add(2).Add(2).Add(1)},
			want: true,
		},
		{
			args: args{head: nil},
			want: true,
		},
		{
			args: args{head: NewListNode(1).Add(2)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
