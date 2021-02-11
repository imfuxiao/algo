package stack

import "testing"

func TestKthLargest_Add(t *testing.T) {
	obj := Constructor(3, []int{4, 5, 8, 2});
	got := obj.Add(3)
	if want := 4; got != want {
		t.Fatalf("want: %d, got: %d", want, got)
	}
	got = obj.Add(5)
	if want := 5; got != want {
		t.Fatalf("want: %d, got: %d", want, got)
	}
}
