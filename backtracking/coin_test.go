package backtracking

import "testing"

func TestCoinChange1(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CoinChange1(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("CoinChange1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkCoinChange1-12    	  263506	      4200 ns/op
func BenchmarkCoinChange1(b *testing.B) {
	coins, amount := []int{1, 2, 5}, 11
	for i := 0; i < b.N; i++ {
		CoinChange1(coins, amount)
	}
}

func TestCoinChange3(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				coins:  []int{3, 7},
				amount: 11,
			},
			want: -1,
		},
		{
			args: args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			want: 3,
		},
		{
			args: args{
				coins:  []int{1, 2, 5, 10},
				amount: 11,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CoinChange3(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("CoinChange3() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkCoinChange3-12    	11186390	       107 ns/op
func BenchmarkCoinChange3(b *testing.B) {
	coins, amount := []int{1, 2, 5}, 11
	for i := 0; i < b.N; i++ {
		CoinChange3(coins, amount)
	}
}
