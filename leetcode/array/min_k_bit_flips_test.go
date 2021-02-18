package array

import "testing"

func Test_minKBitFlips(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				A: []int{0, 0, 0, 1, 0, 1, 1, 0},
				K: 3,
			},
			want: 3,
		},
		{
			args: args{
				A: []int{1, 1, 0},
				K: 2,
			},
			want: -1,
		},
		{
			args: args{
				A: []int{0, 1, 0},
				K: 1,
			},
			want: 2,
		},
		{
			args: args{
				A: []int{1, 1},
				K: 2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minKBitFlips(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("minKBitFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
