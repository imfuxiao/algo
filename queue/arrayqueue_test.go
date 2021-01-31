package queue

import (
	"reflect"
	"testing"
)

func TestArrayQueue_EnQueue(t *testing.T) {
	type fields struct {
		arrayQueue *ArrayQueue
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        bool
		wantElement string
	}{
		{
			fields: fields{
				arrayQueue: New(3),
			},
			args: args{
				v: 1,
			},
			want:        true,
			wantElement: "[1]",
		},
		{
			fields: fields{
				arrayQueue: New(1),
			},
			args: args{
				v: 2,
			},
			want:        true,
			wantElement: "[2]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.arrayQueue.EnQueue(tt.args.v); got != tt.want {
				t.Errorf("ArrayQueue.EnQueue() = %v, want %v", got, tt.want)
			}
			if str := tt.fields.arrayQueue.String(); str != tt.wantElement {
				t.Errorf("ArrayQueue.String() = %v, wantElement %v", str, tt.wantElement)
			}
		})
	}
}

func TestArrayQueue_DeQueue(t *testing.T) {
	type fields struct {
		arrayQueue *ArrayQueue
	}
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name             string
		fields           fields
		argsFirst        args
		wantFirst        interface{}
		wantFirstString  string
		argsSecond       args
		wantSecond       interface{}
		wantElement      string
		wantSecondString string
	}{
		{
			fields: fields{
				arrayQueue: New(2),
			},
			argsFirst: args{
				v: []interface{}{1, 2, 3},
			},
			wantFirst:       1,
			wantFirstString: "[2]",
			argsSecond: args{
				v: []interface{}{3},
			},
			wantSecond:       2,
			wantSecondString: "[3]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, value := range tt.argsFirst.v {
				tt.fields.arrayQueue.EnQueue(value)
			}
			if got := tt.fields.arrayQueue.DeQueue(); !reflect.DeepEqual(got, tt.wantFirst) {
				t.Errorf("ArrayQueue.DeQueue() = %v, want %v", got, tt.wantFirst)
			}
			if str := tt.fields.arrayQueue.String(); str != tt.wantFirstString {
				t.Errorf("ArrayQueue.String() = %v, wantFirstString %v", str, tt.wantFirstString)
			}
			t.Logf("arrayQueue: %+v", tt.fields.arrayQueue.elements)

			for _, value := range tt.argsSecond.v {
				tt.fields.arrayQueue.EnQueue(value)
			}
			t.Logf("arrayQueue: %+v", tt.fields.arrayQueue.elements)

			if got := tt.fields.arrayQueue.DeQueue(); !reflect.DeepEqual(got, tt.wantSecond) {
				t.Errorf("ArrayQueue.DeQueue() = %v, want %v", got, tt.wantSecond)
			}
			if str := tt.fields.arrayQueue.String(); str != tt.wantSecondString {
				t.Errorf("ArrayQueue.String() = %v, wantSecondStr %v", str, tt.wantSecondString)
			}
			t.Logf("arrayQueue: %+v", tt.fields.arrayQueue.elements)

			value := tt.fields.arrayQueue.DeQueue()
			t.Logf("arrayQueue.DeQueue: %+v", value)
			t.Logf("arrayQueue: %+v, head: %d, tail: %d", tt.fields.arrayQueue, tt.fields.arrayQueue.head, tt.fields.arrayQueue.tail)

			tt.fields.arrayQueue.EnQueue(1)
			t.Logf("arrayQueue: %+v, head: %d, tail: %d", tt.fields.arrayQueue, tt.fields.arrayQueue.head, tt.fields.arrayQueue.tail)

			tt.fields.arrayQueue.EnQueue(2)
			t.Logf("arrayQueue: %+v, head: %d, tail: %d", tt.fields.arrayQueue, tt.fields.arrayQueue.head, tt.fields.arrayQueue.tail)

			tt.fields.arrayQueue.EnQueue(3)
			t.Logf("arrayQueue: %+v, head: %d, tail: %d", tt.fields.arrayQueue, tt.fields.arrayQueue.head, tt.fields.arrayQueue.tail)
		})
	}
}
