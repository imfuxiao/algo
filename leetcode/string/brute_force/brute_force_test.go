package brute_force

import "testing"

func TestBruteForce(t *testing.T) {
	type args struct {
		primary string
		model   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				primary: "helloWorld",
				model:   "World",
			},
			want: true,
		},
		{
			args: args{
				primary: "helloWorld",
				model:   "WorlD",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BruteForce(tt.args.primary, tt.args.model); got != tt.want {
				t.Errorf("BruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
