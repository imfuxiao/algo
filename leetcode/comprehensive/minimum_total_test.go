package comprehensive

import "testing"

func TestMinimumTotal1(t *testing.T) {
	type args struct {
		triangle [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumTotal1(tt.args.triangle); got != tt.want {
				t.Errorf("MinimumTotal1() = %v, want %v", got, tt.want)
			}
		})
	}
}
