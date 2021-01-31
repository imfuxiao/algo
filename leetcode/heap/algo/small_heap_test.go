package algo

import (
	"testing"
)

func TestNewSmallHeap(t *testing.T) {
	heep := NewSmallHeap(10)
	heep.Insert(10)
	heep.Insert(9)
	heep.Insert(8)
	heep.Insert(7)
	heep.Insert(6)
	heep.Insert(5)
	heep.Insert(4)
	heep.Insert(3)
	heep.Insert(2)
	heep.Insert(1)
	if ok := heep.Insert(0); ok {
		t.Fatal("insert should false")
	}

	for {
		if value, ok := heep.Delete(); ok {
			t.Logf("value: %d", value)
		} else {
			break
		}
	}
}
