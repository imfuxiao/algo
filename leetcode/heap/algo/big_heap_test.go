package algo

import (
	"testing"
)

func TestNewBigStack(t *testing.T) {
	bigStack := NewBigStack(10)
	bigStack.Insert(1)
	bigStack.Insert(2)
	bigStack.Insert(3)
	bigStack.Insert(4)
	bigStack.Insert(5)
	bigStack.Insert(6)
	bigStack.Insert(7)
	bigStack.Insert(8)
	bigStack.Insert(9)
	bigStack.Insert(10)
	if ok := bigStack.Insert(11); ok {
		t.Fatal("insert should false")
	}

	for {
		if value, ok := bigStack.Delete(); ok {
			t.Logf("value:%d", value)
		} else {
			break
		}
	}
}
