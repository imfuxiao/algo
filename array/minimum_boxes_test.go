package array

import "testing"

func Test_minimumBoxes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	args: args{
		//		n: 3,
		//	},
		//	want: 3,
		//},
		//{
		//	args: args{
		//		n: 4,
		//	},
		//	want: 3,
		//},
		//{
		//	args: args{
		//		n: 5,
		//	},
		//	want: 4,
		//},
		//{
		//	args: args{
		//		n: 12,
		//	},
		//	want: 8,
		//},
		{
			args: args{
				n: 13,
			},
			want: 8,
		},
		{
			args: args{
				n: 14,
			},
			want: 9,
		},
		{
			args: args{
				n: 15,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumBoxes(tt.args.n); got != tt.want {
				t.Errorf("minimumBoxes() = %v, want %v", got, tt.want)
			}
		})
	}
}
