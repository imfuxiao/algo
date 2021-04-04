package array

import "testing"

func Test_numRabbits(t *testing.T) {
	type args struct {
		answers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[]int{0, 0, 1, 1, 1}",
			args: args{answers: []int{0, 0, 1, 1, 1}},
			want: 6,
		},
		{
			name: "[]int{1, 0, 1, 0, 0}",
			args: args{answers: []int{1, 0, 1, 0, 0}},
			want: 5,
		},
		{
			name: "[]int{0, 2, 0, 2, 1}",
			args: args{answers: []int{0, 2, 0, 2, 1}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRabbits(tt.args.answers); got != tt.want {
				t.Errorf("numRabbits() = %v, want %v", got, tt.want)
			}
		})
	}
}
