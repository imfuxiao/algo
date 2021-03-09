package array

import "testing"

func Test_maxEnvelopes(t *testing.T) {
	type args struct {
		envelopes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{envelopes: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 5},
				{5, 6},
				{5, 5},
				{6, 7},
				{7, 8},
			}},
			want: 7,
		},
		{
			name: "2",
			args: args{envelopes: [][]int{
				{5, 4},
				{6, 4},
				{6, 7},
				{2, 3},
			}},
			want: 3,
		},
		{
			name: "3",
			args: args{envelopes: [][]int{
				{30, 60},
				{12, 2},
				{3, 4},
				{12, 15},
			}},
			want: 3,
		},
		{
			name: "4",
			args: args{envelopes: [][]int{
				{2, 100},
				{3, 200},
				{4, 300},
				{5, 500},
				{5, 400},
				{5, 250},
				{6, 370},
				{6, 360},
				{7, 380},
			}},
			want: 5,
		},
		{
			name: "5",
			args: args{envelopes: [][]int{
				{4, 5},
				{4, 6},
				{6, 7},
				{2, 3},
				{1, 1},
			}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}
}
