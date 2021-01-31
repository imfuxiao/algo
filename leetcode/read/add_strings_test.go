package read

import "testing"

func TestAddStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				num1: "9",
				num2: "99",
			},
			want: "108",
		},
		{
			args: args{
				num1: "3876620623801494171",
				num2: "6529364523802684779",
			},
			want: "10405985147604178950",
		},
		{
			args: args{
				num1: "123",
				num2: "321",
			},
			want: "444",
		},
		{
			args: args{
				num1: "100",
				num2: "200",
			},
			want: "300",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("AddStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
