package array

import "testing"

func Test_countDigitOne(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{n: 120},
			want: 53,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDigitOne(tt.args.n); got != tt.want {
				t.Errorf("countDigitOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
