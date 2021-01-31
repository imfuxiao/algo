package single

import (
	"reflect"
	"testing"
)

func TestArrayStack_Pop(t *testing.T) {
	stack := NewArrayStack(3)
	_ = stack.Push(1)
	_ = stack.Push(2)
	_ = stack.Push(3)
	got, err := stack.Pop()
	if err != nil {
		t.Fatalf("do not get error: %+v", err)
	}
	want := interface{}(3)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got: %+v, want: %+v", got, want)
	}
}

func TestArrayStack_Push(t *testing.T) {
	stack := NewArrayStack(3)
	_ = stack.Push(1)
	_ = stack.Push(2)
	_ = stack.Push(3)
	if err := stack.Push(4); err == nil {
		t.Fatalf("must get error")
	}
	if got, want := stack.Print(), []interface{}{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got: %+v; want: %+v", got, want)
	}
}
