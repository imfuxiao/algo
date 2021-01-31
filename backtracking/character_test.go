package backtracking

import "testing"

func Test_lwst1(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				a: "mitcmu",
				b: "mtacnu",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lwst1(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Lwst1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lwst2(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				a: "mitcmu",
				b: "mtacnu",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lwst2(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Lwst2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkLwst1-12    	  405576	      2874 ns/op
func BenchmarkLwst1(b *testing.B) {
	c1 := "mitcmu"
	c2 := "mtacnu"
	for i := 0; i < b.N; i++ {
		Lwst1(c1, c2)
	}
}

//BenchmarkLwst2-12    	 4115496	       282 ns/op
func BenchmarkLwst2(b *testing.B) {
	c1 := "mitcmu"
	c2 := "mtacnu"
	for i := 0; i < b.N; i++ {
		Lwst2(c1, c2)
	}
}

func TestLangSubString1(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				str1: "mitcmu",
				str2: "mtacnu",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LangSubString1(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("LangSubString1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLangSubString2(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				str1: "mitcmu",
				str2: "mtacnu",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LangSubString2(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("LangSubString2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkLangSubString1-12    	 1874386	       617 ns/op
func BenchmarkLangSubString1(b *testing.B) {
	c1 := "mitcmu"
	c2 := "mtacnu"
	for i := 0; i < b.N; i++ {
		LangSubString1(c1, c2)
	}
}

// BenchmarkLangSubString2-12    	 3443365	       325 ns/op
func BenchmarkLangSubString2(b *testing.B) {
	c1 := "mitcmu"
	c2 := "mtacnu"
	for i := 0; i < b.N; i++ {
		LangSubString2(c1, c2)
	}
}
