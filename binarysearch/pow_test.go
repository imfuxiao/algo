package binarysearch

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "2.0, 10",
			args: args{
				x: 2.0,
				n: 10,
			},
			want: 1024.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myPow1(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{
				x: 2.0,
				n: 10,
			},
			want: 1024.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow1(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow1() = %v, want %v", got, tt.want)
			}
		})
	}
}
