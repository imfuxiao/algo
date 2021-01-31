package comprehensive

import "testing"

func TestMinPathSum1(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				grid: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinPathSum1(tt.args.grid); got != tt.want {
				t.Errorf("MinPathSum1() = %v, want %v", got, tt.want)
			}
		})
	}
}
