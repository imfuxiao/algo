package string

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "123456",
			args: args{s: "123456"},
			want: 123456,
		},
		{
			name: "-2+1",
			args: args{s: "-2+1"},
			want: -1,
		},
		{
			name: "-1+2",
			args: args{s: "-1+2"},
			want: 1,
		},
		{
			name: "1 + 1",
			args: args{s: "1 + 1"},
			want: 2,
		},
		{
			name: "2-1 + 2",
			args: args{s: "2-1 + 2"},
			want: 3,
		},
		{
			name: "(1+(4+5+2)-3)+(6+8)",
			args: args{s: "(1+(4+5+2)-3)+(6+8)"},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculate2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1-10+1",
			args: args{s: "1-10+1"},
			want: -8,
		},
		{
			name: "3+2*2",
			args: args{s: "3+2*2"},
			want: 7,
		},
		{
			name: " 3/2 ",
			args: args{s: " 3/2 "},
			want: 1,
		},
		{
			name: " 3+5 / 2 ",
			args: args{s: " 3+5 / 2 "},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate2(tt.args.s); got != tt.want {
				t.Errorf("calculate2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculate3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{s: "0"},
			want: 0,
		},
		{
			name: "11",
			args: args{s: "11"},
			want: 11,
		},
		{
			name: "1+1",
			args: args{s: "1+1"},
			want: 2,
		},
		{
			name: "6-4/2",
			args: args{s: "6-4/2"},
			want: 4,
		},
		{
			name: "2*(5+5*2)/3+(6/2+8)",
			args: args{s: "2*(5+5*2)/3+(6/2+8)"},
			want: 21,
		},
		{
			name: "(2+6*3+5-(3*14/7+2)*5)+3",
			args: args{s: "(2+6*3+5-(3*14/7+2)*5)+3"},
			want: -12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate3(tt.args.s); got != tt.want {
				t.Errorf("calculate3() = %v, want %v", got, tt.want)
			}
		})
	}
}
