package recursion

import (
	"testing"
)

func TestStep(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			args: args{
				n: 5,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Step(tt.args.n); got != tt.want {
				t.Errorf("Step() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllPermutation(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				data: []int{1, 2, 3},
			},
		},
		{
			args: args{
				data: []int{1, 2, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AllPermutation(tt.args.data)
		})
	}
}
