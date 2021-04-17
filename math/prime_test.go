package math

import "testing"

func TestPrintPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "100",
			args: args{
				n: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintPrime(tt.args.n)
		})
	}
}
