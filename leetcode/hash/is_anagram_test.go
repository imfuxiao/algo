package hash

import "testing"

func Test_isAnagram1(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "aacc",
				t: "ccac",
			},
			want: false,
		},
		{
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram1(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAnagram2(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "aacc",
				t: "ccac",
			},
			want: false,
		},
		{
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkIsAnagram1-12    	 1554350	       748 ns/op
func BenchmarkIsAnagram1(b *testing.B) {
	s := "anagram"
	t := "nagaram"
	for i := 0; i < b.N; i++ {
		isAnagram1(s, t)
	}
}

// BenchmarkIsAnagram2-12    	  650690	      1680 ns/op
func BenchmarkIsAnagram2(b *testing.B) {
	s := "anagram"
	t := "nagaram"
	for i := 0; i < b.N; i++ {
		isAnagram2(s, t)
	}
}
