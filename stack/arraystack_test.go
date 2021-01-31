package stack

import (
	"reflect"
	"testing"
)

func TestArrayStack_Push(t *testing.T) {
	type fields struct {
		stack *ArrayStack
	}
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			fields: fields{
				stack: New(0),
			},
			args: args{
				v: []interface{}{1, 2, 3},
			},
			want: "[1 2 3]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.stack.Push(tt.args.v...)
			if s := tt.fields.stack.String(); s != tt.want {
				t.Fatalf("you want: %+v, result: %+v", tt.want, s)
			}
		})
	}
}

func TestArrayStack_Pop(t *testing.T) {
	type fields struct {
		stack *ArrayStack
	}
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     interface{}
		wantSize int
	}{
		{
			fields: fields{
				stack: New(0),
			},
			args: args{
				v: []interface{}{1, 2, 3},
			},
			want:     3,
			wantSize: 2,
		},
		{
			fields: fields{
				stack: New(0),
			},
			args: args{
				v: nil,
			},
			want:     nil,
			wantSize: 0,
		},
		{
			fields: fields{
				stack: New(3),
			},
			args: args{
				v: []interface{}{1, 2, 3, 4},
			},
			want:     nil,
			wantSize: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.stack.Push(tt.args.v...)
			if got := tt.fields.stack.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayStack.Pop() = %v, want %v", got, tt.want)
			}
			if got := tt.fields.stack.size; !reflect.DeepEqual(got, tt.wantSize) {
				t.Errorf("ArrayStack.count = %v, wantSize %v", got, tt.wantSize)
			}
		})
	}
}

func TestArrayStack_Size(t *testing.T) {
	type fields struct {
		stack *ArrayStack
	}
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			fields: fields{
				stack: New(3),
			},
			args: args{
				v: []interface{}{1, 2, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.stack.Push(tt.args.v...)
			if got := tt.fields.stack.Size(); got != tt.want {
				t.Errorf("ArrayStack.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
