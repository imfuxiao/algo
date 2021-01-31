package stack

import "testing"

func Test_isValid1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{s: "(("},
			want: false,
		},
		{
			args: args{s: "()"},
			want: true,
		},
		{
			args: args{s: "()[]{}"},
			want: true,
		},
		{
			args: args{s: "(]"},
			want: false,
		},
		{
			args: args{s: "([)]"},
			want: false,
		},
		{
			args: args{s: "{[]}"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid1(tt.args.s); got != tt.want {
				t.Errorf("isValid1() = %v, want %v", got, tt.want)
			}
		})
	}
}
