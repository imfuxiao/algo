package stack

import "fmt"

// 默认容量大小
var DefaultCap = 32

type ArrayStack struct {
	elements []interface{}
	cap      int // 容量大小
	size     int // 当前元素数量
}

func New(cap int) *ArrayStack {
	stackCap := cap
	if stackCap <= 0 {
		stackCap = DefaultCap
	}

	s := &ArrayStack{
		elements: make([]interface{}, 0, stackCap),
		size:     0,
		cap:      stackCap,
	}
	return s
}

func (a *ArrayStack) Size() int {
	return a.size
}

func (a *ArrayStack) Push(v ...interface{}) bool {
	if a.size == a.cap || len(v) > a.cap {
		return false
	}
	if len(v) > 0 {
		a.elements = append(a.elements, v...)
		a.size = a.size + len(v)
		return true
	}
	return false
}

func (a *ArrayStack) Pop() interface{} {
	if a.size > 0 {
		e := a.elements[a.size-1]
		a.size--
		return e
	} else {
		return nil
	}
}

func (a *ArrayStack) String() string {
	return fmt.Sprintf("%+v", a.elements)
}
