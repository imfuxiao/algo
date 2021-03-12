package trie

import "testing"

func Test_findKthNumber(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "10",
			args: args{
				n: 10,
				k: 3,
			},
			want: 2,
		},
		{
			name: "13",
			args: args{
				n: 13,
				k: 2,
			},
			want: 10,
		},
		{
			name: "23",
			args: args{
				n: 23,
				k: 16,
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthNumber(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findKthNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
