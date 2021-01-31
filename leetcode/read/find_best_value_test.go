package read

import "testing"

func TestFindBestValue(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				arr:    []int{48772, 52931, 14253, 32289, 75263},
				target: 40876,
			},
			want: 8175,
		},
		{
			args: args{
				arr:    []int{4, 9, 3},
				target: 10,
			},
			want: 3,
		},
		//{
		//	args: args{
		//		arr:    []int{60864, 25176, 27249, 21296, 20204}, // 154789 /5 = 30957.8
		//		target: 56803,
		//	},
		//	want: 11361,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindBestValue(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("FindBestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{num: -4},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.num); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
