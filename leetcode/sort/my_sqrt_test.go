package sort

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				x: 2147395600,
			},
			want: 46340,
		},
		{
			args: args{
				x: 5,
			},
			want: 2,
		},
		{
			args: args{
				x: 8,
			},
			want: 2,
		},
		{
			args: args{
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt1(tt.args.x); got != tt.want {
				t.Errorf("mySqrt1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mySqrt2(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				x: 2147395600,
			},
			want: 46340,
		},
		{
			args: args{
				x: 5,
			},
			want: 2,
		},
		{
			args: args{
				x: 8,
			},
			want: 2,
		},
		{
			args: args{
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt2(tt.args.x); got != tt.want {
				t.Errorf("mySqrt2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkMySqrt1-12    	41631570	        26.7 ns/op
func BenchmarkMySqrt1(b *testing.B) {
	x := 2147395600
	for i := 0; i < b.N; i++ {
		mySqrt1(x)
	}
}

// BenchmarkMySqrt2-12    	35535472	        33.7 ns/op
func BenchmarkMySqrt2(b *testing.B) {
	x := 2147395600
	for i := 0; i < b.N; i++ {
		mySqrt2(x)
	}
}
