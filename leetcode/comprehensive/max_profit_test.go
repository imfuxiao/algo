package comprehensive

import "testing"

func TestMaxProfit1(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				prices: []int{3, 2, 6, 5, 0, 3},
			},
			want: 4,
		},
		{
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
		{
			args: args{
				prices: []int{7, 6, 4, 3, 2, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit1(tt.args.prices); got != tt.want {
				t.Errorf("MaxProfit1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxProfit2(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				prices: []int{3, 2, 6, 5, 0, 3},
			},
			want: 4,
		},
		{
			args: args{
				prices: []int{1, 2},
			},
			want: 1,
		},
		{
			args: args{
				prices: []int{2, 1, 2, 0, 1},
			},
			want: 1,
		},
		{
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
		{
			args: args{
				prices: []int{2, 4, 1},
			},
			want: 2,
		},
		{
			args: args{
				prices: []int{7, 6, 4, 3, 2, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit2(tt.args.prices); got != tt.want {
				t.Errorf("MaxProfit2() = %v, want %v", got, tt.want)
			}
		})
	}
}
