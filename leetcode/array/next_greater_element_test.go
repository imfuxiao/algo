package array

import "testing"

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2147483486",
			args: args{
				n: 2147483486,
			},
			want: -1,
		},
		{
			name: "12",
			args: args{
				n: 12,
			},
			want: 21,
		},
		{
			name: "0",
			args: args{
				n: 0,
			},
			want: -1,
		},
		{
			name: "132",
			args: args{
				n: 132,
			},
			want: 213,
		},
		{
			name: "4321",
			args: args{
				n: 4321,
			},
			want: -1,
		},
		{
			name: "1243",
			args: args{
				n: 1243,
			},
			want: 1324,
		},
		{
			name: "230241",
			args: args{
				n: 230241,
			},
			want: 230412,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.args.n); got != tt.want {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
