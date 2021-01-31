package queue

import "testing"

func TestMyCircularDeque(t *testing.T) {
	circularDeque := Constructor(3) // 设置容量大小为3

	check := circularDeque.InsertLast(1) // 返回 true
	if !check {
		t.Fatalf("InsertLast is error")
	}

	check = circularDeque.InsertLast(2) // 返回 true
	if !check {
		t.Fatalf("InsertLast is error")
	}

	check = circularDeque.InsertFront(3) // 返回 true
	if !check {
		t.Fatalf("InsertFront is error")
	}

	check = circularDeque.InsertFront(4) // 已经满了，返回 false
	if check {
		t.Fatalf("InsertFront is error")
	}

	value := circularDeque.GetRear() // 返回 2
	if value != 2 {
		t.Fatalf("GetRear is error")
	}

	check = circularDeque.IsFull() // 返回 true
	if !check {
		t.Fatalf("IsFull is error")
	}

	check = circularDeque.DeleteLast() // 返回 true
	if !check {
		t.Fatalf("DeleteLast is error")
	}

	check = circularDeque.InsertFront(4) // 返回 true
	if !check {
		t.Fatalf("InsertFront is error")
	}

	value = circularDeque.GetFront() // 返回 4
	if value != 4 {
		t.Fatalf("GetFront is error")
	}
}

func TestMyCircularDeque2(t *testing.T) {

	// ["MyCircularDeque","insertFront","insertLast","insertLast","getFront","deleteLast","getRear","insertFront","deleteFront","getRear","insertLast","isFull"]
	// [[3],[8],[8],[2],[],[],[],[9],[],[],[2],[]]
	// [null,true,true,true,8,true,8,true,true,8,true,true]

	circularDeque := Constructor(3)

	check := circularDeque.InsertFront(8)
	if !check {
		t.Fatalf("InsertFront is error")
	}

	check = circularDeque.InsertLast(8)
	if !check {
		t.Fatalf("InsertLast is error")
	}

	check = circularDeque.InsertLast(2)
	if !check {
		t.Fatalf("InsertLast is error")
	}

	value := circularDeque.GetFront()
	if value != 8 {
		t.Fatalf("GetFront is error")
	}

	check = circularDeque.DeleteLast()
	if !check {
		t.Fatalf("DeleteLast is error")
	}

	value = circularDeque.GetRear()
	if value != 8 {
		t.Fatalf("GetRear is error")
	}

	check = circularDeque.InsertFront(9)
	if !check {
		t.Fatalf("InsertFront is error")
	}

	check = circularDeque.DeleteFront()
	if !check {
		t.Fatalf("DeleteFront is error")
	}

	check = circularDeque.InsertLast(2)
	if !check {
		t.Fatalf("InsertLast is error")
	}

	check = circularDeque.IsFull()
	if !check {
		t.Fatalf("IsFull is error")
	}
}
