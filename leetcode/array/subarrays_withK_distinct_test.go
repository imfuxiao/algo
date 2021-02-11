package array

import "testing"

func Test_subarraysWithKDistinct(t *testing.T) {
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
				A: []int{2, 1, 2, 1, 2},
				K: 2,
			},
			want: 10,
		},
		{
			args: args{
				A: []int{1, 2, 1, 2, 3},
				K: 2,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysWithKDistinct(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("subarraysWithKDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubstringKDistinct(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s: "abaccc",
				k: 2,
			},
			want: 4,
		},
		{
			args: args{
				s: "ccaabbb",
				k: 2,
			},
			want: 5,
		},
		{
			args: args{
				s: "ab",
				k: 1,
			},
			want: 1,
		},
		{
			args: args{
				s: "a",
				k: 0,
			},
			want: 0,
		},
		{
			args: args{
				s: "eceba",
				k: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringKDistinct(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("lengthOfLongestSubstringKDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
