package comprehensive

import "testing"

func TestIsMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			args: args{
				s: "aa",
				p: "a*",
			},
			want: true,
		},
		{
			args: args{
				s: "ab",
				p: ".*",
			},
			want: true,
		},
		{
			args: args{
				s: "aab",
				p: "c*a*b",
			},
			want: true,
		},
		{
			args: args{
				s: "mississippi",
				p: "mis*is*p*.",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("IsMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
