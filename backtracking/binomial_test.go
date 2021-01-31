package backtracking

import "testing"

func Test_binomial1(t *testing.T) {
	type args struct {
		binomialArray [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{[][]int{
				[]int{5},
				[]int{7, 8},
				[]int{2, 3, 4},
				[]int{4, 9, 6, 1},
				[]int{2, 7, 9, 4, 5},
			}},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binomial1(tt.args.binomialArray); got != tt.want {
				t.Errorf("binomial1() = %v, want %v", got, tt.want)
			}
		})
	}
}
