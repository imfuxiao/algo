package read

import "testing"

func TestMaxScoreSightseeingPair(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{A: []int{1, 3, 5}},
			want: 7,
		},
		{
			args: args{A: []int{8, 1, 5, 2, 6}},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxScoreSightseeingPair(tt.args.A); got != tt.want {
				t.Errorf("MaxScoreSightseeingPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
