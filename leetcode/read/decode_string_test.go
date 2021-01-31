package read

import "testing"

func TestDecodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{s: "abc3[cd]xyz"},
			want: "abccdcdcdxyz",
		},
		//{
		//	args: args{s: "3[a]2[bc]"},
		//	want: "aaabcbc",
		//},
		//{
		//	args: args{s: "3[a2[c]]"},
		//	want: "accaccacc",
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeString(tt.args.s); got != tt.want {
				t.Errorf("DecodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
