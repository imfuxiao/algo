package read

import "testing"

func TestSubArraysDivByK(t *testing.T) {
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
				A: []int{4, 5, 0, -2, -3, 1},
				K: 5,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubArraysDivByK(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("SubArraysDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
