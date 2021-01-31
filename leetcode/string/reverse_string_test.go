package string

import (
	"reflect"
	"testing"
)

func Test_reverseString1(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{s: []byte{'h', 'e', 'l', 'l', 'o'}},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if reverseString1(tt.args.s); !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Fatalf("got: %+v, want: %+v", tt.args.s, tt.want)
			}
		})
	}
}
