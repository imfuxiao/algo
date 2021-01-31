package backtracking

import (
	"testing"
)

func TestQueen(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Queen()
		})
	}
}

func Test_checkPositionIsOk(t *testing.T) {
	type args struct {
		queens [9]int
		row    int
		column int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				queens: [9]int{0, 1, 8, 6, 0, 0, 0, 0},
				row:    4,
				column: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPositionIsOk(&tt.args.queens, tt.args.row, tt.args.column); got != tt.want {
				t.Errorf("checkPositionIsOk() = %v, want %v", got, tt.want)
			}
		})
	}
}
