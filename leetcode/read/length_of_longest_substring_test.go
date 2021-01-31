package read

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{s: ""},
			want: 0,
		},
		{
			args: args{s: " "},
			want: 1,
		},
		{
			args: args{s: "abcabcbb"},
			want: 3,
		},
		{
			args: args{s: "bbbbb"},
			want: 1,
		},
		{
			args: args{s: "pwwkew"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
