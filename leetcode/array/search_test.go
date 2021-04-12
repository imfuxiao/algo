package array

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "2, 2, 2, 3, 2, 2, 2",
			args: args{
				nums:   []int{2, 2, 2, 3, 2, 2, 2},
				target: 3,
			},
			want: true,
		},
		{
			name: "2,5,6,0,0,1,2",
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 0,
			},
			want: true,
		},
		{
			name: "2,5,6,0,0,1,2",
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 3,
			},
			want: false,
		},
		{
			name: "1, 0, 1, 1, 1",
			args: args{
				nums:   []int{1, 0, 1, 1, 1},
				target: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
