package tree

import "testing"

func Test_hasPathSum1(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
				sum:  1,
			},
			want: false,
		},
		{
			args: args{
				root: &TreeNode{Val: 1},
				sum:  1,
			},
			want: true,
		},
		{
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   11,
							Left:  &TreeNode{Val: 7},
							Right: &TreeNode{Val: 2},
						},
					},
					Right: &TreeNode{
						Val:  8,
						Left: &TreeNode{Val: 13},
						Right: &TreeNode{
							Val:   4,
							Right: &TreeNode{Val: 1},
						},
					},
				},
				sum: 22,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum1(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasPathSum2(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
				sum:  1,
			},
			want: false,
		},
		{
			args: args{
				root: &TreeNode{Val: 1},
				sum:  1,
			},
			want: true,
		},
		{
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   11,
							Left:  &TreeNode{Val: 7},
							Right: &TreeNode{Val: 2},
						},
					},
					Right: &TreeNode{
						Val:  8,
						Left: &TreeNode{Val: 13},
						Right: &TreeNode{
							Val:   4,
							Right: &TreeNode{Val: 1},
						},
					},
				},
				sum: 22,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum2(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
