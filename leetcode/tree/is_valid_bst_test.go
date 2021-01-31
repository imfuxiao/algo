package tree

import "testing"

func Test_isValidBST1(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				root: &TreeNode{
					Val: 2147483647,
				},
			},
			want: true,
		},
		{
			args: args{
				root: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 4},
					Right: &TreeNode{
						Val:   6,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			want: false,
		},
		{
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
			},
			want: true,
		},
		{
			args: args{
				root: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{
						Val:   6,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 6},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST1(tt.args.root); got != tt.want {
				t.Errorf("isValidBST1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBST2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				root: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 4},
					Right: &TreeNode{
						Val:   6,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			want: false,
		},
		{
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
			},
			want: true,
		},
		{
			args: args{
				root: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{
						Val:   6,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 6},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST2(tt.args.root); got != tt.want {
				t.Errorf("isValidBST2() = %v, want %v", got, tt.want)
			}
		})
	}
}
